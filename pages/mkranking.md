[Home](/)

# MK Ranking
* Date: 2022
* Tech: Go, Docker, PostgreSQL
* Source: [github](https://github.com/thalkz/mkranking)
* Demo [fr]: [mkranking.thalkz.com](https://mkranking.thalkz.com)

An elo-based leaderboard for Mario Kart competition previously hosted on mkranking.com

## Motivation

In 2022, the Mario Kart competition was strong among the team at Revolt and the grind to get the number one spot was real. However, because of the randomness of the game (blue shells, hum hum), one or two races are not enough to know who the best player was. This is why I started working on a leaderboard, where we could enter our scores and know *for sure* who consistantly performed better.

Sounds too intense ? Well, during about 6 months, we used MK Ranking to add scores and for a while it was really part of the experience. Inevitably, it started to be more stressful than fun and we decided to leave it. In the few months that the competition lastest, many iterations to the project were made and the final state of the project was quite polished.

## Minimum Viable Product

The MVP for this project was all about getting a leaderboard up and running. Early on, I realised that the main friction point for users was that the score needed to be manually enterer after each game. With this in mind, I decided **not** to go for a mobile app, since both the deployment process (for me) and the downloading, installing and updating of the app (for everyone) would really hinder the leaderboard's adoption. I went for a server-side rendered webapp, which turned out to have some very nice benefits, first of all in terms of adoption.

The stack is as follows:

* A web server in Go, that uses templates to generate HTML files
* A PostgreSQL database for keeping scores

In retrospect, using Postgres for this project was overkill, but I haven't really used SQL databases that much before, and this made for a good learning experience.

## Early focus on deployment

From earlier projects, I learned that having well automated deployment process from the very start helped getting it to production. Why ? Partly because the friction of complicated deployments later or can be enough to abandon a project, and partly because showing progress to other people is motivating and allows for valuable feedback while things are not too much set in stone. Realising that the project needs a major pivot after months of work can break the motivation.

For deployment, I used a combinaison of Github Actions and Docker:

1) To start a deployment, I create a semver tag using [autotag](https://michaelcurrin.github.io/auto-tag/)
2) Push the commit and the tag, which triggers a Github Action
3) The job builds the Docker image of the webapp and pushes it to [ghrc.io](ghcr.io)
4) When the upload is completed, it connects to my VPS with SSH and runs a `update.sh` script, that pull the image with the correct version and start it with `docker compose` .

This deployment process makes it easy to have both local and production versions with different configurations. It's also easy to handle steps manually or to rollback to an earlier version if needed.

## A generalized ELO rating system

Elo is a rating system that add or removes points to players after each match. It's particularity is that is takes the "strength" of each player into account: if a high-rated player beats a low-rated player, it does not grant him many points. However if the reverse happens (an upset), than the high-rated player losses many points. For this, the algorithm makes an estimation of the probability that each player will win, and compares it to the actual outcome. The bigger the difference, the more points are granted.

This algorithm (or one of its variations) is used in chess, Starcraft, tennis and many other 2 player games. I needed to find a way to generalize the algorithm so that it would work for 2 to 4 players. Thankfully, [this article](https://towardsdatascience.com/developing-a-generalized-elo-rating-system-for-multiplayer-games-b9b495e87802) gives a very clear explanation on how elo works and how to generalize it to more players. My implementation can be found [here](https://github.com/thalkz/mkranking/blob/main/server/elo/elo.go). 

## Conclusion and learnings

This project was a learning experience on several aspects:

* for the developer experience, actually being a user of my own product made all the difference. It boosts motivation and attention to detail.
* in terms of product design, I could experience first hand the reactions to certain product choices, and how seemingly insignificant changes could make or break a product.
* a product does not necessarily need to be particularly beautiful to work. How it *looks* helps drive adoption, but retention is more related to how it *feels*.
* on the technical side, I've learned much about Server Side Rendering, SQL databases, Contenerization and deployment in general.

Many of those learnings were re-used for my next project: trikount.
