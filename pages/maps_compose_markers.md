[Home](/)

> NOTE: This is a written version of a conference given at Android Makers 2025, in Paris. The video should be available on the [Android Makers Youtube Channel](https://www.youtube.com/@DroidconParis) and goes a bit more in depth into the theory behind the implementation. Make sure to check it out (when it's out)

# Composable Markers for Google Maps

The Google Maps sdk is the most widly used map renderer in Android apps. There are other renderers of course, like Mapbox, but they all share the same underlying architecture: the map rendering is performed separately from the rest of the UI. This makes a lot of sense, of course, but it has some practicle consequences: it is not possible to render Composables (or other piece of dynamic UI) directly on the map. In other words: markers on the map are mostly static. They are essentially Bitmaps!

Using [official Maps library for Compose](https://github.com/googlemaps/android-maps-compose), there is a way to take a snapshot of a Composable, transfer it to map renderer and display it. However, if this Composable is animated, contains state or intercations —like clickable buttons— then the result will be dissapointing: all of these will be ignored when rendered.

Thankfuly, there are a some workaround to solve this issue, so let's explore them here.

> DISCLAMER: Many things in this post are simplified for making the point clearer and easier to follow and some of the code doesn't run as-is.

### The easy option - using Google Maps' Projection

The fastest way to render a Composable on Google Maps is to use the provided `Projection` class from the sdk. This class is provided by `CameraPositionState` and has a `toScreenLocation()` method. In the context of maps, the concept of projection is pretty straightforeward: it's the transformation of a value in 3D space (a LatLng) into a 2D space (a pixel on screen). I'm over-simpliying a bit, but that's the general idea.

With this in mind, we can create a minimal setup to display both a `GoogleMaps()` and a Composable `Marker` on top. Here, the `Projected` Composable is what we want to implement.

```
@Composable
fun MapScreen() {
    val cameraPositionState = rememberCameraPositionState()

    GoogleMap(
        modifier = Modifier.fillMaxSize(),
        cameraPositionState = cameraPositionState,
    )

    Projected(
        latlng = target,
        cameraPositionState = cameraPositionState,
    ) {
        Marker() // Composable we want to render on the map
    }
}
```

Let's create a basic `Projected()` Composable, using the `cameraPositionState`. This Compomsable has a clear purpose: to project provided `content` onto the Map, making it "follow" any map movement.

```
@Composable
fun Projected(
    target: LatLng,
    cameraPositionState: CameraPositionState,
    modifier: Modifier = Modifier,
    content: @Composable () -> Unit,
) {
    Box(
        modifier = modifier
            .fillMaxSize()
            .graphicsLayer {
                cameraPositionState.position // read position state to trigger re-draws when position changes

                val offset = cameraPositionState.projection?.toScreenLocation(target)
                if (offset != null) {
                    translationX = offset.x.toFloat() - size.width / 2
                    translationY = offset.y.toFloat() - size.height / 2
                }
            }
    ) {
        content()
    }
}
```

What's happening here ? We read the map's camera position during the drawing phase of `Projected`. Everytime the camera moves (by panning, zooming or rotating) we trigger a re-draw, computing a new offset using `toScreenLocation` and correctly placeing `Marker` on the screen.

This works in practice, but there is a catch! We are doing some two slow operations on the Main thread: getting the projection and calling `toScreenLocation`. Both are just wrapper that communicate with Google Play Services, via a Binder. Besides the actual computation that is needed to transform a LatLng into an offset, all transfered data is parcelized back and forweward, which can become costly when done multiple times per frame.

For just a few markers, this might be a perfectly viable option. Less code, less bugs!

However, when adding hundreds or thousands of markers, it starts to look bad: the Main thread is spending significant time getting the projection & calling `toScreenLocation`, resulting in missed frames.

There is a more advanced solution that requires move code, but can be very effective. How ? By performing the projection directly on the Main thread, avoiding the communication costs.

### The advanced option - implementing the projection kotlin-side

Just like most online maps, Google Maps is using a variation of the Mercator projection for its map called "web mercator" or "pseudo mercator". It is a simplified (and therefore less accurate) version of Mercator, optimized of oneline map rendering. It was popularized by Google Maps itself and has become the standard.

The rest of this post explains how to implement web-mercator in Kotlin, and ways to use it with Compose. We will start with a target LatLng (the position of our marker) as well as the current camera position (e.g latlng + zoom level + bearing). The output will be the position of the marker on the screen, in pixel, that well use in `Projected`.

To make things easier, let's decompose the computation into multiple steps: 

* step a) first, we take the target LatLng and apply the web-mercator formulas to them. This will output a point in "web mercator coordinates". 
* step b) using the zoom level, we will tranform that point into pixel coordinates. These coordinates are "absolute" and do not depend on the camera's position
* step c) we will transform the absolute pixel coordinates into screen coordinates, our final output

All these steps need to be re-run everytime the camera moves (panning, zooming or rotating the map), which can add up in terms of CPU usage. The good news is that we can reuse parts of this pipeline for multiple markers, effectively lowering the scaling costs.

In the final code, we will replace the `cameraPositionState.projection?.toScreenLocation(target)` line by `customProjection.toScreenLocation(target)`. `customProjection` is a class that we create ourselves.

Te begin with, here is the custom implementation of `toScreenLocation`, where we can clearly see the 3 steps:

```
fun toScreenLocation(target: LatLng): Offset {
    val point = target.toWebMercatorPoint() // step a
        .toPixels(scale) // step b
    return matrix.map(point) // step c
}
```

This methods uses two variables, `scale` and `matrix`. We'll explain what they are and how to compute them just below.

For step a, all we need it to implement the web-mercator projection. The formulas can be found online. In Kotlin, it looks like this:

```
fun LatLng.toWebMercatorPoint(): WebMercatorPoint {
    val lat = latitude.coerceIn(LAT_MIN_VALUE, LAT_MAX_VALUE)
    return WebMercatorPoint(
        x = (EARTH_RADIUS_METERS * toRadians(longitude)),
        y = (EARTH_RADIUS_METERS * ln(tan(PI / 4.0 + toRadians(lat) / 2.0))),
    )
}

private const val EARTH_RADIUS_METERS = 6378137.0
private const val LAT_MAX_VALUE = 85.05112878
private const val LAT_MIN_VALUE = -LAT_MAX_VALUE
```

There is some trigonometry involved, but overall there is not much going on! That's precisely the reason why this projection is so wildly used: it's simple ande fast to compute.

The output, a `WebMercatorPoint`, can be seen as a point on a 1:1 scale map of earth. So the units we are using here are actually meters.

Next, let's transforms those meters into pixels. For that we need to scale the map by a precise amount: the `scale`. The way `scale` is computed is linked to the [tile system](https://en.wikipedia.org/wiki/Tiled_web_map) of online maps, which we won't get into here. But this is where the formulas come from.

To the camera has a zoom (between 22 for max zoom and 1 for minimum zoom on Google Maps), and using that zoom, screen density and a constant we can compute a `scale` value:

```
fun computeScale(zoom: Double, screenDensity: Double): Double {
    return 2.0.pow(this) * TILE_SIZE * screenDensity
} 

const val TILE_SIZE = 256.0
```

This scale is used to, implement `toPixels`. We're going from a 1:1 maps of earth to a map correctly sized for our device.

```
fun WebMercatorPoint.toPixels(scale: Double): Offset {
    return (this + worldCenter) * scale / WEB_MERCATOR_SIZE
}

const val WEB_MERCATOR_SIZE = 40075016,6855784
val WEB_MERCATOR_CENTER = WebMercatorPoint(WEB_MERCATOR_SIZE / 2f, WEB_MERCATOR_SIZE / 2f)
```

The `Offset` returned here is in absolute coordinates on the map, correclty sized for our device. The last step is to take the camera's target & rotation into account, in order to know where to place the point on the screen. The good news here is that the only things left to do are linear tranformations, which can be represented with a `Matrix`. Using a matrix is both an efficient way to do the computations and a good way to share the pre-computed matrix between markers.

Here is how we create the Matrix, taking `scale`, `cameraTarget` and `cameraBearing` into account for handeling camera rotations. Note that we also need the center of the screen in pixels (assuming the map is full screen, of course), because its our pivot point for the rotation.

```
private fun computeMatrix(
    scale: Double,
    screenCenter: Offset,
    cameraTarget: LatLng,
    cameraBearing: Float,
): Matrix {
    val cameraTargetPx = target.toWebMercatorPoint().toPixels(scale)

    return Matrix()
        .apply {
            translate(
                x = screenCenter.x,
                y = screenCenter.y,
            )
            rotateZ(-cameraBearing)
            translate(
                x = -cameraTargetPx.x,
                y = cameraTargetPx.y,
            )
            scale(y = -1f)
        }
}
```

Finally, we can create our `CustomProjection` class, using all the helper methods defined above.

```
class CustomProjection(
    screenDensity: Float,
    screenCenter: Offset,
    camera: CameraPosition,
) {
    val scale = computeScale(camera.zoom, screenDensity)
    val matrix = createMatrix(scale, screenCenter, camera.target, camera.bearing)

    fun toScreenLocation(target: LatLng): Offset {
        val point = target.toWebMercatorPoint().toPixels(scale)
        return matrix.map(point)
    }
}
```

That's pretty much it! We have all the tools to perform a projection using Kotin only. For each new `CameraPosition` provided by Google Maps, we create an associated `CustomProjection` and pass it down to each `Projected` Composable. Both `scale` and `matrix` are computed only once per camera position, and than we call `toScreenLocation` for each marker, which is relatively cheap. No communication cost with Google Maps sdk!

Here is a way to link everything together in our `MapScreen`:

```
@Composable
fun MapScreen(modifier: Modifier = Modifier) {
    val cameraPositionState = rememberCameraPositionState()
    val density = LocalDensity.current
    val customProjection by remember { derivedStateOf {
        CustomProjection(density.density, screenCenter, cameraPositionState.position)
    } } 

    GoogleMap(
        modifier = modifier.fillMaxSize(),
        cameraPositionState = cameraPositionState,
    )

    Projected(
        latlng = target,
        customProjection = customProjection,
    ) {
        Marker()
    }
}
```

### Conclusion

This approach solves one of the issues with modern map sdk. It does have many drawbacks, though. 

Users move their map pretty fast, and any frame were the markers are not perfectly in sync with the underlying map will be noticable. Since we don't have any sync mechanism in place, there in no way to guarantee that this will never happen. For example, if the camera position returned by the map is out-of-date (which happens from time to time) or if we drop frames in the main thread, the markers will glitch a bit.

Another, harder-to-notice issue is related to point inputs, and more specifically: taps. For some reason, Compose is not always handling correctly the interoperability between the `AndroidView` (the map) and Composables on top. Adding a `clickable` modifier on our marker will make it impossible to start a pan gesture from that marker. It sounds unlikely, but it actually happens a lot. One workaround can be to handle hit-testing ourselfs for markers, which works but is far from ideal.

The code in this post is pretty basic, and there are many more things to add to this idea, both to improve performance, but also to unlock new features.

Hopefully this post helped you create a better, more dynamic map! 