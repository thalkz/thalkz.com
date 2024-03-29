[Home](/)

# Speed Typer
* Date: 2022
* Tech: Flutter
* Source: [github](https://github.com/thalkz/speed_typer)
* Demo: [gameplay video](/images/speedtyper.mp4)

A typing game in Flutter, made during a train ride from Lyon to Lille.

## Motivation

Around that time, Flutter's support for desktops was getting ready for production. The best way to test it is by making a small project, like Speed Typer. Turns out it's also a nice way to spend time during a long train ride. Speed Typer is a typing game, where players need to type as many words correctly in a given time. Each correct word gives a few more seconds. The words get longer and longer, which constrains the game time to under 1 or 2 minutes.

## Design Process

The game was so simple, the core features were really fast to implement. The main focus quickly became to maximize the satisfaction of playing the game. Putting the game into the hands of testers made this discovery process much faster. It's the small details that make a product like this work. Careful observation and empathy during user tests help to find what's preventing the player's experience from being great.

## Iterations from user tests

1) In one of the first iterations, when a typing mistake occurred, the next word would instantly be displayed. This turned out to be an issue for fast typers: a single mistake would often become a multi-word blunder. To fix this, I've added a one-second pause, with the whole word turning red and shaking. It's a classic, but very effective.

2) The first iteration had a timer in seconds displayed in the corner of the screen. But this turned out to be very distracting: when the time was almost up, players would bounce their eyes between the timer and the word to type, which turned into more typing mistakes and added unnecessary stress. This was solved by replacing the timer with a bar, right in the center.

3) Since the game is in French, some words have accents, which are harder to type. I tried removing them, but accents helped players understand at first glance what the word is. To accommodate for that, the game allowed players to type both the accentuated and unaccentuated letters, whichever is more convenient.

## Conclusion: it's all about the experience

Even the simplest game loop can be very satisfying when enough care is put into the experience. For Speed Typer, the current design is extremely minimalistic (and boring) and there is no sound. But with a proper game engine, particles, effects, and sound could be added, and with small user-tested iteration, the game could become interesting. Do one thing, but do it well.
