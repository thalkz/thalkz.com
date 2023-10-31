[Home](/)

# Trikount
* Date: 2022-23
* Source: [github](https://github.com/thalkz/mkranking)
* Webapp [fr]: [trikount.com](https://trikount.com)

A webapp to track expenses in groups, for travels, projects, etc.

## Motivation
Projects that are actually used (even by a few people) are much more motivating than anything else. As I was getting increasingly frustrated by the app I usualy used to share expense with my friends (ads, friction, bloated), I started working on a remake that I would actually enjoy to use.

## Low friction is a feature
Even if it's not usually seen as one, having very little friction for the users *IS* a feature, and one of the main feature, for trikount, at least. This translates to several design choices:

1) Webapp instead of a mobile app
2) Urls for sharing
3) No registration required
4) Maximize compatibility by using plain HTML

## Webapps vs mobile apps
I'm coming from a "mobile app" background, I have a natural tendency to want to turn everything into an app. But today, most mobile apps simply make no sense, and they add more friction than anything else (compared to the webapp equivalent). In some cases, turning a product from webapp to mobile app will *not* drive better retention. On the contrary: 

- Mobile apps add friction on the first use. Most people are reluctant to download another app: phone are a personal space, and allowing anything in is only done if convinced that it's worth it.
- Mobile apps have longer release cycles than webapps. That's just how the mobile stores work, the deployments, even when well automated, takes significantly more time than a webapp deployment. This usually leads globally for harder to maintain software, that more easily gets outdated (again, I'm talking big picture, here).

## Urls for sharing
I'd say that urls are well understood, at least by the kind of people that would be interested in trikount. Urls have the ultimate compatibility: they can be copied, shared on conversations, bookmarked, stored in notes, ect. I'm seeing more and more "multiplayer" products using plain links to share a session, from video calls, online documents, tools, etc. It's event better if the url does not point to an "invitation", but rather to the resource (document, call, etc) itself. This makes it possible to re-use the same-link later on, bookmark it, etc.

## No registration required
No registration is required for using trikount. Actually, it's not even possible to create an account. This is one of the central design choice for trikount: it allows someone to get a link and directly have access to all the shared expenses and add another expense directly. This was important to me, since I had seen many cases of people wanting to participate to shared expenses, but getting frustrated and defaulting to the "one person in the group tracks all expense", which is not ideal for many reasons.

Not having an account creates issues in itself: how to avoid loosing access to a group ? For now, a combinaison of web cookies and incitations to share the link in a conversation (both for sharing and "not loosing access") work well enough. One day, there might be an optional registration, maybe using [Web Authn](https://webauthn.io/) or equivalent.

## Maximize compatibility by using plain HTML
For this project, I've chosen to maximize compatibility by avoiding Javascript as much as possible. I have been pleasantly surprised to see the quality of the default HTML components, and how easily they can be combined to express interactions with the webapp. It has it's limits, of course, but the benefits are real:

- Compatibility with any kind of browser
- Future compatibility (aka: [this page is designed to last](https://jeffhuang.com/designed_to_last/))

## Conclusion
Most products cannot (and should probably not) be that focused purely be focused on that aspect. Users are used to the extremely high quality products they use every-day, most of them for free. This is usually supported by very intentional focus on quality, and immense resources to back it up. Having a feature-rich product is better, but only if the quality is there. Teams with lower amounts of resources (like a solo developer in his spare time) should focus on the very few features that actually make sense, and not try to beat the competition on every aspect, because it's not going to happen.