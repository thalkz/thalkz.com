[Home](/)

# Trikount
* Date: 2022-23
* Tech: Go, Sqlite, Docker
* Source: [github](https://github.com/thalkz/mkranking)
* Webapp [fr]: [trikount.com](https://trikount.com)

A web app to track expenses in groups and easily split the bill.

## Motivation
From previous projects (like [MK Ranking](/pages/mk_ranking)) I've seen how much better it when when working on a project that is used by people, even by a very few. Instead of guessing what to work on next, ideas come naturally from experience.

Thus, I started looking for an app or service I frequently use, but that didn't bring much satisfaction. At the time I was getting increasingly frustrated with an app we used to share group expenses with my friends. The concept is simple:

- when someone pays for the rest of the group, he adds the expense (manually) to the app
- at the end, the app displays which money transfers should be done to even out the expenses for everyone

This app was frustrating for me. It wasn't just the ads: many small interactions could be improved. It was feature-rich, for sure, and could technically do every we needed, but not particularly well (also, it looked like a fun side-project, which was probably the biggest deciding factor in the end).

I started working on a remake and soon realized what the main focus of this project would be: *reducing friction as much as possible*.

## Low friction is a feature
Even if it's not usually seen as one, having very little friction for the users *is* completely a feature. Lowering friction is an iterative process, which requires many feedbacks. The result is very far from perfect, but many interesting things came out of this. In the end, the focus was translated to several design choices:

1) Webapp instead of a mobile app
2) Using URLs for sharing
3) No registration required
4) Maximize compatibility

## Web app instead of a mobile app
I come from a "mobile app" background, I have a natural tendency to want to turn everything into an app. But today, most mobile apps simply make no sense, and they add more friction than anything else (compared to the web app equivalent). In some cases, turning a product from a web app to a mobile app will *not* drive better retention. On the contrary: 

- Mobile apps add friction on the first use. Most people are reluctant to download another app: phones are a personal space, and allowing anything in is only done if convinced that it's worth it.
- Mobile apps have longer release cycles than web apps. That's just how the mobile stores work, the deployments, even when well automated, take significantly more time than a web app deployment. This usually leads globally for harder to maintain software, that more easily gets outdated (again, I'm talking big picture, here).

## Using URLs for sharing
I'd say that URLs are well understood, at least by the kind of people that would be interested in trikount. Urls have the ultimate compatibility: they can be copied, shared on conversations, bookmarked, stored in notes, etc. I'm seeing more and more "multiplayer" products using plain links to share a session, from video calls, to online documents, tools, etc. It's even better if the URL does not point to an "invitation", but rather to the resource (document, call, etc) itself. This makes it possible to re-use the same link later on, bookmark it, etc.

## No registration required
No registration is required for using trikount. It's not even possible to create an account. This is one of the central design choices for trikount: it allows someone to get a link and directly have access to all the shared expenses and add another expense directly. This was important to me since I had seen many cases of people wanting to participate in shared expenses, but getting frustrated and defaulting to the "one person in the group tracks all expenses", which is not ideal for many reasons.

Not having an account creates issues in itself: how to avoid losing access to a group? For now, a combination of web cookies and invitations to share the link in a conversation (both for sharing and "not losing access") works well enough. One day, there might be an optional registration, maybe using [Web Authn](https://webauthn.io/) or equivalent.

## Maximize compatibility
For this project, I've chosen to maximize compatibility by making the web app as static as possible, which means avoiding Javascript as much as possible. I have been pleasantly surprised to see the quality of the default HTML components, and how easily they can be combined to express interactions with the web app. It has its limits, of course, but the benefits are real the greater compatibility with any kind of browser and overall more future-proof.

Plain and minimalistic HTML pages like this also load fast, lowering the need for any further optimization. It's fast by default, uses less data, and works better on bad internet connections.

## Conclusion
Of course, I'm not arguing that all products need to be plain HTML web apps without registration. For most apps, those limitations are way too hard and do not allow for the complexity required by a complete product. However, it's easy to overestimate the feature scope that a given team can handle and to underestimate the number of iterations necessary to make it work well.

Users are used to the high-quality products they use every day, and this is why teams try to match this, by adding a ton of features to their products. But having a feature-rich product is better only if the quality is there to back it up. This requires true focus and resources to bring quality.

Teams with lower amounts of resources should focus on the very few features that make sense, and not try to beat the competition in every aspect.