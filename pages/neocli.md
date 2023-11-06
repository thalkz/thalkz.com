[Home](/)

# Neocli
* Date: 2021-22
* Tech: Javascript
* Co-authored with: [nihiloproxima](https://github.com/nihiloproxima)

A Javascript CLI to handle all kinds of game-related operations in Neopolis.

## Motivation

Part of the experience of Neopolis relies on how dynamic the game feels. Dynamic meaning that the game itself changed over-time, and was able to react on what players did. For this, the dev team needed to ba able to quickly and efficiently "interact" with the game by releasing and piloting live-events, interacting directly with players, as well as have a lot of control over things like the "seasonal updates". The moderation team also had a lot of work, in terms of chat moderation, cheat detection, conflict resolution, and overall had a very direct impact on the world of Neopolis.

For this, nihiloproxima created a powerful and extremely useful back-office (web app + backend), called the Neoconsole, which was used both by the dev team, the support staff and the moderators. However, since it was rather "heavy", updates on the Neoconsole itself took time and there was a struggle to keep up with the rapid evolutions of Neopolis.

At Revolt Games, we embraced the chaos, in a way. Everything changed fast and decisions were taken all the time: this was simply the kind of culture we had. So the tech needed to be able to to accomodate for that. This is why I started working on a CLI that would handle different kinds of tasks and better "stay in sync" with the rest of the game.

## When CLI make sense

Building a CLI for an internal tool is useful for multiple reasons:

1) They have a very high "logic to code" ratio. Meaning that most of the code in a CLI is directly useful and provides real value. All the outputs are in text formats, which is a exellent limitation: it mostly cancels the craving to "make it beautiful". Tge internal tool's job is not to be beautiful: it is to be useful.
2) When designed with the Unix philosophy in mind, it can be combined with other programs for even more power. Piping through your CLI allows for very flexible workflows and removes the need to update the tool each time something new needs to be done.
3) Finally CLIs are easily distributed. Our setup used Github Action to deploy the CLI to ghcr.io as a private npm package. After a simple configuration on the devs laptop, a new version of the tool is installed with a simple `npm install tool-name` command. A nice NPM package, `update-notifier`, helps to notify devs that there is a new version available, by displaying a message like this upon CLI usage.

```
   ╭────────────────────────────────────────────────╮
   │                                                │
   │        Update available 0.9.2 → 1.0.0          │
   │   Run npm i -g @revoltgames/neocli to update   │
   │                                                │
   ╰────────────────────────────────────────────────╯
```

## Conclusion
At Revolt Games, I spent time thinking how different tools and workflows could and should be combined togeother. Good automation, tools and process can multiply the productivity of a team, devs or not. This is defenitely what we were looking for: keeping the team small but extremely effective. When visiting our office, visitors were surprised to see how small the dev team was. In this we took much pride in.
