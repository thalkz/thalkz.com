[Home](/)

# Terra
* Date: 2022
* Tech: Go, GCP, PostgreSQL, Redis, H3, OpenAPI
* Links: [Official docs](https://revoltgames.github.io/docs/terra/intro), [OpenAPI specs](https://terra.neopolis.app/docs/)

A program that processes and combines various sources of real-world data al well as an API to efficiently serve the data in hexagonal tiles

## Motivation

[Neopolis](/pages/neopolis) being based on real-world data and requires fast access to relevant data. Over the course of the years in developement of Neopolis, multiple systems were designed to handle core gameplay features, like :

* find the buildings closest to the player's location
* a travel system that used virtual visas for visiting foreign countries
* closest city and regions, to create geo-located chat rooms and rankings
* a tiling system for players to viretually "own land"

Even though theses independant systems were more or less compatible, they had both as performance issues and inconsistencies (edge-cases like borders, cities lists, etc.) that made development slower. At the time, part of the team was working on another project called [Neoland](https://neoland.io), which needed to be compatible with Neopolis on a geographical level. Moreover, the goal of Revolt Games was to create more real-world based projects.

With this in mind, I started working in early 2022 on Terra, that would solve these problems, unify existing projects and create a solid base for further geo-located projects.

## A unified world-view

Terra aims to provide a unified "ground truth", with what's required for devs. In reality, there is not one ground trouthand multiple world-views exist, depending on who you ask. Terra's goal is not to take a political stance (or at least tries to take the most commonly accepted world-view), but simply provide useful real-world data for Neopolis and similar projects. Careful selection of data sources attempts to mitigate this issue, but will never be perfect for everyone. 

Here were some of the things Terra provides :

* A world-wide tiling system, based on [Uber's h3 hexagonal tiles](https://h3geo.org)
* precise definition of where countries border stand, and how it is handeled by the tiling system, as well as some terrain marked as "disputed grounds"
* what actually counts as a separate country, or not (there is no worldwide concensus on this, of course)
* what the major cities in each country are, as well as unique identifiers for each city (Neopolis requires a curated list of major cities)
* what terrain is associated with each major city (each tile has a "closest city" field)
* which "buildings" (called places in the Terra) exist in each tile, as well as categorisation (based on [Foursquare's Place API](https://fours))
* what kind of terrain is contained in each tile, in percent of landuse (forest, water, etc)

## Sources of data

The data provided by Terra can be grouped into 3 categories:

* Cultural data, countries, and major cities
* Natural data, called biomes in Terra (urban, forst, desert, etc)
* Curated places (aka. real world buildings), and useful meta-data about them

A large part of this project was to find high quality sources of data. I've landed on these main sources:

* For cultural data, [Natural Data](https://www.naturalearthdata.com/) has an incredible list of vector maps, updated by a Map Update Committe that is well regarded in the carthography world. This data is denerously given as public domain.
* For natural data, [GLAD](https://glad.umd.edu/) provied a CC BY-licensed dataset of [Global Land Cover and Use](https://glad.earthengine.app/view/global-land-cover-land-use-v1), that for Terra was downloaded as a (huge) GeoTiff image
* The curated places comes from [Foursquare](https://foursquare.com) and is therefore NOT redistributed outside Revolt Games.

## Data processing

All the downloading from these different sources, processing, curating, combining them into a well-defined format is done by a single Go program. Why Go ? Because the huge amounts of data (yes the world is really big) required to take full advantage of multiple cores. My Macbook Pro M1, the 8 CPU cores working at max capacity, was powerful enough to process the whole world in about 10 hours in total. For a while, my computer looked a bit like this:

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

The final output of the program is a PostgreSQL database, with multiple related tables that can be easily queried, joined and served

 - 198 countries
 - 5k cities
 - 700M+ lands (e.g terrain tiles), not including the 70% of oceans covering earth

Because of the size of input and output data as well as the the different formats in data sources (vector data, csv, GeoTiffs, ect.) I designed the processing in multiple stages, which could be turned on or off and configured via [a configuration file](/pages/terra_config) in JSON before lauching the pipeline. Broadly, the stages are:

1) Bronze stage: Download data and cache raw data locally (when possible), to avoid dowloading huge file on each start
2) Silver stage: Input all bronze data (in different formats) into a "Silver" Postgres database, with little to no modification
3) Gold stage: Trough queries to the silver database, combine and process data into its final format and input it into a "gold" database

The 3 stages of processing allowed for clearer responsability of code, as well as reusability of stages, which made it easy to restart only the relevant stages and save hours of computing. Once fully processed, I've uploaded the final 20Go (compressed) gold database to GCP and created an API to serve it.

## Serving the data
The data is served via an API, used by both the Neopolis backend and the Neoland backend. The server is pretty straightforeward: it makes (carefully crafted) queries to Postgres, with databses indexes to make common queries fast.

A caching layer takes full avantage of both the fast that some data is rarely updated, as well as the tiling system. On some queries, large amounts of "lands" need to be queries and therefore, as an optimization, tiles are sometimes grouped together, which was implemented using the parenting hierarchy of h3 tiles.

TO easier adoption, I've documented on [a github page](https://revoltgames.github.io/docs/terra/intro) the API's enpoints with example queries and responses, as well as detailed explanations of the data contained in Terra. A [Open API page](https://terra.neopolis.app/docs/) is also generated at compile time.

## Conclusion and learings
Working on Terra has been especially interesting on the data-processing side. Breaking the processing pipeline into multiple stages (bronze, silver and gold) that are idempotent and produce intermediary steps was a game-changer. I was impressed by how powerful PostgreSQL is: carefully crafted SQL queries goes a (very) long way. 

Finally, this project was a good optimization challenge: some stages took multiple hours of computing time, so each performance boost was useful. My biggest focus here was to cleanly spread the workload acrodd multiple CPUs, for which Go channels and Goroutines were very convinient.