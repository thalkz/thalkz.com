[Home](/)

# Terra
* Date: 2022
* Tech: Go, PostgreSQL, Redis, H3, OpenAPI, Google Cloud Platform
* Links: [Official docs](https://revoltgames.github.io/docs/terra/intro), [OpenAPI specs](https://terra.neopolis.app/docs/)

A program that processes and combines various sources of real-world data as well as an API to efficiently serve the data as hexagonal tiles.

## Motivation

[Neopolis](/pages/neopolis) is based on real-world data and therefore requires fast access to geographical data for the users. Over the years of iterations, multiple systems were designed to handle core gameplay features, like :

* a system to efficiently find available buildings closest to the player's location
* a virtual travel system for visiting foreign countries
* a closest city and regions, to create geo-located chat rooms and rankings
* a tiling system for players to "own" virtual land

Even though these independent systems stayed more or less compatible, they had both performance issues and inconsistencies (poorly defined borders created bugs, for example) that made development slower and posed compatibility issues. At the time, part of the team was working on a project called [Neoland](https://neoland.io), which needed to be fully compatible with Neopolis on a geographical level. Moreover, the goal of Revolt Games was to create more real-world based projects.

With this in mind, I started working on Terra in early 2022 to solve these problems, unify existing projects, and create a solid base for further geo-located projects.

## A unified world-view

Having a unified worldview is complicated. In reality, there is not one single ground truth that everyone agrees on. Some countries are recognized here but not there, there are disputed borders, there are different views on which cities are important or not, etc. All of this means that multiple worldviews exist, depending on who you ask. Terra's goal is *not* to take any political stance: it tries to take the most commonly accepted world-view and for this is based on carefully selected sources that seem trustworthy and well recognized. However, I'm not a geopolitical expert nor a cartographer, so errors might very well have slipped in.

Anyways, here are some of the things Terra provides :

* A worldwide tiling system, based on [Uber's h3 hexagonal tiles](https://h3geo.org)
* precise definition of where countries' border stand, and how it is handled by the tiling system, as well as some terrain marked as "disputed grounds"
* what counts as a separate country, or not (there is no worldwide consensus on this, of course)
* what the major cities in each country are, as well as unique identifiers for each city (Neopolis requires a curated list of major cities)
* what land is associated with each major city (each tile has a "closest city" field)
* which "buildings" (called places in the Terra) exist in each tile, as well as categorization (based on [Foursquare's Place API](https://fours))
* what kind of terrain is contained in each tile, in percent of land use (forest, water, etc)

## Sources of data

The data provided by Terra can be grouped into 3 categories:

* Cultural data like countries and major cities
* Natural data, represented as "biomes" in Terra (urban, forest, desert, etc)
* Curated places (aka. real-world buildings), and useful metadata about them

Part of the project was to find high-quality sources of data. I've settled on these main sources:

* For cultural data, [Natural Data](https://www.naturalearthdata.com/) has incredible vector maps, updated by a Map Update Committee. This data is provided in the public domain.
* For natural data, [GLAD](https://glad.umd.edu/) provies a CC-BY-licensed dataset of [Global Land Cover and Use](https://glad.earthengine.app/view/global-land-cover-land-use-v1). In practice, it is a (huge) GeoTiff image of the world, where the color of each pixel corresponds to a category of land use (forest, water, etc)
* The curated places are from [Foursquare](https://foursquare.com), which closed license data and is therefore not redistributed.

## Data processing

A single program, the "Terra Pipeline" does all the downloading from these different sources, processing, curating, and combining of the data. Why Go? Because of the massive amounts of data (yes, the world is big) required to take full advantage of multiple cores. My Macbook Pro M1, with the 8 CPU cores working at max capacity, was powerful enough to process the whole world in about 10 hours in total. For a while, my computer looked a bit like this:

```
CPU         Use%
All
AVG         94%
CPU0        93%
CPU1        92%
CPU2        99%
CPU3        96%
CPU4        92%
CPU5        97%
CPU6        98%
CPU7        98%
```

The final output of the program is a PostgreSQL database, with multiple related tables that can be easily queried, joined, and served. It has :

 - 198 countries
 - 5k+ cities
 - 700M+ lands (e.g terrain tiles), not including the 70% of oceans covering the earth
 - A few more things things

Because of the size of input and output data, as well as the different formats in data sources (vector data, CSV, GeoTiff...), the processing is done in multiple stages. Stages can be turned on or off and configured via a configuration file ([see example configuration](/pages/terra_config)) in JSON before launching the pipeline. Broadly, there are 3 main stages: bronze, silver, and gold.

1) Bronze stage: Download and cache raw data locally (when possible)
2) Silver stage: Insert bronze data into a "Silver" Postgres database, with little to no modification
3) Gold stage: Through queries to the silver database, combine and process data into its final format and insert it into a "gold" database

The 3 stages of processing allowed for clearer responsibility of each code part, as well as reusability of stage outputs. This made it easy to restart only the relevant stages and save hours of computing time. Once fully processed, I've uploaded the final 20Go (compressed) gold database to GCP and created an API to serve it.

## Serving the data

The data is served via an API, used by both Neopolis and Neoland. The server is pretty straightforward: it makes Postgres queries, with indexes to make common queries faster.

A caching layer with Redis takes full advantage of both the fact that some data are rarely updated, as well as the tiling system. For some queries, large amounts of "lands" need to be returned and therefore, as an optimization, lands are sometimes grouped together. This optimization uses properties of H3, like parenting hierarchies.

The whole project is documented on [a GitHub page](https://revoltgames.github.io/docs/terra/intro) with the API's endpoints, example queries, and responses, as well as detailed explanations. A [Open API page](https://terra.neopolis.app/docs/) is also generated at compile time.

## Conclusion and learnings

Working on Terra has been particularly interesting on the data-processing side. Breaking the pipeline into multiple stages that are idempotent and produce intermediary steps made the development much faster. I've had to upgrade my SQL knowledge and overall I was impressed by how powerful PostgreSQL is: carefully crafted SQL queries go a (very) long way. 

Finally, this project was a good optimization challenge: some stages took multiple hours of computing time, so each performance boost was useful. My biggest focus here was to cleanly spread the workload across multiple CPUs, for which Go channels and Goroutines were very convenient.