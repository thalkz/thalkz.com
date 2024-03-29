[Home](/)

# Ma Petite Planète
* Date: 2020
* Tech: Flutter, Firebase
* Official website: [mapetiteplanete.org](https://mapetiteplanete.org)

The MVP of an app with ecological challenges played in teams.

## Motivation
Ma Petite Planète (MPP) is a non-profit organization that aims to get more people to take action for the protection of the environment. Many people want to help prevent climate change but do not know where to start. Realizing this, the MPP team designed a game, where players take on ecological challenges with their team. It's fun, full of learning, and has a long-lasting positive impact on the participants. It has been wildly successful in groups of friends, schools, and companies all around France, and is starting to be exported to other countries.

I met one of the co-founders 3 years ago when the organization was still in its early stages (it now has multiple hundred collaborators !). Everything was done through shared Google Sheets. It worked at a small scale, but the limitations made it hard to get more participants on board. I helped build a first draft of a mobile app, that would then become their main product.

## The impact of initial tech choices
Technical choices made at the very beginning of a project have a huge impact on a product's life. Most choices are close to impossible to undo, especially if there is good traction on the product: the team is "pulled" to focus on improving the product. During development, the risks associated with a major technical switch are great: it's difficult to estimate how long it will take, does not bring any direct value to users, requires re-training for the devs, etc.

I've witnessed this at MPP, where much of the app relies on Firestore. Firestore helped a lot in getting the app up and running: it virtually eliminates the need for a backend server, and therefore for a backend developer. All the business logic can be handled client side, which worked well but has some hard limitations. Firestore is amazing tech, but it is not for everyone.

Over time, some of MPP's core feature set changed (which is completely normal for this kind of product), and the initial choices no longer matched the tech, especially for Firestore. This forces devs to build complicated workaround and triggers internal discussions around of idea of switching to something else, or not. This hesitation itself draws a lot of energy from the teams and reduces focus.

As some of the tech choices as MPP come from the first version I helped create, this is mostly my bad. I wish I had taken more consideration in these initial choices. Today, the devs are handling the project incredibly well, but better initial choices would have been much more convenient for them.

## Making better choices
There are ways to mitigate these difficulties, of course, like having a clear separation of concern in the codebase, encapsulating outside dependencies, etc. Strict application of these principles takes time and requires experience and great engineering efforts overall. Good initial tech choices help the team focus on what matters the most: delivering value to users. 

So what are good technical choices, initially? There is no single answer to this question, of course. A good understanding of the tech *and* its limits helps. Brand-new tech comes with unknown limitations, which can be an issue. Making benchmarks and testing stuff helps too. It's also essential to know what the vision of the product is: in which global direction is it going to evolve? What are its most important aspects? Visual quality? Long-term maintainability? Feature-completeness? Robustness? Flexibility? Answering these questions helps to build an overall idea of what kind of tech is required.

A good rule of thumb for me is that a product should have mostly [boring tech](https://boringtechnology.club/), except in one or two very carefully chosen cases, depending on what the product does.

## Going forward
My direct contribution to the codebase didn't continue much after the initial launch of the first version (working on Neopolis was taking all my available time). I'm glad I've had the chance to keep in contact with the MPP team and especially [Leila](https://github.com/LeilaCoquard), who's leading their internal teams of devs. 

Working on the first versions of something that grows so much and so fast is rewarding. The fact that it's a project for good that makes people take steps towards a more eco-friendly future is amazing. Even though my participation has been limited, it shows that working on something that makes sense and that people believe in makes all the difference.