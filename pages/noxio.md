[Home](/)

# Noxio
* Date: Built in 2020
* Tech: Flutter, Firestore, Flare
* Source: [github](https://github.com/thalkz/noxio)

Noxio is a two-player board game for mobile phones, inspired by Bacteria. 

## Motivation

When starting this project, I had a clear idea in mind: I wanted to do a remake of a game I had been obsessed with some years ago. A mini-game called Bacteria on a platform that no longer exists today. It looked [like this](https://www.youtube.com/watch?v=QHdSQ2QTEDg&ab_channel=crabest2) (warning: 2007 vibes ahead). Bacteria was easy to understand, but still had a great amount of depth to it, both in terms of tactics and strategy. It's not chess, but there is definitely a learning curve.

## The MPV & first iterations

The core feature would be player vs player matches, which could easily be done offline at first, with the two players playing on the same phone. Getting this to work was the very first milestone.

After implementing a simple grid with the few rules of the game, I quickly realized that to make PvP work, it needed to have an online option too: being in the same room to test the game is simply too much of a constraint for this game. The online mode uses Firestore, which stores the state of each match inside a Firestore document, where each update triggers an update to the document and updates both player's UI.

At this point, the game was mostly working, but the minimalistic design (a simple grid with colored cells) made it hard to see what your opponent's last move was: the new cells simply "blinked" into existence. The next iteration was to add some animations on the cells, and for this, the Flare (now called rive.app) allowed to create stateful animation, that can be triggered with code.

## Adding an "AI" Player

For someone to play against the computer, an "IA" needed to be added to the game. This IA is obviously not the kind of IA we now see in modern machine learning techniques (though it would be interesting to try to use one for Noxio): I used a technique often used for chess IAs: [Minmax](https://www.chessprogramming.org/Minimax). My [implementation](https://github.com/thalkz/noxio/blob/master/lib/utils/ai.dart) is not very efficient but still uses a technique called "alpha-beta pruning" to discard bad moves early on and make the computation faster.

My implementation has multiple levels of difficulty, which computes the next move with more or less depth. This IA has some obvious flaws that can easily be exploited, but even with this advantage, I've never managed to beat the "expert" level...

## A few learnings from trying to re-run Noxio

Reviving a project that has been taking dust for 3 years is a mixed experience, some good findings, but a lot of frustration too. It took me a while to be able to start it on my Macbook again since I wanted some screenshots for this post.

* The lack of README for this project wasn't a real blocker, but made the process a bit less pleasant. Good Readme helps to get going (especially with a few screenshots, a todo list, etc)
* Having few dependencies makes it easier to re-use the project 3 years later, but since semver is rather well respected for the packages I used, it wasn't too hard to do the breaking change upgrades one by one.
* The Dart migration to null-safety is rather painful. I made an attempt with the dart migration tool from dart 2.19.6, but the automatic results were not very good, so I won't migrate it for now.

## Appendix: Noxio's gameplay

Noxio is played on an 8 by 8 grid and each player controls a single "cell" at the beginning. Their goal is to control the most cells when the game ends, which is when one of the two players can no longer make any move. Each turn, players can either:

* A) Duplicate one of their cells to a nearby empty cell
* B) Make one of their cell "jump" over 1 cell to another empty cell.

When one of your cells is created (by duplication or jumping) next to the cell of your opponent, you convert all the cells that are directly adjacent to this cell (diagonals included). This simple rule makes it complicated to see what's going to happen in a few moves.

For more diversity, different terrains exist, where some cells are marked as "rocks" and are blocked (cells can still jump over one rock, though).
