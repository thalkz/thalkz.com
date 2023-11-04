# Example config.json for Terra
```
{
  "steps": {
    "silver_countries": "override",
    "silver_cities": "override",
    "silver_biomes": "override",
    "silver_buildings": "skip",
    "silver_altitudes": "override",
    "silver_mountains": "override",
    "silver_markers": "override",
    "gold_countries": "override",
    "gold_cities": "override",
    "gold_biomes": "override",
    "gold_lands": "override",
    "gold_places": "skip",
    "gold_markers": "skip",
    "gold_countries_stats": "override",
    "france_gold_lands": "override"
  },
  "global": {
    "land_h3_resolution": 8,
    "workers_count": 8
  },
  "silver_database": {
    "host": "localhost",
    "port": "5432",
    "user": "postgres",
    "password": "postgres",
    "db_name": "terra_small_silver"
  },
  "gold_database": {
    "host": "localhost",
    "port": "5432",
    "user": "postgres",
    "password": "postgres",
    "db_name": "terra_small_gold"
  },
  "silver_cities": {
    "shp_folder": "assets/natural_earth",
    "all_cities_shp_filename": "ne_10m_populated_places_simple.shp",
    "primary_cities_shp_filename": "ne_50m_populated_places_simple.shp"
  },
  "silver_countries": {
    "shp_folder": "assets/natural_earth",
    "shp_filename": "countries_large.shp",
    "neopolis_countries_csv_folder": "assets/cities_api",
    "neopolis_countries_csv_name": "neopolis_countries.csv"
  },
  "silver_markers": {
    "filepath": "assets/foursquare/usa_quality.tsv"
  },
  "silver_biomes": {
    "pixel_resolution": 5,
    "image_size_x": 40000,
    "image_size_y": 40000,
    "image_size_lat": 10,
    "image_size_lng": 10,
    "tiff_folder": "assets/earth_engine",
    "download_url": "https://storage.googleapis.com/earthenginepartners-hansen/GLCLU_2019/strata/",
    "tiff_filenames": ["40N_000E.tif"]
  },
  "silver_altitudes": {
    "pixel_resolution": 10,
    "image_size_x": 6000,
    "image_size_y": 6000,
    "image_size_lat": 5,
    "image_size_lng": 5,
    "tiff_folder": "assets/srtm",
    "download_url": "https://srtm.csi.cgiar.org/wp-content/uploads/files/srtm_5x5/TIFF/",
    "tiff_filenames": []
  },
  "silver_mountains": {
    "pixel_resolution": 3,
    "image_size_lat": 140,
    "image_size_lng": 45,
    "tiff_folder": "assets/global_mountain/",
    "tiff_filenames": []
  },
  "gold_countries": {
    "population_ceil_to": 1000,
    "remove_min_population": 0,
    "remove_country_ids": [],
    "buffer_size": 0.1,
    "fra_geojson_file": "assets/geojson/fra.geojson",
    "guf_geojson_file": "assets/geojson/guf.geojson",
    "guf_i18n_file": "assets/i18n/guf.json"
  },
  "gold_cities": {
    "population_ceil_to": 1000,
    "remove_min_population": 1000,
    "remove_city_ids": [],
    "keep_city_ids": [],
    "set_primary_city_ids": []
  },
  "gold_biomes": {
    "remove_min_percent": 5,
    "low_mountain_weight": 2,
    "scattered_low_mountain_weight": 1,
    "high_mountain_weight": 3,
    "scattered_high_mountain_weight": 2
  },
  "gold_lands": {
    "france_lands_csv_file": "assets/neoland/lands_france.csv",
    "process_france_only": true
  }
}
```