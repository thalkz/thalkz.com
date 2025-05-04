[Home](/)

# Year of Knowledge
* Date: 2023
* Tech: Obsidian, Webdav
* The 365 notes: [Will be published later]

A year of learning one thing per day on different subjects, including learnings when traveling in Europe. I only started in March, but found the time to write something almost every day.

## Motivation
The Year of Knowledge is part of my [yearly theme](https://www.youtube.com/watch?v=NVGuFdX5guE&ab_channel=CGPGrey) for 2023. The theme being purposefully vague, I've taken a bit of time to figure out how to concretely work towards this theme. Then, it clicked: just write something I learned each day, whatever it is. Like most people, I remember information much better when I've done the job of summarizing it and writing it down.

## The setup
To achieve 365 days of writing something new each day, I needed to have an editor both on my phone and my computer. The phone for most days, is accessible from anywhere, and the computer is mostly for editing, for longer inputs, or simply when I am in the mood for it.

[Obsidian](https://obsidian.md/) turned out to be perfect for that purpose: the text editor is simple and powerful and the text is plain markdown. Obsidian has a paid synchronization feature, that looks very well made and completely worth it in most cases. Still, I started looking for alternatives and handling the sync "myself". After a quick failed attempt of using Git, I settled with the community plugin [Remotely Save](https://github.com/remotely-save/remotely-save) and a WebDAV server on my VPS, using a nice Go implementation of the protocol: [dave](https://github.com/micromata/dave) (now unmaintained).

## All days (french only)

# 2023-03-05 IFSC Climbing World Cup

#escalade #competition

La IFSC (International Fédération of Sport Climbing) organise les Climbing World Cups, une série de compétitions d'escalade depuis 1989. Chaque cup rapporte des points aux participants et a la fin de la saison, le top de chaque discipline (lead, boulder, speed, combined) est déclaré champion. Les Climbing World Championships sont organisées tous les 2 ans. 2023 Berne, 2021 Moscou, 2019 Hashioji, 2018 Innsbruck.

# 2023-03-06 Le Matterhorn

#geographie #histoire #escalade 

Le mont Cervin (Matterhorn) a été gravi pour la première fois par l'alpiniste Edward Whymper (britannique) en 1865. L'expédition était la 7 ou 8ème tentative de Whymper et résultat dans la mort de 5 personnes de la cordée, tombés dans le vide.

Le mont Cervin, culminant a 4476m se situe sur la frontière Italienne et Suisse. Le village le + proche côté suisse est Zermatt.

# 2023-03-07 Les Workspace VS Code

#code #tools

Il est possible de créer un Workspace sur VS Code pour regrouper plusieurs projets dans la même fenêtre. Pour cela il suffit de créer un fichier .code-workspace et de lister les dossiers concernés. Ce Workspace peut avoir ses propres settings et extensions.

# 2023-03-08 La geo-ingenieurie

#ecologie #energie #technologie 

La geo-ingenierie solaire est un type de geo-ingenierie qui consiste à limiter le rayonnement solaire sur la Terre afin de conserver un réchauffement climatique sous la barre des 1.5°C. Fortement contestée, ces techniques s'attaquent aux symptômes du réchauffement climatique sans résoudre les causes. Elles présentent de plus d'énormes incertitudes par rapport à leur Impact sur l'atmosphère et le climat.

Il existe deux techniques principales étudiées : la diffusion de sulfate dans la stratosphère (effet aérosol), la propulsion d'eau salée dans l'atmosphere (reflexion sur le sel) et l'amincissement des cirrus, nuages renforçant l'effet de serre.

# 2023-03-09 Le Merkle Tree

#algorithme #git #code

Un Merkel tree (ou Hash tree) est une data structure sous forme de tree où chaque leaf node contient le hash d'un blob de data (par ex un fichier) et chaque inner node (ou inode) contient le hash de ses nodes enfants.

Un Merkle tree permet de vérifier rapidement l'intégrité d'un gros volume de données, notamment dans le cas de transfert de données en peer to peer. Il est utilisé, entre autre, sur IPFS (InterPlanetary File System), git, Ethereum et Bitcoin.

# 2023-03-10 Les réacteurs nucléaires en France

#ecologie #energie #france

Entre 1950 et 1960, 11 réacteurs nucléaires (1ere génération) ont été construits en France. Ils sont ajd tous a l'arrêt

Entre la fin des années 70 et les années 80, 58 réacteurs de 2eme génération (appelés REP) ont été mis en service.

La 3eme génération en préparation sont les EPR. Ils ont un rendement de 37% vs les 33% actuels et d'une durée de vie de 60 ans.

# 2023-03-11 Les ours des Pyrénées

#ecologie #geographie

Il reste environ 70 ours dans les Pyrénées.

# 2023-03-12 Les notations en Escalade

#escalade

Il existe différents systèmes de notation de la difficulté des voies en escalade. Pour la voie (sport climbing), les deux principaux sont le Yosemite Decimal System (YDS), utilisé aux US, et le French Scale (ou Fontainebleau Scale), utilisé partout ailleurs. 

Le YDS s'étend de 5.2a à 5.15c (le 5 signifie que c'est de l'escalier sur rocher). 5.9 correspond a 6a et 5.12 à 7a.

Le French Scale s'étend de 3a à 9c et s'applique aussi bien pour le bloc que pour la voie.

Pour le bloc, le Vermin Scale est utilisé aux US. V4 correspond a 6a et V6 à 7a.

# 2023-03-14 Orpierre

#escalade #france

Orpierre est situé dans le département des Hautes Alpes (05), dans la région Province-Alpes Côte d'Azur. La commune se trouve dans le parc régional naturel des Barronies provincales.

Neuf falaises de calcaire entourent le village, et la première voie d'escalade a été ouverte en 1980. La mairie est propriétaire de la plupart des falaises. Il y a aujourd'hui prêt de 500 voies accessibles.

# 2023-03-15 Les Saisons

#calendrier

Le printemps commence ~20 mars
L'été commence le ~20 juin
L'automne commence le ~20 septembre
L'hiver commence le ~20 décembre

# 2023-03-16 Starlink

#espace #technologie 

Starlink est une constellation de satellites opérée par SpaceX, fournissant internet n'importe où dans le monde. Cela coûte 450€ pour l'antenne + 50€/mois. En France, les débits en download sont compris entre 50 et 200 Mbps.

Les premiers satellites ont été envoyés en 2019. En février 2023, il y en avait déjà 3600 en orbite basse. Au total, 12 000 satellites doivent êtres deployés. En 2022, déjà plus d'un million d'utilisateurs se sont inscrits à l'abonnement.

# 2023-03-18 Les Okonomiyakis

#nutrition #voyage 

Un okonomiyaki est un plat japonais, qui est la contraction de okonomi = "ce que vous aimez" et Yaki = "grillé". C'est une sorte d'omelette dont la pâte est faite d'oeuf, de farine de blé, de dashi et de chou blanc. A la pate est rajouté un ingrédient au choix, tel que du porc, de la dinde, du kimchi de la sèche ou encore des crevettes.

# 2023-03-20 Les Origines de Daech

#histoire #geographie 

Daech, aussi appelé État Islamique (Islamic State of Irak and Syria, ISIS en anglais) est une organisation terroriste issue du groupe Al-Qaïda en Irak. Formée en 2006, elle proclame un Califat en sur le territoire contrôllé. Ce territoire est reparti entre la Syrie et l'Irak. A son apogée, en 2014/2015, le territoire de Daech s'étend sur une surface similaire a celle se l'Angleterre. De nombreux groupes djihadistes, tel que Boko Haram au Nigeria, ont prêté allégeance à Daech, renforçant leur autorité. Depuis 2015, la puissance de Daech s'affaiblit jusqu'a la mort du chef d'Abou Bakr al-Baghdadi, en 2019.

# 2023-03-21 Le déconventionnement des Falaises

#escalade #france

Depuis de nombreuses années, la FFME (Fédération Française de Montagne et d'Escalade) proposait une "convention" aux propriétaires de falaise en France (privé ou collectivités) : la FFME s'occupe de l'équipement et de l'entretien en échange de l'ouverture du site au public. La FFME assure aussi la responsabilité légale en cas d'accident (car le Code Civil stipule que le propriétaire est responsable, par défaut).

Cependant, suite à de graves accidents sur des falaises conventionnées (ex: accident de Vingrau en 2010, ou la fédération a payé + de 1.5M en dommages et intérêts), il a été décidé de "deconventionner" les falaises, rendant la responsabilité au propriétaire.

En 2020, un amendement limite la responsabilité du propriétaire (facilitant donc le deconventionnement). Malgré cela, des falaises telles que Presles (Vercors), Aurielle (Alpilles), les Goudes (Calanques) ont été fermées. Depuis janvier 2023, il n'y a plus de falaises conventionnées en France.

# 2023-03-22 Classification Naïve Bayésienne

#code #machine-learning #algorithme 

La classification naïve bayésienne est un type de classification basée sur le théorème de Bayes. Cette technique peut être utilisé pour détecter du spam, en se basant sur la probabilité qu'un mot donné apparaisse dans un spam ou non.

# 2023-03-23 Talibans en Afghanistan

#histoire #geographie 

Les talibans sont au pouvoir en Afghanistan Dupuis le retrait des troupes américaines, en 2021. Les talibans reprennent alors Kaboul sans combat, 20 ans après avoir été chassés.

# 2023-03-24 Le terme Juice

#user-experience #product

Le terme "Juice" represente tous les éléments non-essentiels: visuels, audio et haptic qui améliorent l'expérience utilisateur.

Ajouter des transitions, des feedbacks, des animations, des sons permet a l'utilisateur de s'impliquer plus émotionnellement dans un produit/jeu et d'avoir plus de satisfaction.

https://garden.bradwoods.io/notes/design/juice

# 2023-03-25 Mejdi Schalck et Paul Jenft

#escalade #competition 

Mejdi Schalck est un grimpeur français né en 2004. Il intègre l'équipe de France d'escalade en 2019 et remporte la médaille d'or aux Championnats de France de Bloc à Valence en 2023.

Paul Jenft est un grimpeur français né en 2003. Il s'est placé 10eme aux Climbing World Cup a Briançon en 2020 (en difficulté) puis 6 et 8eme a Villars et Chamonix en 2021

# 2023-03-26 Les Voies Lyonnaises

#lyon #ecologie

Les voies Lyonnaises est un projet de la Métropole de Lyon visant à aménager 12 voies traversant la métropole. Les travaux ont débutés en 2022 avec la ligne 1 (quais de Rhône) et la ligne 3 (à Couzon). En fin d'année 2023 commenceront les travaux de la ligne 2 à Part-Dieu et notamment Boulevard Vivier Merle (boulevard qui passe devant la gare). Les travaux dans Lyon doivent arriver à leur terme en 2026, avec les extensions de ligne se terminant en 2030.

# 2023-03-27 Map et Object en JavaScript

#code #javascript

The correct way to add values to a `Map` in JavaScript is with the `map.set(key, value)` function. Using the `map[key]=value` notation seems to work at first, but the key-value pair is ignored by methods on the Map, like `map.has(key)` and `map.delete(keys)`.

Maps and Objects are not the same thing. They are similar In the fact that they are both key-value pairs. But there are notable differences :
- an Object contains values by default (prototype)
- an Object is vulnerable to object injection attacks, so a Map is safer
- any value can be used as key in a Map. For Objects, only String and Symbols are allowed
- Map is iterable and the keys are ordered in the order they were created
- Map has a size property
- Map has better performances when adding & removing many keys

`Map` was added with ES6 (2015) and is not supported by all browsers.

# 2023-03-28 Le Garbage Collector

#code #algorithme #flutter

Le Garbage Collector (GC) est une technique de gestion de la mémoire allouée dans un programme. Inventée en 1959 par John McCarthy pour le langage Lisp, elle permet d'automatiser la désallocation de mémoire inutilisée, simplifiant donc la tâche du programmeur.

Le GC présente des avantages (évite les "dangling pointers", les "double free bugs" et certains memory leaks) et ses inconvénients (moindres performances).

Il existe différentes méthodes de GC :
- Le "Tracing", le plus commun, parcours à intervalles réguliers les objets du programme et libère ceux qui n'ont aucune référence. C'est utilisé en Java, en Go, en Dart en C#, D et dans les langages de scripts.
- Le "référence counting": chaque objet compte le nombre de fois qu'il est référencé. Si ce nombre tombe a zéro, il est libéré.
- La technique du "Escape Analysis" permet d'allouer un objet sur la stack plutôt que la head, si l'objet ne sort pas de son scope (utilisé en Go)

# 2023-03-29 La Convention Citoyenne pour le Climat

#ecologie #france 

D'octobre 2019 à juin 2020 s'est tenu la Convention Citoyenne pour le Climat, comité de 150 citoyens tirés au hasard (par numéro de téléphone). Les citoyens tirés allait de 16 à 80 ans.

L'objectif était de trouver des propositions de lois afin de réduire de 40% les émissions de gaz a effet de serre d'ici 2030, dans un esprit de justice sociale. Le gouvernement ayant promis de soumettre ces lois sans modifications au vote par référendum ou au Parlement.

La convention a abouti à 149 propositions, vues par beaucoup d'experts comme très pertinentes. Cependant, aucun référendum n'a été tenu a ce jour et les propositions de lois on été fortement amoindrie par rapport aux propositions de la convention.

# 2023-03-30 Le Syndicat des Copropriétaires

#appartement 

Toute copropriété forme automatiquement un syndicat des copropriétaires (=ensemble des copropriétaires). Un syndic (!=syndicat),professionnel ou non, représente le syndicat (dans notre cas: Espace Immobilier Lyonnais). Le syndic s'occupe de :
- la gestion financière (budget prévisionnel, tenue des comptes, compte bancaire au nom du syndicat)
- la gestion administrative (réalisation de la fiche synthétique de copropriété, tenir a jour la liste des copropriétaires, représenter le syndicat en justice)
- Organise une AG au moins 1 fois par an, où le syndicat vote différentes propositions

En plus du syndic, un conseil syndical (élu parmi les copropriétaires) non-rémunéré peut être formé pour coordonner les relations entre le syndic et le syndicat. Le conseil syndical peut être consulté et assure le contrôle du syndic.

# 2023-03-31 Eco-village Emmaüs Lescar Pau

#france #ecologie 

Il existe un eco-village près de Pau : l'eco-village Emmaüs Lescar Pau. 

Créée en 1982 à Pau dans une ancienne usine de textile, la communauté Emmaüs déménage par la suite à Lescar, puis, en 1995 fonde le village Emmaüs. Le village compte aujourd'hui prêt de 150 habitants.

Emmaüs est un mouvement créé par l'Abbé Pierre en 1949. Les "compagnons" ont petit a petit rejoins le mouvement, initialement pour aider à reloger des familles et pour transformer et réutiliser des déchets.

# 2023-04-01 Code de Hamming

#code #algorithme 

Le Code de Hamming est une manière d'encoder des données binaires pour qu'elles soient resiliantes aux erreurs. Si un bit est inversé (flipped), alors l'encodage permet de détecter l'erreur ET de la corriger, afin de garder l'intégralité de la data.

Le Code de Hamming est une sorte d'évolution du "parity check", où un bit est réservé pour la détection d'erreur. Si, dans la data, le nombre de bits a 1 est pair, alors le parity bit doit être mis à 0, sinon il est mis a 1. Cela permet de savoir si un bit a été flipped, mais pas lequel.

Avec le Code de Hamming sur 16 bits :
- 11 bits sont utilisés pour la donnée a transmettre 
- 4 bits sont utilisés pour la détection d'erreur
- le dernier bit est utilisé pour le parity check du tout (permettant de détecter mais pas de corriger un 2eme bit inversé)

3Blue1Brown : https://youtu.be/X8jsijhllIA

# 2023-04-02 Le Chiffrement AES

#algorithme #cryptographie #gentlehumanclub

Le chiffrement AES (Advanced Encryption Standard) est un algorithme de chiffrement inventé par Rijndael. Cet algorithme remporta le concours AES lancé par NIST en 1997 pour trouver l'algorithme de chiffrement utilisé par le gouvernement américain.

Il s'agit d'un chiffrement symétrique (la clé d'encryption permet également de décrypter) et utilisé une clé de 128bit, 192bit ou 256bit, selon le besoin de sécurité.

Dans sa version 128bit, l'algorithme encrypted des blocs de 128bit de data en performant 10 rounds des operations suivantes :
- SubBytes (remplace chaque byte en utilisant une lookup table)
- ShiftRows (shift les lignes)
- MixColumns (multiplie chaque colonne par une matrice donnée)
- AddRoundKey (XOR la data avec la round key, générée a partir de la key)

Cet algorithme est implémenté directement dans les processeurs intels et est extrêmement efficace, permettant de lire et écrire des fichiers chiffrés sans délai.

Numberphile : https://youtu.be/O4xNJsjtN6E

# 2023-04-03 Le Ray Tracing

#code #graphics #algorithme 

Le Ray Tracing est une technique qui permet de réaliser des rendus de modèles 3d, en emulant le parcours de la lumière et son interaction avec l'environnement. Il consiste a envoyer des rayons depuis un point origine, la caméra, et de le faire diffuser, refléter, réfracteur (ou autre) sur les matériaux rencontrés. 

Un matériau qui diffuse verra les rayons lumineux être envoyés dans tous les directions après le contact, alors qu'un matériau qui reflète envoie dans une seule direction.

Pour réaliser l'effet "depth of field", où les objets trop proche où trop éloignés sont flou, il faut envoyer des rayons depuis plusieurs origines, toutes convergentes vers un point du plan focal.

Vidéo de Sabastan Lague: https://youtu.be/Qz0KTGYJtUk

# 2023-04-04 Tendons et ligaments

#biologie #escalade 

Le tendon est le prolongement du muscle et le relie aux os.

Un ligament relie deux os dans une articulation (par exemple, les ligaments croisés du genou)

# 2023-04-05 10 tonnes de CO2eq par an

#ecologie 

L'empreinte moyenne d'un français est de 10 tonnes de CO2 équivalent émis dans l'atmosphère par an.

Ce chiffre prendre un compte les importations et exportations du pays.

Pour atteindre l'objectif de l'accord de Paris, il faut diminuer ce chiffre a 2 tonnes, nécessaire pour que la température mondiale ne dépasse pas 2°C d'augmentation.

# 2023-04-06 Les apps Slack

#tools #web

Il est possible de créer des apps Slack très simplement, avec la librairie bolt (JS). l'App peut être self hosted ou utiliser l'infrastructure gérée par Slack.

Une app Slack peut réagir a partir de plusieurs triggers (events, webhooks, scheduled) et exécuter une fonction ou un workflow complet. L'app peut publier des messages et même des modales avec des éléments interactifs (ex: boutons)

# 2023-04-07 La compensation Carbone Volontaire

#france #ecologie 

Pour respecter l'Accord de Paris, les entreprises doivent atteindre le "net zéro émissions" d'ici 2050. Pour cela, elles doivent :
1. Réduire leurs émissions de CO2 de 80%
2. Séquestrer elles-même les 20% restants, c'est à dire re-capter le CO2 émis via la reforestation, la gestion des sols, voire la captation de CO2 via de nouvelles technologies (si cela voit le jour)

Note: Le terme de "neutralité carbone" ne doit plus être utilisé aujourd'hui par les entreprises, car trompeur : il prennait en compte la "compensation carbone". La neutralité carbone a du sens uniquement à l'échelle de la planète.

Aujourd'hui, presque aucune entreprise n'a encore atteint le "Net Zéro Émission". Cependant, certaines peuvent faire de la Compensation Carbone Volontaire, à savoir acheter des "tonnes de CO2 evités" via divers projets (reforestation, gestion des sols, remploi de matériaux). Cela prend la forme de Crédits Carbone, validés par dans labels comme le "Label bas carbone" en France, qui s'assure que le projet en question est sérieux. Ces crédits sont donc acheté par l'entreprise, ce qui permet également de financer le projet.

Il existe des "méthodologies", validées par le ministère de la transition écologique, qui permettent d'évaluer la quantité d'émission de carbone evité pour chaque type de projet. Certains projets, comme l'isolation d'un immeuble d'habitation par exemple, ne peuvent pas faire financer les tonnes de CO2 evités par l'isolation, car ces rénovations sont déjà requises par la loi.

# 2023-04-08 Les noeuds d'escalade

#escalade #manip

Il existe différents types de nœuds en escalade.

D'abord, les noeuds d'encordement, qui permettent d'attacher la corde a son baudrier
- le noeud de huit
- le noeud de chaise

Les noeuds d'assurage et de blocage 
- le noeud d'arrêt 
- le cabestan
- le demi-cabestan
- le noeud de mule
- le machard (variantes: le français, le machard tressé)
- le prusik

Les noeuds de jonction, permettant de rabouter deux cordes entre elles.
- le noeud de pêcheur double 
- le noeud de poing

Autre noeud utiles:
- le noeud en double huit
- la tête d'alouette

# 2023-04-09 Norvegian Allemansretten

#voyage #europe #geographie 

En Norvège, l'accès à la nature est inscrit dans la loi, c'est la loi "Norvegian Allemansretten". Il est donc permis de bivouaquer sur des terrains publics et privés, y compris dans les parcs naturels.

Cependant certaines restrictions s'appliquent : il ne faut pas rester 2 jours au même endroit sans demander l'autorisation, et le campement doit se trouver à plus de 150m de toute habitation.

Il est autorisé d'allumer un feu, sauf entre le 15 avril et le 15 septembre, ou bien dans certaines zones protégées.

Ma cueillette de bais et champignons est autorisée, ainsi que la pêche en eau salée, tant que cette pêche est utilisée pour la consommation personnelle.

Il est évidemment demandé de laisser les emplacements aussi propre qu'à son arrivée. De même la vidange des vans doit être réalisée dans certains emplacements prévus pour cela, indiqués par des panneaux.

# 2023-04-10 L'OTAN

#geographie #histoire 

L'Organisation du Traité de l'Atlantique Nord (OTAN ou NATO en anglais) est une alliance politique et militaire entre les pays des l'Amérique du nord et d'Europe.

Crée en 1949 via le traité de Washington, l'OTAN avait pour objectif initial de contrer la menace que représentait l'union soviétique. Elle compte aujourd'hui 31 états membres.

Les membres de l'OTAN adhérent notamment au principe de défense collective, selon lequel une attaque armée contre un pays membre est considéré comme une attaque contre tous. Ce principe n'a été invoqué qu'une unique fois: à la suite du 11 septembre 2001 par les Etats Unis. 

L'OTAN possède des effectifs armés et un commandant militaire, ainsi qu'un budget issue de ses état membres.

Fortement critiquée par la France et les Etats Unis ces dernieres années, l'OTAN retrouve sa raison d'être au vue de la guerre en Ukraine. La Finlande est le dernier pays a avoir rejoint l'alliance en avril 2023, son adhésion s'étant accélérée au vu de la menace que représente la Russie.

# 2023-04-11 Le Kazakhstan

#geographie

Le Kazakhstan est un pays d'Asie Centrale, peuplé de 19M d'habitants. Le mot "Kazakh" est un dérivé de "qaz" en turc qui signifie se déplacer et le suffixe "stan" signifie territoire en Perse : c'est donc le pays des nomades.

Composé d'immenses steppes mais également de grandes chaînes montagneuses le Kazakhstan a une superficie particulièrement vaste : la 9e plus grande mondiale, ce qui en fait d'ailleurs une des densité de population la plus faible. Le fleuve Oural, frontière entre l'Europe et l'Asie, parcours le pays pour se jeter dans la mer caspienne.

La religion majoritairement islamique et minoritairement chrétienne. Côté ethnie : 75% de la population est Kazakh et parmie reste 15 est Russe.

Au 13e siècle, le territoire a été conquis par l'Empire Mongol de Genghis Khan. Au 15e siècle, lors de la dislocation de cet empire, a été établi le Kazakh Khanate qui deviendra plus tard le Kazakhstan.

Le pays a fait partie de l'Union soviétique et on peut noter qu'il s'agit de la dernière république soviétique à se détacher de l'union.

# 2023-04-12 Le Baseball

#sport

Le baseball est un sport d'équipe à 9 joueurs par équipe. Le jeu se déroule en 9 manche où chaque équipe prend 'e rôle d'attaquant (avec un batteur) ou de défenseur. Le but du batteur est de frapper la balle, de parcourir les 3 bases et de retourner au marbre pour marquer un point. Cependant s'il est touché par un défenseur avec la balle en main, il est éliminé.

La Major League Baseball est la compétition principale de baseball et existe depuis 1903. Elle est composée de 30 équipe (29 américaines + 1 canadienne).

# 2023-04-13 Les Shaders

#code #graphics #flutter

Un shader est un programme exécuté sur le GPU, et qui permet d'accélérer certaines types de tâches, comme le rendu 3d ou encore des effets de post-processing (comme le blur changement de couleur, de contraste, effet bokeh, light bloom, etc)

Il en existe deux types principaux :
- les Vertex Shaders, exécutés pour chaque vertex (=point dans l'espace) et qui peut modifier des données liées à ce Vertex, tel que sa position, par exemple 
- Les Fragment Shaders (ou Pixel Shaders), exécutés pour chaque pixel et où l'output est la couleur du pixel (entre autre)

Les Shaders sont souvent écrit avec le langage GLSL (pour OpenGL Shader Language), maintenu par Khronos Group.

Il va bientôt être possible de compiler ses propres Fragment Shaders sur Flutter, afin de réaliser de nouveaux effets. Cela est d'ailleurs déjà possible sur le channel `master` de Flutter.

Pour apprendre à écrire des Shaders : https://thebookofshaders.com/

# 2023-04-14 Le WebDAV

#technologie#tools #web

Web-based Distributed Authoring dans Versioning (WebDAV) est un protocol qui permet de lire et d'écrire des fichiers via le web. C'est une extension du protocol HTTP, créée par une Internet Engineering Task Force (IETF) à partir des années 1996.

Le protocol permet à plusieurs utilisateurs, après s'être authentifié, de récupérer, déposer, ou synchroniser des fichiers et dossiers rapidement.

Concrètement, il s'agit de nouveaux verbes HTTPS, tel que "COPY", "LOCK", "MKCOL", "MOVE", et d'autres.

Ce Vault Obsidian est synchronisé via un server WebDAV hébergé sur thalkz.com

https://datatracker.ietf.org/doc/rfc4918/

# 2023-04-15 La remontée sur corde

#escalade #manip

Pour faire de la remontée sur corde, il faut poser deux auto-blocants :
- En haut: un machard avec un noeud français (cf [[2023-04-08 Les noeuds d'escalade]]). Sur ce machard, accrocher un sangle pour se hisser avec le pied.
- En bas: le reverso en mode auto-blocant, accroché au pontet du baudrier.

Pour remonter, alterner le poid sur les autobloquants en remontant l'autre a chaque fois.

# 2023-04-16 Faire un mouflage

#escalade #manip 

Le mouflage est une manip d'escalade qui permet d'aider un second bloqué. Il existe différents mouflages qui multiplient plus ou moins les forces (effet poulie).

Pour faire un mouflage simple, il faut comme par poser un relais, et installer le reverso en mode assurage du second.

A partir de cette situation, il faut installer un machard sur la corde partant du second, y accrocher un mousqueton et refaire passer la corde dans ce mousqueton. En tirant sur cette corde, la force est multipliée par 3 pour tirer le second.

![[Pasted image 20230416185248.png]]

# 2023-04-17 L'architecture de Flutter

#code #flutter 

Flutter is composed of two parts :
- the Flutter Framework, written in Dart
- the Flutter Engine, written in C/C++

Both parts communicate via an abstraction layer, called FlutterWindow. This abstraction layer provides an API to communicate to the device and also allows the Engine to notify the Framework of :
- device events (orientation changes, device settings changes, memory issues, etc)
- user gestures (taps, swipes, etc)
- platform channels
- most importantly: when the engine is ready to render a new frame

Source: https://www.didierboelens.com/blog/en/flutter-internals

# 2023-04-18 Processor Architectures

#code #hardware 

Computers have different processor architectures, which can be categorized based on their instructions set.

Instruction Set Architectures (ISA) define the communication rules between the software and the hardware, and especially the set of instructions that a particular hardware can handle.

There are two main ISAs :
- Clomplex Instruction Set Computers (CISC) include complexe instructions, like multi-step and energy consuming instructions. This is used by most Intel processors.
- Reduced Instruction Set Computers (RISC) have a limited instruction Set, that take only 1 clock cycle to complete.

Here are the main processor architectures available today:
- ARM processors use the RISC architecture. (ARM = Advanced RISC Machine). It exists in 32bits and 64bits
- Intel x86 (or x86-32, or x32) processors use the CISC architecture. It is the 32 bit version of the processor, limited to 4Go of RAM and that is used on older computers
- AMD64 (or x86-64, or x64) is the newer, 64bit version of the instructions set. It is retro-compatible with the 32 bit version.

# 2023-04-19 Les Certificats TLS

#cryptographie #web

Les certificats SSL (Secure Socket Layer) ou TLS (Transport Secure Layer) permettent de sécuriser la connexion entre un client et un site web, en établissant un canal de communication crypté. Ils permettent la communication via HTTPS.

Note : TSL est le nouveau standard qui a remplacé SSL en 2015, après que ce dernier soit compromis par certaines vulnérabilités. Le terme SSL est cependant toujours utilisé de nos jours.

Ces certificats sont délivrés par des Certificate Authority (CA), un tiers de confiance, qui garantissent l'identité du détenteur du certificat (le site web). A minimat, le CA vérifie que le nom de domaine est bel et bien possédé par le site.

Lors de l'établissement de la connexion SSL, 
1. le client demande le certificat au site
2. Le site renvoie ce certificat
3. Le client vérifie que le certificat est signé par un CA reconnu et qu'il s'agit du bon nom de domaine 

Ce processus s'appelle le TLS Handshake (qui inclut le TCP Handshake, d'ailleurs)

![[b83b75dbbf5b7e4be31c8000f91fc1a8.svg]]

Chrome, Firefox, Safari possèdent chacun entre 100 et 150 CA reconnus (environ). Les principaux étant Verisign, Let's Encrypt, Digicert

# 2023-04-20 Les RenderObjects sur Flutter

#code #flutter 

Les RenderObjects jouent un rôle clé dans le rendering d'une application Flutter. Les RenderObjects forment un arbre : le Render Tree.

Note: En Flutter il existe 3 trees :
1. Widget Tree (immutable)
2. Le Element Tree
3. Le Render Tree (mutable)

Les RenderObjects ont plusieurs responsabilités :
- Calculer la mise en page (layout)
- Dessiner l'interface (oainting)
- Détecter les taps (hit testing)
- Accessibilité (sementics)

La plupart des widgets utilisent une sous-classe de RenderObject: les RenderBox, qui utilisent des coordonnées 2D.

Les RenderObjects sont réutilisables d'un frame à l'autre (contrairement aux Widgets). C'est d'ailleurs a cela que sert l'attribut "key" sur les Widgets : permettre de réutiliser le "Element" et le "RenderObject" associé au Widget, même si ce dernier à été déplacé dans de Widget Tree.

# 2023-04-21 Les Mixins en Dart

#code #flutter 

Sur Dart, une classe donnée ne peut `extend` qu'une seule autre classe à la fois. Cependant, il est possible d'utiliser un ou plusieurs `mixins`.

Les Mixins sont déclarés avec le keyword `mixin` (plutôt que `class`) et peuvent implémenter une ou plusieurs méthodes, qui seront donc accessible à la classe qui utilise le mixin.

Note : Les Mixins ne peuvent pas implémenter de constructeur.

Si un mixin override une méthode, il est possible d'utiliser super.laFonction() pour exécuter également la fonction originelle. Cette technique est utilisée par le Frame Flutter pour initialiser tous les Bindings (qui sont des mixins) via la fonction `initInstances()`: chaque invocation commence par `super.initInstances()`, permettant d'appeler la méthode récursivement sur les autres Bindings.

Documentation Dart : https://dart.dev/language/mixins

# 2023-04-22 Le WebAssembly

#web #code #technologie #javascript 

Le WebAssembly est un format de binaire exécutable dans le navigateur web. Plutôt qu'un langage à part entière, le WebAssembly est une "target" de compilation pour différents langages tel que Rust, C++, Go, Zig, etc.

Le WebAssembly est aujourd'hui compatible avec les principaux browser modernes (Chrome, Firefox, Safari, et même Explorer)

Le WebAssembly possède plusieurs caractéristiques intéressantes :
- les performances d'exécution sont bonnes
- la taille de l'exécutable (en binaire, donc) est de plus petite taille qu'en JavaScript
- la sécurité est bonne car Memory-Safe et Sandboxed
- La portabilité, car le même code WA peut être exécuté sur différents navigateur, voir directement sur la machine (via le WASI)

Le WebAssembly n'a pas pour objectif de remplacement le JavaScript, mais plutôt de le compléter pour certaines tâches plus lourdes. Il peut notamment être appelé depuis le JavaScript et inversement.

Note: Le WebAssembly ne se limite pas au web browser : il peut servir à créer des CLI, des serveurs ou autre, notamment en utilisant le WASI (WebAssembly System Interface).

L'équipe #flutter travaille sur l'option de pouvoir compiler une app Flutter en WebAssembly.

https://webassembly.org/

# 2023-04-23 Les Push Notifications sur le web

#code #web #javascript 

Il est possible d'envoyer des Push Notifications sur le web, tout comme sur des applications natives standards.

Les Push Notifications sont la combinaison deux technologies : les Push Messages (envoi d'événements d'un serveur vers un navigateur web) et les Notifications (affichage d'une "notification" par l'OS).

Il est possible d'afficher une notification sans Push Message, et, dans le cas des applis, il est possible de recevoir un Push Message sans afficher de Notification. Ce sont donc deux parties séparées, bien que souvent utilisées ensemble. C'est pourquoi on parle de Push Notification quand les deux sont utilisés ensemble.

Pour envoyer une Push Notification sur le web, il faut :
1. Demander la permission a l'utilisateur via  l'API JavaScript `Notification.requestPermission()`
2. Authentifier l'utilisateur auprès d'un Push Service, ce qui est réalisé via la Push API. Si l'inscription est réussie, le navigateur reçois un objet `PushSubsription`
3. Envoyer les données contenues dans PushSubsription sur le serveur, et les enregistrer dans une database
4. Le webserver n'envoie pas de push notification directement au client, il doit passer par le Push Service, en lui envoyant le contenu de la notifications, ainsi qu'un identifiant du client
5. Côté client, un `service worker` doit être configuré pour recevoir le push message et déclencher une notification, même si l'utilisateur ne se trouve pas actuellement sur le site web.

La push notification ne peut être reçue seulement si le navigateur est ouvert. Si ce n'est pas le cas, la notification a une durée de vie limitée configurable (time-to-live)

Ce processus a été standardisé par une Internet Engineering Task Force : [Generic Event Delivering Using HTTP Push](https://datatracker.ietf.org/doc/html/rfc8030)

# 2023-04-24 Les Splines

#algorithme #graphics #mathematiques 

Les Splines sont des courbes définies par des points de contrôle. Il existe différents types de Splines, qui présentent différentes caractéristiques et qui sont utilisées dans différents contextes, comme dans le gaming (ex: les paths), les vectors graphics (ex: le "pen tool" sur Figma), le design industriel (ex: carrosserie d'une voiture), l'animation (mouvement de caméra), mais aussi à des endroits plus inattendus, comme les gradients de couleur.

Les Splines principales sont :
- Linear Spline (C0): lignes droites entre les points de contrôle 
- Béziers Spline (C0/C1): la courbe passe par 1 point de contrôle sur 3, et les 2 autres donnent la forme a la courbe. (Figma cruve tool)
- Hermite Spline (C1/C1): Définie par des points de la courbe + leur vélocité 
- Catmull-Rom Spline (C1): la vélocité au pont est définie par le point précédent et le point suivant.
- B-Spline (C2): l'accélération est continue, mais la courbe ne passe pas par les points de contrôle. La courbature de la courbe est également continue (jolies reflexions sur la surface)

Une notion centrale aux Splines est la notion de continuité. 
- Si tous les points de la courbe sont liés, alors la courbe est C0 continue
- Si, en plus, la vélocité d'un point qui parcours la courbe est continue, alors la courbe est C1 continue.
- Si l'accélération est continue, alors la courbe est C2 continue

 Cela vient s'ajouter la notion de continuité géométrique :
 - Si la tangente est continue, alors la courbe est G1 continue
 - Si la courbure est continue, alors elle est G2

https://youtu.be/jvPPXbo87ds

# 2023-04-25 L'équilibre de Nash

#mathematiques #economie 

L'équilibre de Nash est une situation en théorie des jeux, où aucun joueur ne peut améliorer sa situation en changeant de stratégie.

Cependant, ce point d'équilibre n'est pas forcément le point optimal. L'exemple du dilheme du prisonnier ou des vendeurs de glaces montrent que, sous certaines conditions, le point optimal ne peut pas être atteint par la seule poursuite des intérêts personnels : une coopération est donc nécessaire.

La théorie de Nash vient directement contredire (ou nuancer) la théorie économique d'Adam Smith selon lequel "La recherche des intérêts particuliers aboutit à l'intérêt général".

La mathematicien John Nash a remporté le Prix Nobel en 1994 pour cette théorie et le film "A Beautiful Mind" a été réalisé sur sa vie (ainsi que sa lutte contre la schizophrénie). Il est décédé en 2015.

# 2023-04-26 OAuth et OpenID connect

#web #code #user-experience 

OAuth 2.0 (Open Authorization) est un standard web permettant a une application donnée (le client) d'accéder à des informations sur une autre application, le "resource server" (le serveur qui détient les ressources, c'est à dire les informations en question). L'utilisateur s'appelle le "ressources owner".

Le OAuth flow se déroule en quelques étapes :
1. Le resource owner (l'utilisateur) veut permette au client d'accéder à des informations qui se trouvent sur un autre service (ex: je veux autoriser Supabase à accéder à mes repositories Github)
2. Le client redirige vers l'autorization server, avec `client_id`, `redirect_url`, `response_type` et `scopes`
3. L'autorization serveur affiche une page pour demander le consentement de l'utilisateur pour le partage des données.
4. Si accepté, l'utilisateur est rédigé sur e client (grace au `redirect_url`), avec un `authorization_code`.
5. Le client recontacte directement l'autorization serveur, avec `l'autorization_code` et un `client_secret`, mis en place au préalable.
6.  L'autorization serveur renvoie un `access_token`, qui pourra désormais être utilisé par le client pour demander la ressource (dans le scope).

OpenId Connect est un surcouche d'OAuth permettant d'accéder en particulier à la ressource "identité de l'utilisateur". Dans ce cas, le resource server est appelé un "identity provider". C'est ce qui est utilisé pour les boutons Facebook ou Google sign in.

Note: Ne pas confondre OAuth (standard) et Auth0 (solution commerciale d'authentification)

# 2023-04-27 Le Crontab

#tools #code

La commande `crontab` permet d'exécuter à intervalles réguliers des commandes ou scripts sur une machine.

Les crons sont configurés en éditant un fichier. Pour éditer ce fichier, il faut run `crontab -e`

Dans ce fichier, il faut indiquer la fréquence sous la forme:
```
minute hour day(month) month day(week)
```

Par exemple, cette ligne exécute my_command tous les jours à 10h00

```
0 10 * * * my_command
```

Cette ligne exécute la commande toutes les 5 minutes

```
*/5 * * * * my_command
```

Pour la syntaxe détaillée, aller sur https://crontab.guru

# 2023-04-28 Compensation de vol annulé

#voyage #lifeprotips 

Lorsqu'un vol est annulé par la compagnie aérienne, en plus du remboursement du billet, une compensation de 250€ à 600€ peut être réclamée, selon la distance du vol. 

L'indemnisation n'est pas valide si :
- l'annulation du vol à été déclarée plus de 14 jours à l'avance 
- un vol de remplacement est proposé avec une arrivée de moins de 4h de retard 
- la compagnie prouve que l'annulation est due à des circonstances exceptionnelles (erruption volcanique, inondations, guerre civile, tremblement de terre)
- les grèves sont des cas particuliers, qui ne donnent pas toujours droit a une compensation

# 2023-04-29 Le Cloud Native

#web #technologie #devops

Le Cloud Native est une philosophie, un ensemble de techniques et de technologies permettant de développer, de déployer et de manager des applications modernes dans le Cloud. "Cloud" désignant une infrastructure privée ou publique (tel que AWS, GCP, Azure) constituée d'un ensemble de machines. "Native" signifie "Développé pour". En pratique, le Cloud Native tourne autour de Docker et Kubernetes, ainsi que des outils et techniques associés.

La "Cloud Native Architecture" regroupe un ensemble de concepts tels que :
- Les Microservices
- Le design d'APIs
- Le Service Mesh (layer qui manage la communication entre plusieurs Microservices)
- La Containerisation

Le "Cloud Native Development" se joue sur le plan culturel, et regroupe les concepts:
- Continuous Integration
- Continuous Delivery
- Le DevOps
- Le Serverless

Cette philisophie est promue entre autre par la "Cloud Native Computing Foundation" (CNCF), une fondation open source, crée en 2015 et issue de la Linux Formation, ayant pour but de promouvoir le Cloud Native ainsi que de développer des applications Cloud Native, tel que Kubernetes, Containerd, Prometheus, etc. L'ensemble des projets "graduated" peut se trouver ici: https://www.cncf.io/projects/

# 2023-04-30 Les droits d'accès des fichiers Unix

#code #unix #cli

Sur les systèmes Unix, les fichiers peuvent être configurés pour limiter les droits d'accès. Les droits d'accès en lecture, écriture et exécution sont défini intependement pour le propriétaire du fichier, le groupe associé puis tous les autres.

La commande `ls -l` affiche pour chaque fichier une indication du type `- rwx rwx rwx` 
- Le `-` initial indique qu'il s'agit d'un fichier. (`d` pour les dossiers)
- Le 1er `rwx` (Read Write eXecute) concerne le propriétaire du fichier 
- le 2eme concerne son groupe
- le 3eme concerne tous les autres 

Il est possible de voir chaque rwx sous forme binaire, puis de le traduire en chiffres. Par exemple: `rw-` devient `110` et donc 6 en décimal.

Avec cette notation :
- `777` = tous les droits pour tous (non recommandé)
- `755` = tous les droits pour le propriétaire, mais seulement lecture et exécution pour les autres 
- `644` = lecture et écriture pour le propriétaire, lecture uniquement pour les autres

Pour les dossiers, rwx a une autre signification :
- r = pouvoir voir le contenu du dossier (avec les, par exemple )
- w = pouvoir créer ou modifier des fichiers dans le dossier
- x = pouvoir entrer dans le dossier (avec `cd`, par exemple)

# 2023-05-01 Le refactoring

#code 

Le Refactoring consiste à améliorer la structure interne du code, **sans modifications le comportement** externe.

Le Refactoring n'est donc pas la même chose que du code rewrite ou restrucuring, de la résolution de bug. Le Refactoring se fait par toutes petites étapes, qui, bout à bout, font toute la différence.

Dans son livre "Refactoring: Improving the Design of Existing Code", Martin Fowler décrit en détail ces petites étapes qui permettent d'améliorer la qualité du code, et donc d'augmenter la lisibilité, la facilité de maintien et la rapidité de développement de nouvelles features.

https://refactoring.com
Résumé non officiel : https://github.com/HugoMatilla/Refactoring-Summary

# 2023-05-02 La Specification OpenAPI

#code #web #tools #open-source 

La OpenAPI Specification (OAS) standardise la description d'API. L'OpenAPI Description document consiste en un fichier YAML (ou JSON) décrivant chaque endpoint, les paramètres attendus ainsi que les réponses revoyées. La spécification est agnostique du langage utilisé, tant que l'API est basé sur HTTP.

Le document OpenAPI, étant un format standard assez répandu, permet de définir proprement l'API et ainsi d'augmenter la lisibilité pour les développeurs. Il existe de nombreux outils se basant sur le document OpenAPI Description, tel que la génération de docs, la génération de code client (SDK), la validation de données, la création de serveurs mocks, le testing automatique des endpoints, et d'autes.

OpenAPI, initialement appelée Swagger, est actuellement dans sa version 3.1. La spécification est soutenue par l'OpenAPI Initiative, qui est une organisation en gouvernance ouverte, sous la Linux Foundation.

Le site web officiel: https://www.openapis.org/

# 2023-05-03 Le GitOps

#code #git #devops

Le GitOps est une variante (ou évolution) du DevOps. Cela consiste à gérer une infrastructure directement depuis git. Pour cela, un repository contient le (ou les) fichiers d'Infrastructure as Code (IaC) et dès qu'une Pull Request (PR) est mergée sur main, le deploiement continue (CD) déclenche les changements d'infrastructure. Le GitOps est généralement utilisé en conjonction avec Kubernetes et un framework d'IaC comme Terraform ou Ansible.

L'avantage du GitOps réside dans le fait qu'il y a une parité 1-1 entre l'état de l'infrastructure actuel et la branche main. Utiliser git permet aussi d'utiliser les PRs comme lieu de collaboration, de facilement revert un changement (via git revert) mais également d'augmenter l'accountability des changements (via git blame).

Les composantes principales du workflow sont:
1) Un repository Git est la seule source de vérité sur la configuration de l'infrastucture et du code
2) Une pipeline de CD est responsable de mettre à jour du build, des testing et du déploiement
3) Un système de monitoring est mis en place pour vérifier l'état de l'infrastructure (performances), mais aussi de vérifier la parité avec le repository git

Article sur GitLab: https://about.gitlab.com/topics/gitops/

# 2023-05-04 ClickHouse, une base de données OLAP

#tools #data #open-source

ClickHouse est une base de données OLAP (OnLine Analytical Processing) et column-based. Cela signifie qu'elle est optimisée pour faire de l'analyse à partir des données. 

- OLAP = OnLine Analytical Database (utilisée pour les analytics)
- OLTP = OnLine Transactional Database (utilisée pour les operations)

ClickHouse peut faire des requêtes très rapidement, grâce au stockage des données par colomnes (plus grande proximité des données sur le disque dur, donc lecture plus rapide). Cela permet également une meilleure compression des données, car les données sont plus semblables.

Cependant, d'importantes limitations sont a noter :
- le support pour des jointure de tables est très limité
- le nombre d'utilisateur concurrent à la database est limité
- Les données ne peuvent pas facilement être supprimées ou modifiées

https://clickhouse.com/

# 2023-05-05 Buckingham Palace

#histoire #europe 

Buckingham Palace est le palais résidentiel de la famille royale anglaise. Il se trouve dans un parc au centre de Londres et est utilisé pour différentes réceptions ou éventuellement, let que la couronnement de Charles III, fils d'Élisabeth II le 6 may 2023.

# 2023-05-06 L'égalité en JavaScript

#code #javascript 

Il y a 3 méthodes pour tester l'égalité de deux variables en JavaScript :
- La loose equality `==`
- La strict equality `===`
- `Object.is()`

`Objet.is()` est presque identique à `===` mais ne gère pas `NaN` et `-0` de la même manière. 

La loose equality converti les valeurs en nombres avant de les comparer.

La strict equality fait la comparaison sans faire de conversation. Autrement dit: si les types sont différents, alors c'est `false`.

Attention ! Ces trois méthodes ne comparent les objets que par référence. Par exemple :
- `{} == {}` est false
- `{} === {}` est false
- `Objet.is({}, {})` est false également 

Il faut aussi noter que dans `if({})`, `{}` est évalué a `true` .

Le #lifeprotips est de toujours utiliser la stricte equality, sauf si il y a une bonne raison d'utiliser la loose equality.

https://developer.mozilla.org/en-US/docs/Web/JavaScript/Equality_comparisons_and_sameness

# 2023-05-07 Le prototype en JavaScript

#code #javascript 

Tous les objets en JavaScript ont une priorité built-in, `__proto__` (généralement appelé prototype). Cette propriété est set automatique à la création d'un objet (même un objet vide comme `{}`). 

Le prototype permet d'hériter méthodes d'autres objets. Tous les objets créés ont comme prototype `Object.prototype`, donnant la méthode `toString()`, par exemple.

Les prototypes peuvent être chaînés. Et donc si l'interpréteur ne trouve pas la methode directement dans l'objet, il remonte la chaîne des prototypes jusqu'à trouver la méthode en question ou arriver au prototype `null`.

De David

# 2023-05-08 La loi de la trivialité (de Parkinson)

#psychologie #management

La Loi de la trivialité est une loi empirique selon laquelle les organisations et les individus accordent plus d'importance à des futilités. Cela s'expliquerait par le fait qu'il est plus facile de débattre sur des choses que l'on connait, alors que pour des sujets compliqués on suppose que les autres savent ce qu'ils font.

Formulé autrement : le temps d'un sujet passé sur un agenda est inversement proportionnel à l'argent en jeu

Le concept a été appliqué aux développement logiciel par Poul-Henning Kamp en 1999, sous le nom de "bike shed effect" ou "bike shedding".

# 2023-05-09 Le chômage en France

#france #finance

A partir de 1er février 2023, les ARE (Aide au Retour à l'Emploi) sont calculés de cette façon (pour les moins de 53ans)

ARE par jour = 12.7 + 0.4 * SJR

Le SJR (Salaire Journalier de Référence) est calculé a partir des revenus brut des 2 dernières années, divisé par 731 (nombre de jours travaillés ou non).

Par exemple, pour un revenu brut de 32k€ par an:

SRJ = 32000/731 = 85.5€
ARE = 12.7 + 0.4 * 85.5 = 46.9€
ARE (par mois) ~= 1400€

A noter que les ARE ne peuvent pas être supérieur à 75% ni être inférieur à 57% du SJR

# 2023-05-10 Le package fs-promises sur Node.js

#code #javascript 

A la place du package la la librairie standard Node.js "fs" (pour FileSystem), il est possible d'importer la version basée sur les `Promise` avec `import * from "fs/promises"`. Avec ce package, la fonction `readFileAsync` renvoie une `Promise`, par exemple.

# 2023-05-11 Web Platform Baseline

#web #tools 

Baseline est un projet visant à augmenter la visibilité sur la compatibilité des navigateurs webs. Baseline permet de savoir clairement si une feature HTML, CSS ou JavaScript est supportée par tous les navigateurs.

Baseline est disponible sur web.dev et sur MDN (Mozilla Developper Network)

Baseline est également soutenu par le W3C WebDX community group.

Le W3C (World Wide Web Consortium) est une communauté internationale, comprenant du staff à temps plein ayant pour objectif de développer les standards du web.

Les standards sont, entre autre :
- HTML5
- CSS
- SVG
- WOFF
- The Semantic Web
- XML
Au total : 1400 spécifications (dont WebAssembly et WebGPU)

# 2023-05-12 Le format APNG (Animated PNG)

#graphics #web #flutter 

Il existe un format pour les PNG animés : .APNG. C'est une extension (non officielle) du format PNG, assez semblable à GIF. Le format est supporté par Chrome, Safari, Opéra et Edge.

APNG a 24bit colors + 24 bits transparency tandis que GIF a seulement 8bit transparency.

Depuis Flutter 3.10, les APNGs sont supportés dans la librairie standard (Image)

# 2023-05-13 Les cellules cancéreuses

#biologie 

Les cellules du corps humain ont un noyau rempli avec de l'ADN (Acide DésoxyriboNucléique), constitué de gènes. Ces gènes sont des instructions pour construire des protéines (par des ribosomes).

L'ADN, étant dupliqué pour chaque cellule, subi régulièrement des mutations aléatoires. Certaines de ces mutations peuvent êtres corrompues, produisant des protéines corrompues. 

Les erreurs de copies de l'ADN sont plus fréquentes chez les fumeurs, en cas d'obésité, mais sont principalement liés à l'âge.

Certaines cellules corrompues peuvent êtres cancereuses. Un cancer est une cellule corrompue de l'organisme qui se duplique da manière incontrôlée. Cela peut se passer pour n'importe quelle cellule dans n'importe quelle partie du corps.

Bien que normalement détruite par le système immunitaire (T cells et Killer cells), le cancer peut se développer et rendre des organes défectueux.

https://youtu.be/zFhYJRqz_xk

# 2023-05-14 Hauteur d'une table

#appartement 

La hauteur moyenne d'une table a manger est de 70cm. Une table haute mesure généralement 90cm, et un bar 110cm.

L'assise de la chaise doit faire entre 25 et 30 cm de moins que la table.

D'après Yann : les tables sont plus autour de 75cm aujourd'hui

# 2023-05-16 Les locations meublées

#appartement 

Les revenus issus de la location de meublés sont soumis à l'impôts sur les revenus (IR) dans la catégorie "Bénéfines Industriels et Commerciaux" (BIC). 

Il existe des cas d'exonération des impôts en cas de location d'une partie seulement du logement principal.

Du point de vue fiscal, une différenciation est faite entre un Loueur en Meublé Professionel (LNP) et un Loueur en Meublé Non Professionnel (LMNP). Le loueur est professionel si les deux conditions sont remplies (sinon, c'est non-professionnel):
1. Les recettes dépassent 23k€ par an
2. Les recettes dépassent le salaire

Pour les LMNP, si les revenus ne dépassent 72 600€, le régime Micro s'applique. Dans ce cas, un abbattement forfaitaire de 50% s'applique. Autrement dit: seulement la moitié est imposé.

Les loueurs professionnel ou non doivent se déclarer au Guichet des Formalités des Entreprises (GFE) au plus tard 15 jours après le début de l'activité: [formalites.entreprises.gouv.fr](https://formalites.entreprises.gouv.fr/). Cela permet de reçevoit un numéro de SIRET.

Les loueurs en meublés (pro ou non) sont soumis aux Cotisation Foncières des Entreprises (CFE). Il faut souscrire à la "déclaration initiale de CFE" la 1ere année de l'activité.

---

Une location ne peut être considérée comme meublé que si elle possède les éléments de mobiler suivants:
- Lit + couettes
- Rideaux ou volets aux fenêtes dans les chambres
- Plaques de cuisson
- Four ou micro-ondes
- Frigo ou congélateur
- Vaisselle
- Table et chaises
- Etagères de rangement
- Luminaires
- Matériel d'entretien

https://www.impots.gouv.fr/particulier/location-meublee

# 2023-05-17 La base 58

#cryptographie #blockchain

La base 58 est un set de caractères permettant de compter, au même titre que les chiffres décimaux (base 10), hexadecimaux (base 16) ou encore binaire (base 2).

La base 58 utilise tous les chiffres, les lettres minuscules et majuscules SAUF les chiffres 0 et 1 et les lettres I et O (i et o majuscules). 

Cela permet d'avoir un set de caractères compacte, facile à écrire et qui n'a pas d'ambiguïté à la lecture (grâce aux caractères retirés)

Ces propriétés en font une base parfaite pour des clés de cryptographie et est utilisé notamment pour représenter les clés publique sur Bitcoin et Solana.

# 2023-05-18 Le nonce en cryptographie

#cryptographie 

Un nonce est un chiffre aléatoire envoyé par le serveur au client et que le client renvoie dans sa prochaine requête. Un nonce ne peut être utilisé qu'une seule fois.

Le nonce permet d'éviter une "replay attack", où un attaquant ayant intercepté le message peut le renvoyer au serveur pour répéter l'action. C'est notamment utile pour de l'authentification: sans nonce, un attaquant ayant intercepté les messages d'authentification pourrait simplement les rejouer à l'identique pour s'authentifier)

Le nonce est donc utilisé pour la secret-key authentification et pour la public-key authentification.

Note: Le nonce est aussi utilisé dans l'autre sens: du serveur vers le client.

# 2023-05-19 Les types de cordes

#escalade 

Il existe 3 homologations de cordes : les cordes à simple, les cordes à double et les cordes jumellées.

Les cordes jumellées ne sont quasiment plus utilisés car présentent aujourd'hui peu d'intérêt par rapport aux cordes à double. Elles sont prévues pour être toujours utilisées ensemble (à toujours clipper dans le même mousqueton). Elles sont reconnaissables par le symbole de deux O superposés en bout de corde.

Les cordes à double sont généralement utilisées en grande voie. Prévues pour être utilisés par deux afin de faire des rappels, elles permettent également de grimper en flèche. Les cordes à double sont plus élastiques que les cordes à simple, grâce aux deux brins à l'intérieur de chaque corde. Elles sont reconnaissables par un 1/2 marqué en bout de corde.

Les cordes à simple n'ont qu'un seul brin, sont donc plus rigide et plus épaisses que les cordes à double et sont à utiliser seule (ne pas les utiliser à deux corde, car elles risquent de "sécher" un grimpeur qui chute, car plus rigide). Ces cordes sont donc utilisées seules pour la couenne. Elles sont reconnaissables par un 1 marqué en bout de corde.

# 2023-05-20 Faire une rappel sur une corde tonchée

#escalade #manip 

Il existe deux techniques principales pour faire un rappel sur une corde tonchée (qui a un point très fragilisé)

1. Ne descendre en rappel que sur une seule corde (la saine) et utiliser la corde tonchée pour rappeler les cordes. Dans ce cas : faire un noeud de 8 sur la corde saine et clipper avec un mousqueton sur l'autre. Cela évite complétement de soliciter la corde tonchée lors de la descente.
2. Faire un noeud de papillon au niveau de la coupure. Puis utiliser deux machards pour passer au dessus du noeud lors de la descente. Pour cela : faire un rappel avec un machard tressé (idéalement avec une sangle) et avec le machard au dessus du Reverso, pour passer plus facilement le noeud.

# 2023-05-21 Le GR5

#france #voyage 

Le GR5, aussi appelé la Grande Traversée des Alpes, part du lac Léman à Thonon les Bains et traverse les Alpes jusqu'à la mer méditerranée, à Nice.

Il faut compter entre 30 et 40 jours de marche en général pour effectue les 600+ km de trajet, avec plusieurs journées avec un dénivelé positif ou négatif à 1000+ mètres.

Le trajet traverse les pre-alpes Savoyardes, le massif du Mont Blanc, la Tarentaise, le massif de la Vanoise, le massif du Thabor puis le Briançonnais, suivi de Queyras, Ubaye pour enfin arriver dans le Mercantour.

# 2023-05-22 L'oxygène de l'océan

#ecologie #biologie 

Presque tout l'oxygène respirable sur la planète ne provient pas des forêts, mais des océans. Ce sont de micro-organismes marins tels que les cyanobacteries et les micro-algues planctoniques qui le produisent, avec la photosynthèse.

En plus de ce rôle de fournisseur d'oxygène, les océans régulent le climat. Entre autre, ils absorbent près de 30% des émissions de gaz à effet de serre. Cependant, cet excès de CO2 absorbé diminue le pH de l'eau. Autrement dit: les océans s'acidifient. Cela déséquilibre la biosphère marine et empêche le plancton et les coraux de se renouveler.

# 2023-05-23 Le Plan d'Epagne en Actions

#france #finance 

Le Plan d'Epargne en Actions (PEA) est un produit d'épargne réglementé donnant une exonération d'impôts (il reste cependant les charges sociales), sous certaines conditions.

Un PEA ne permet d'investir seulement dans des entreprises européennes et possède un plafond (150k€ pour un PEA classique)

Il existe deux types de PEA : le PEA classique et le PEA-PME, ce dernier étant dédié aux PME (Petites et Moyennes Entreprises) et ETI (Entreprises de Taille Intermédiaire)

Il y a une limite d'un PEA par personne et 2 par foyer fiscal. Chacun peut avoir son PEA dans un couple Pacsé.

Retirer de l'argent avant 5 ans entraine la fermeture du compte et est imposé à 30% dont 17,2 de charges sociales (sauf conditions exceptionnelles). Un retrait après 5ans n'est imposé que de 17.2% (charges sociales) et n'entraîne pas la fermeture.

# 2023-05-24 Le label AB

#france #ecologie #nutrition 

Le label AB (Agriculture Biologique) certifie que la production des aliments n'utilise pas de produits de synthèse (pesticides, engrais) et ne contient pas plus que de fiables traces d'OGM. Le label s'est aligné sur le label européen en 2009, aussi appelé l'Eurofeuille.

Ce label est critiqué car jugé trop laxiste:
- il n'interdit pas les fermes-usines (agriculture industrielle)
- il n'incite pas au circuit court
- il ne tient pas compte des problématiques liées au commerce équitable (rémunération des producteurs)

Les labels Bio Équitable, Bio Français Équitable, Demeter et Bio Cohérence sont quant à eux plus stricte.

# 2023-05-25 Les engrais azotés

#ecologie #nutrition 

L'azote est un des composants essentiels des protéines produites par les plantes lors de leur croissance. Naturellement présente dans le sol, il existe aujourd'hui des engrais synthétiques permettant d'augmenter le rendement d'une parcelle de terre.

Le procédé Haber-Bosch, mis au point vers l'an 1909 permet la synthetisation d'ammoniac à partir de l'azote présent dans l'air et d'hydrogène (avec consommation d'énergie). L'ammoniaque est ensuite combiné avec d'autres éléments pour former l'ammonitrate, l'urée, les sulfates d'ammoniac.

Dans le monde, 130 millions de tonnes sont produites chaque année. Cette utilisation abondante entraîne des dégâts majeurs sur la biodiversité, sur la santé animale et humaine) et plus généralement sur l'environnement. [[2023-05-24 Le label AB]] entre autre, garantie une culture qui n'utilise pas d'engrais azotés.

# 2023-05-26 Strict Mode in JavaScript

#code #javascript 

Strict Mode en JavaScript est un mode qu'il est possible d'activer, soit en utilisant les modules ECMAScript, soit en précisant "use strict"; dans le fichier. Le mode par défaut est appelé (non officiellement) Sloppy mode.

Strict mode a une sémantique et un comportement différent de Sloppy mode : ce n'est pas juste un subset des fonctionnalités. Strict Mode permet :
1. De throw plutôt que d'avoir des erreurs silencieuse dans certains cas
2. Permet de meilleurs optimisations de l'interpréteur
3. Prépare aux évolutions de syntaxe d'ECMAScript

# 2023-05-27 Pointure de pieds

*File not found.*

# 2023-05-28 Les plus gros festivals en Europe

#europe #musique

Les plus gros festivals en Europe sont :
- Primavera Sound Barcelona, Espagne (début juin)
- Rosklide Festival à Rosklide au Danemark (fin juin)
- Glastonbury à Pilton au Royaume-Unis (fin juin) 
- Rolling Loud à Portimao au Portugal (début juillet )
- L'Exit à Novo Sad en Serbie (début juillet)
- Awakenings à Hilvarenbeek aux Pays bas (début juillet)
- Tomorrowland à Boom en Belgique (fin juillet)
- le Sziget à Budapest en Hongrie (mi aout)
- Creamfields à Daresbury en Angleterre (fin août)
- Rock en Scène à Paris (fin août)

# 2023-05-29 La journée de solidarité

#france #calendrier 

La journée de solidarité est un jour férié travaillé. Cela signifie que le jour doit être travaillé comme n'importe quel autre jour, mais le salaire de la journée est reversé directement à l'état pour financer des actions pour les personnes âgées ou handicapées.

La journée de solidarité tombe le lundi de la Pentecôte (fin mai), mais peut être déplacé à un autre jour férié (sauf le 1er mai)

# 2023-05-30 Les licences de software

#open-source #code #unix 

Les licences permettent de définir les conditions d'utilisation, de modification et de distribution d'une codebase. Les licences sont plus ou moins permissives.

- La licence GPL (GNU General Public License) est en réalité une famille de licences. Les principales versions sont GPLv3, LGPLv2.1 et GPLv2. C'est une licence Copyleft : toute modification du software doit être publiée sous la même licence GPL (ne peut donc pas être rendu propriétaire). Cette licence est utilisée par le kernel Linux, GNU, Wording et Gît.
- La licence MIT est très permissive. Le code peut être utilisé dans du code propriétaire sans être redistribué. Il faut uniquement donner une copie de la license MIT dans le projet. C'est utilisé par Nodejs, Angular, React.
- La license Apache est un entre deux entre MIT et GPL. N'est pas une licence Copyleft. C'est utilisé par Android, Kubernetes.
- La licence BSD ressemble à MIT, à la différence qu'il n'est pas permis d'utiliser le nom des créateurs du software pour promouvoir la modification.
- Les licences Creative Common est une famille de licences pour tout type de travail créatif.

# 2023-05-31 Climatisation et Pompe à Chaleur

#appartement #ecologie 

Une climatisation utilise un fluide caloporteur pour transporter de la chaleur de l'intérieur vers l'extérieur d'un logement. Une pompe à chaleur fait l'inverse. Il existe des clims réversibles et des pompes à chaleur réversible. Ce n'est pas clairement dit sur internet, mais il s'agit de la même chose.

# 2023-06-01 La commune de Vesc

#geographie #france 

La commune de Vesc se situe dans le département de la Drôme (26) en Auvergne Rhône Alpes. Un des produits locaux est le fromage Picodon.

A vol d'oiseau, Vesc se trouve à une quinzaine de kilomètres de Saou et à une trentaine de kilomètres de Crest. La grande ville la plus proche est Montélimar.

Le domaine de Damian, où se marient Thomas et Marlène, se trouve dans ce village.

# 2023-06-02 Validité des billets TER

#france #lifeprotips

La durée de validité d'un billet de TER est de 1 jour. Il est donc possible de prendre n'importe quel TER dans la journée pour un billet donné.

Les conditions de remboursement varient selon les régions. Généralement le remboursement est gratuit ou coute 10% de prix du billet, jusqu'à J-1.

# 2023-06-03 Électricité de France

#france #energie 

Électricité de France (EDF) est une une entreprise détenue a 90% par l'état français, gérant la production et la fourniture d'électricité en France.

Créé après la 2eme guerre mondiale, EDF était une entreprise publique qui est privatisée en 2004 puis nationalisée de nouveau en 2023.

Son résultat net est de -13Mrd d'€ en 2022, malgré un chiffre d'affaires en hausse.

La production d'énergie est principalement nucléaire (77%). Le reste étant constitué d'hydraulique et renouvelables (10%), gaz (8%) et thermique hors gaz (5%).

Les filiales principales sont 
- Enedis (distribution d'électricité en France)
- RTE (Réseau de transport d'électricité en France)
- Dalkia (services énergétiques et production décentralisée)
- EDF Renouvelables

# 2023-06-04 Stefan Zweig

#littérature #histoire 

Stefan Zweig est un auteur juif Autrichien né en 1881 à Vienne et mort, suicidé en 1942 au Brésil. Il a entretenu une correspondance avec Freud.

Il quitte l'Autriche en raison de la montée du nazisme pour s'expatrier à Londres et au Brésil. Il écrivit des biographies (Marie-Antoinette, Joseph Fouché, Marie Stuart) et des fictions (Amok, La Pitié dangereuse, Le joueur d'échecs, La confusion des sentiments

# 2023-06-05 Daniel Balavoine

#musique 

Auteur, compositeur et interprète français, Daniel Balavoine est né en 1952 dans le département de l'Orne (61), en Normandie. Il devient connu à la fin des années 70, avec notamment son album "Le Chanteur". Il joue Jonny Rockfort dans Starmania (Michel Berger et Luc Plamondon).

Il meurt en 1986 d'un accident d'hélicoptère en survolant la course Paris-Dakar.

# 2023-06-06 Le théorème de Fermat Wiles

#mathematiques #histoire 

La conjecture de Fermat (aussi appelé "le dernier théorème de Fermat" ou "le grand théorème de Fermat") a été écrit dans la marge d'un livre par Pierre de Fermat. Il écrivit également qu'il avait une preuve remarquable, mais pas assez de place pour la rédiger dans la marge.

La conjecture est la suivante : Il n'existe pas d'entiers strictement positifs x, y et z tel que x^n + y^n = z^n pour n > 2.

Les mathématiciens du monde entier ont tenté d'apporter un preuve à cette conjoncture, pendant plus de 2 siècles, jusqu'à ce que le mathématicien Andrew Wiles en apporte la preuve en 1993, après 7 ans de travail acharné. La conjecture s'appelle désormais : Théorème de Fermat Wiles.

# 2023-06-07 L'infarctus

#biologie 

L'infarctus, aussi appelé Crise Cardiaque, est une obstruction d'une artère coronaire, artère alimentant en oxygène le muscle du coeur : le myocarde.

Cette obstruction est dûe à l'accumulation de cholestérol (graisse) en plaques diminuant le diamètre de l'artère et où un caillot de sang peut venir obstruer complément le passage.

Chaque année, 80k infarctus se produisent en France (environ 1 personne sur 1000), et cela entraîne la mort en moins d'une heure dans 10% des cas (selon l'emplacement de L'infarctus, entre autre)

Les risques d'infarctus sont plus élevés chez les hommes et augmentent avec l'âge, le tabagisme, l'alcool, le surpoids, la forte [[2023-06-15 La tension artérielle|tension artérielle]], le haut taux de cholestérol, le diabète mais aussi les antécédents familiaux.

# 2023-06-08 Kill un process via son port

#cli 

La commande `lsof` liste tous les fichiers actuellement ouvert sur le système. Avec cette commande, il est possible de trouver le PID (Process ID) qui écoute sur un port donné:

```
lsof -i :PORT
```

Une fois le PID récupéré, pour terminer le process, la commande kill peut être utilisée :

```
kill PID
```

- Par défaut, `kill` envoie un `SIGTERM` (terminaison "normale" du process)
- Avec `kill -2 PID`, c'est un `SIGINT` (le même signal qu'avec Ctrl + C) 
- Avec `kill -9`, c'est un `SIGKILL` (terminaison forcée du process)

# 2023-06-09 Atelier 2tonnes

#ecologie 

L'atelier 2tonnes est un atelier de 3h à faire en groupe avec un animateur et avec l'objectif de répondre à la question : "comment faire pour réduire ses émissions de CO2 d'ici 2050".

Depuis l'accord de Paris de 2015, les pays du monde entier se sont engagés à ne pas dépasser la limite des +2 degrés par rapport à l'époque pré-indistrielle. En France cela signifie que les émissions de CO2eq par habitant doivent passer de 10 tonnes en moyenne à 2 tonnes d'ici 2050.

L'atelier a été créé par François Laugier (Mars) en 2019 et compte d'50000 participation en début 2023.

# 2023-06-10 Les types de roche

#escalade #nature

Il existe 4 types de roches principales en montagne :

+ Le calcaire, roche sédimentaire constituée de sédiments d'animaux marins: coquilles, squelettes. Ex : Pre-Alpes (Chartreuse, Vercors)
+ Le granite, roche qui est une roche plutonique. Elle est formée par le refroidissement lent d'une bulle de magma. Il contient des cristaux de quartz, mica et feldspath. Ex: Mont Blanc, Yosemite, Trango (Pakistan)
+ Le gneiss, roche métamorphique composée de quartz, mica, de feldspath et d'autres éléments non visibles a l'oeil nu. Se forme sous haute pression et haute température. Elle est structurée en feuillets de de qq centimètres à qq décimètres. Ex: Valais (Suisse), sommet des Écrins (Meije, Olan), K2, Nanga Parat.
+ Le schiste, composé de roche sédimentaire et de cristaux. Il est formée sous la pression des plissements de la croûte terrestre. Ex: Massif du Queyras.

# 2023-06-11 Saou

#france #escalade 

Saoû est un village de 500 habitants dans la Drôme (26) situé à proximité de grande voies et couennes. Le rocher est en calcaire.

Avec le CAF, nous avons grimpé dans le secteur "Aiguille de la Tour"

# 2023-06-12 La NEF

#economie #ecologie 

La NEF est une banque éthique française, basée à Vaulx-en-Velin. Elle porte des valeurs de transparence, du collectif et de l'éthique sociale et écologique. Depuis 1988, la NEF accorde des prêts à des projets à impact écologique, social ou culturel. Elle se structure sous forme d'une association où les sociétaires on chacun une voix et décident ensemble du futur de la banque. Chaque année, une rapport est publié sur les investissements. Les projets sont généralement lié au commerce équitable, à la filière bio, aux énergies renouvelables, au vrac, aux projets de méthanisation, à la mobilité douce et aux ressourceries.

La NEF à investi dans:
- Biocoop
- Veja
- Ethiquable
- Enercoop
- New Tree
- Les Arpenteurs (Lyon)
- Cyclable (Lyon)

# 2023-06-13 Wifi 2.4GHz et 5GHz

#technologie #appartement 

La connection Wifi existe sur deux bandes: 2.4GHz et 5GHz. La bande 2.4GHz est la fréquence utilisée à l'origine et est donc compatible, même avec les machines les plus anciennes. La bande 5GHz a été ajoutée pour palier au problème d'interférences de la bande 2.4GHz.

Ce problème d'interférences provient du fait que 
1. la bande de fréquences est utilisé par d'autres appareils életroniques (micro-ondes, bluetooth)
2. Il n'existe que 11 channels pour le Wifi en 2.4GHz, avec seulement 3 channels qui ne se chevauchent pas
3. Le signal est de plus long portée

Le Wifi en 5GHz a une bande passante 3x supérieure et est moins vulnérable aux interférences (plus de channels sans chevauchement). Cependant, la portée est plus faible et le signal traverse mal les murs.

# 2023-06-14 Les sophismes

#philosophie

Les sophismes sont des raisonnements faux malgré une apparence de vérité. Ils sont généralement difficiles à réfuter.

Le terme provient de la Grèce Antique, au 5e siècle BC, prenant sa racine dans le mot sophia: la sagesse. Les sophistes sont des orateurs dont le métier est de convaincre le plus grand nombre. Fortement critiqués par Platon pour leur absence de recherche de vérité, le terme est aujourd'hui péjoratif.

Certains sophismes bien connus sont souvent utilisés aujourd'hui
- l'appel à la popularité : "Tout le monde le dit, c'est bien connu"
- les généralisations hâtives : Généralisation d'un cas à partir d'un exemple personnel
- la pente glissante : "si on fait ça, alors demain on fera ça" (même si ça n'a riya voir)
- l'argument d'autorité : "telle personne pense ça, donc c'est vrai"
- la caricature : exagérer la position de l'adversaire pour la rendre plus critiquable 
- le faux dilheme : "la France tu l'aimes ou tu la quitte"
- l'attaque contre la personne : "Mouai, comment croire quelqu'un qui a quitté à 16 ans"

# 2023-06-15 La tension artérielle

#biologie #santé

La tension artérielle est la pression exercée par le sang sur les parois des artères, en mm de mercure (mmHg). Une tension trop forte ou trop faible peuvent être problématique.

Lors d'une mesure de tension, deux valeurs sont relevées: la pression systolique mesure la pression maximale exercée par le coeur lors de la contraction du ventricule gauche. La pression diastolique indique la pression lors du relâchement du coeur.

Les tensions sont catégories :
- faible si inférieur à 100/60
- optimale jusqu'à 120/80
- normale jusqu'à 130/85
- normale-haute jusqu'à 140/90
- hypertension légère jusqu'à 160/100
- hypertension modérée jusqu'à 180/110
- hypertension sévère ensuite

# 2023-06-16 Fontainebleau

#escalade #france 

Fontainebleau est un spot de grimpe mondialement connu pour ses blocs de grès, avec plus de 25k voies sur 200 secteurs, les plus connus étant Cuvier, Roche aux Sabots et les Trois Pignons

Dans les années 1910, les alpinistes français commencent à utiliser Fontainebleau comme entraînement.

Dans les années qui suivent, des voies de plus en plus difficiles sont ouvertes. Comme en 1935, où Pierre Alain realise l'Angle Allain (5c) au Cuvier Rempart grâce à sa nouvelle invention : les chaussons d'escalade PA. EN1946, René Ferlet réalise la Marie-Rose au Cuvier, premier 6a de Fontainebleau.

En 1962 est créé le CoSiRoc (Comité de Défense des Rochers d'escalade), organisme de protection de Fontainebleau, en collaboration avec l'ONF (Office National des Forets)

La forêt de Fontainebleau est classée "réserve mondiale de la Biosphère" par l'UNESCO.

# 2023-06-17 Le rayon de la terre

#geographie

Le rayon de la terre est de 6300km

# 2023-06-18 Les corbeaux et les corneilles

#nature 

Le corbeau (corbeau freux et grand corbeau) et la corneille sont tous deux de la famille des Corvidés, de pelage noir et peuvent être observés en France. Cependant de nombreuses différences permettent de les distinguer: 

Les corneilles sont plus petites avec un tête plus plate et la queue arrondie au bout. Elles ne peuvent pas planner (elles doivent battre des ailes tout le temps) et au sol se déplacent en sautant.

Les corbeaux sont grands et robustes, avec la tête bombée et la queue en losange. Ils peuvent planner et au sol se déplacent en marchant.

Le cri des corneilles est plus long et fâché "raaaaah" et celui des corbeau est plus bref et aigu.

# 2023-06-19 La RGPD

#technologie #europe #legal 

La RGPD (Règlement Général de Protection des Données) est un réglementation européenne entrée en vigueur en 2018 visant à uniformiser la gestion des données personnelles au niveau européen.

Les données personnelles sont des données permettant d'identifier une personne : soit directement (nom, adresse mail) soit par combinaison (adresse+ date de naissance, ou identifiant unique, etc)

Sur ces données, les entreprises ont un devoir de 
- transparence: expliquer pourquoi ces données sont nécessaires, combien de temps elles sont conservées, qui a accès ces données, etc
- minimisation des données collectées : collecter des données non essentielles n'est pas autorisé.
- permettre aux producteurs de gérer leurs données: droits d'accès, de rectification, d'opposition, d'effacement et à la portabilité. 1 mois max de traitement de ces demandes
- sécurité: chiffrement des données, gestion des mots de passe, etc

# 2023-06-21 Le Duomo di Milano

#europe #histoire 

Le Duomo di Milano (Dôme de Milan ou cathédrale de Milan) est une des plus grande église du monde. Dans un style inspiré du gothique, sa construction a débuté au 14e siècle. Elle est dédiée à la nativité sainte Marie.

# 2023-06-22 De Marxisme au Fondamentalisme de Justice Sociale

#histoire #sociologie #economie 

- 1860-1900: Le Marxisme, c'est la lutte des classes entre travailleurs et patrons, qui s'exerce à travers la politique et l'économie. 
- 1930-1970: Critical Theory puis Critical Social Justice, est une généralisation de Critical Theory à tous les conflits opprimant/opprimés: homme/femme, blanc/noir, hétéro/gay, etc. Il faut questionner la narratrice culturelle dominante et promouvoir la vie des groupes opprimés.
- 1990: Intersectionality: l'intersection de groupes opprimés sont d'autant plus désavantagés (ex: femme noire)
- 2000: Social Justice Fundamentalism (Tim Urban), l'oppression est exercée par tout le discours dominant, y compris le langage, les sciences, la morale, etc. Il faut de la discrimination positive afin de retourner à la justice sociale.

# 2023-06-23 Les types de cafés

#nutrition 

Il existe deux familles principales de graines de café : le café Arabica et le Robusta.

Le café Arabica est cultivé principalement en Amérique du sud et en Afrique, dans des conditions particulières (en hauteur). Il est considéré comme meilleur que le Robusta, avec un arôme plus complexe et une pointe d'acidité. Il est également plus cher que le Robusta.

Le Robusta a un goût plus corsé, avec un arrière goût de cacahuète ou de grillé. Il est cultivé en Orient. Il est également deux fois plus caféiné que l'Arabica.

La plupart du café acheté est un mix des deux.

# 2023-06-24 Le christianisme

#religion #histoire 

Le christianisme est une famille de religions rapportant aux paroles et écrits de Jésus Christ. Il existe trois courants principaux: les catholiques, les protestants et les orthodoxes.

Les protestants ne reconnaissent que la Bible comme source d'autorité. Les orthodoxes et les catholiques reconnaissent la tradition orale, les dogmes et doctrines.

Les protestants refuse les saints et ne vouent pas un culte à Marie, contrairement aux catholiques et orthodoxes, chez qui ces icônes sont omniprésents.

Le pape n'est une figure d'autorité que pour l'église catholique. Les prêtres catholiques n'ont pas le droit de se marier, contrairement aux portes orthodoxes. Chez les protestants, les prêtres sont remplacés par des pasteurs, hommes ou femmes, qui peuvent se marier. Ils ne représentent pas le Christ mais ont plus un rôle de soutien à la communauté.

L'église orthodoxe s'est séparée en 1054, lorsque le patriarche de Constantinople, Michel Cérulaire, rompt avec Rome, créant une séparation Occident/Orient.

Le protestantisme commence en 1517, avec le moine Martin Luther en affichant ses 95 thèses contre les indulgences papales.

# 2023-06-25 La foudre

#nature #sciences

Les éclairs se forment dans un certain type de nuages : les cumulonimbus, qui sont d'énormes nuages verticaux. 

Dans ces nuages, des vents ascendant font remonter des particules de glaces, grêlons, qui grossissent pendant l'ascension, car les températures peuvent atteindre les -20°C au sein du nuage.

Lorsque la masse est trop important, le greffon retombe et entre en collision avec ceux qui montrent, en arrachant des élections. 

Au bout d'un certain temps, le haut du cumulonimbus est chargé positivement et le bas négativement. La charge négative en bas du nuage repose les électrons du sol, qui se charge positivement.

Lorsque la différence de potentiel est suffisamment forte pour ioniser l'air, une décharge s'effectue: c'est l'éclair.

# 2023-06-26 L'alphabet grec

#histoire #voyage 

L'alphabet grec se compose de 24 lettres, dont 7 voyelles. Il prend ses origines en -800 avant JC, lorsque les Grecs commercent avec les Phéniciens, qui utilisent cet alphabet. Les Grecs ont ajouté les voyelles.

![[Pasted image 20230626204528.png]]

# 2023-06-27 Égalité et équité

#philosophie #sociologie 

L'égalité, c'est appliquer le même traitement à chacun, quel que soit sa situation actuelle ou son passé.

L'équité prends en compte les inégalités et les cas particuliers et tente de rétablir un équilibre. C'est le concept de justice sociale. Ce concept est notamment exprimé par Aristote, et théorisé plus tard par John Rawls dans son ouvrage Théorie de la Justice.

# 2023-06-28 Les domaines .new

#web #tools 

Les domaines en .new permettent de faire une action rapide, en se rendant directement sur la page de création de action.

Par exemple :
- sheets.new
- docs.new
- slides.new
- linear.new 
- gist.new
- link.new
- glitch.new
- flutter.new

De Seb

# 2023-06-29 Les jeux Olympiques antiques

#histoire #europe #sport 

Les jeux Olympiques ont débuté en 766 avant JC dans le village d'Olympie (1000 habitants). Ces jeux antiques se sont déroulés tous les 4 ans sans interruption pendant plus d'un millénaire. C'est l'empereur romain Théodore I qui y mis fin, tentant d'unifier l'empire sous [[2023-06-24 Le christianisme]]. Son successeur, Théodore II, ordonna même de détruire les monuments.

A Olympie, les jeux antiques ont évolués au fil des années pour atteindre les six épreuves principales : course, pentathlon, course de char, lutte, course de hoplites.

Les vainqueurs retournaient en héro dans leur cité et étaient exempt d'impôts à vie.

Les femmes et esclaves n'étaient pas autorisées à participer ni à assister aux jeux.

L'origine des jeux est attribuée soit à Héraclès (Hercule) soit au roi Pélops (nom à l'origine du Péloponnèse)

Les JO modernes ont repris en 1896 à Athènes, et aujourd'hui encore, la flamme olympique est allumée à Olympie.

# 2023-06-30 Le bail de mobilité

#france #appartement 

En France, un bail [[2023-05-16 Les locations meublées|logement meublé]] est habituellement fixé à 1 an renouvelable (ou de 9 mois si le locataire est étudiant). Il est possible pour un propriétaire de récupérer son logement pour certaines raisons précises, mais cela reste compliqué.

Le bail de mobilité permet de louer son appartement pour une durée de 1 à 10 mois, non renouvelable (cependant, il peut être converti en bail classique).

Le bail de mobilité compte quelques limitations : 
- le logement doit être meublé
- le propriétaire a interdiction de demander un dépôt de garantie
- et plus important, seuls certains locataires sont éligibles: le locataire doit être en formation professionnelle, en études supérieures, en contrat d'apprentissage, en stage, en service civique ou en mutation professionnelle.

# 2023-07-01 L'Acropole d'Athènes

#voyage #histoire 

L'Acropole d'Athènes est un plateau rocheux dans la ville d'Athènes qui a servi de citadelle pendant l'occupation ottomane puis de lieu sacré pendant l'antiquité.

Le sanctuaire était destiné au culte d'Athena et comprenant plusieurs temples dont le Parthénon, l'Erechtheion et le temple d'Athena Nikè.

# 2023-07-02 Définition du racisme

#sociologie 

D'après Tim Urban dans "What's our problem ?", La définition du racisme a évolué ces dernières années. La définition traditionnelle du racisme, comme inscrite sur le site de l'Anti-Diffamation League :

Le racisme est la croyance qu'une race est supérieure ou inférieure à une autre et que les traits sociaux et moraux d'une personne sont prédéterminés par ses caractéristiques biologiques.

Cette définition a été remplacée par :

Le racisme est la marginalisation et/ou l'oppression de personnes de couleur via une hiérarchie socialement construite qui privilégie les personnes blanches.

Par exemple, selon cette deuxième définition, l'audition "à l'aveugle" des orchestres aurait un caractère raciste, car perpétue le status quo où les personnes blanches et asiatiques sont sur-représentées dans les orchestres.

# 2023-07-03 Les Deeplinks

#tools #code 

Les deeplinks permettent de rediriger un utilisateur sur une page précise dans son application.

Par exemple, ce deeplink permet de rédiger vers la page de partage de l'application Lydia : https://lydia-app.com/site/open/share

Sur Android les deeplinks s'appellent "App Link" et nécessite d'être configurés:
1) Ajouter un fichier json sur son site web https://domaine/.well-known/applink.json
2) configurer le AndroiManifest pour accepter des deeplinks

Sur iOS, cela s'appelle des "Universal Links" et il faut
1) Ajouter un fichier sur https://domain/.well-known/apple-app-site-association
2) Configurer l'App depuis XCode pour accepter le deeplink

Une fois la configuration terminée, cliquer sur un lien vers le site redirige sur l'application (et avec un router bien configuré, sur la bonne page de l'application)

# 2023-07-04 L'effet Dunning Krüger

#psychologie 

L'effet Dunning Krüger est un biais cognitif qui pousse les personnes moins qualifiés à sur-estimer leurs compétences. Il s'explique par le fait que ces personnes auraient une difficulté à identifier leur propre incompétence.

Le corollaire est que les personnes plus qualifiées ont tendance à sous estimer leurs compétences, en pensant que ce qui est facile pour eux l'est aussi pour les autres.

# 2023-07-05 L'échelle de la mobilité

#urbanisme

Une manière de voir le partage de l'espace de mobilité urbaine et rurale est d'associer chaque moyen de transport à son niveau de privilège. Du plus privilégié au moins privilégié :

- les voitures, camions et bus (gros véhicules motorisés)
- les scooters et motos
- les vélos et trottinettes électriques
- les piétons

Les moyens les plus privilégiées appliquent une sorte de pression, même passive, sur les autres, car ils sont plus imposant, lourd et rapide. De plus, les infrastructures sont souvent plus avantageuses pour les plus gros véhicules (cas extrême : les US)

Pourtant, à de nombreux endroits comme dans les villes, les mobilités plus douces sont majoritaires et souvent en croissance. 

Pour "re-équilibrer" les forces, deux moyens :
1) (soft) Donner la priorité au véhicule le moins privilégier, le plus souvent possible (sauf si ça implique de piler au milieu de la route et mettre tout le monde en danger bien sûr)
1) (hard) Faire des infrastructures plus justes (piétonniser, créer des pistes cyclables, augmenter la taille des trottoirs, limiter la vitesse sur les routes)

# 2023-07-06 Le permis français en Europe

#europe #legal

Le permis français permet de conduire dans tout pays européen. Il est également valable au Royaume-Uni, à Monaco et en Suisse pour des séjours de courte durée

# 2023-07-07 Le principe de Peter

#management #psychologie 

Le Principe de Peter, énoncé par Laurence J. Peter et Raymond Hull en 1969, est un loi empirique en rapport avec les organisations hiérarchiques.

> Dans une hiérarchie, tout employé a tendance a s'élever à son niveau d'incompétence

Le corollaire étant

> Avec le temps, tout poste sera occupé par un employé incapable de d'en assumer la responsabilité.

D'Adrien

# 2023-07-08 RATP et SNCF

#france 

La RATP (Régie Autonome des Transports Parisiens) est une entreprise publique locale appartenant à la Ville de Paris créée en 1948 pour assurer l'exploitation des transports en commun de Paris.

La RATP exploite
- les 16 métros parisiens
- 9 tramways (sur 12)
- la plupart des lignes de bus
- les RER A et B, en partie 

La RATP a un contrat avec Île De France Mobilités qui est le syndicat des transports d'Ile de France (Stif). Autrement dit, c'est l'autorité organisatrice des transports en Ile de France.

La SNCF (Société Nationale des Chemins de fer Français) est une entreprise publique (re-privatisée en 2020) composée de deux entités: SNCF Réseau (gestionnaire du réseau ferroviaire) et SNCF Voyageurs (exploitation des trains). Elle est active sur l'ensemble du territoire Français mais également à l'étranger.

# 2023-07-09 Départements

#france #geographie 

Il y a 101 départements français.
- De 1 à 95 en métropole (la Corse compte pour 2)
- de 971 à 975 en outre mer

Les départements initiaux ont été créés peu après la révolution française, en 1790 par l'assemblée constituante, afin de remplacer les provinces de France.

Depuis, même si la plupart des départements sont restés identiques, de nombreuses modifications ont eu lieu :
- En 1810, le nombre de départements grimpe à 130 pendant l'Empire puis redescend à 86 après la chute de l'empire en 1815.
- Les Alpes-Maritimes et la Savoie sont rattachés en 1860.
- En 1871, une partie du Bas-Rhin, Haut-Rhin, Moselle, Meurthe et des Vosges sont cédés à l'Allemagne, puis rétrocédés en 1919. 
- Le Territoire de Belfort est créé en 1922 (90).
- Des départements existent en Algérie entre 1848 et 1959. Ils sont supprimés lorsque l'indépendance de l'Algérie est reconnue.
- En 1964, la région parisienne est réorganisée en Paris, Essonne, Yvelines, Hauts-de-Seine, Seine-Saint-Denis, Val de Marne et Val d'Oise.
- La Corse est divisée en 1976.
- Quatre départements d'outre mer sont créés en 1946, puis un 5eme, Mayotte en 2011 porte le total a 101.

# 2023-07-10 Bivouaquer sans sac de couchage

#lifeprotips 

Bivouaquer sans sac de couchage est possible lorsque la température la plus basse de la nuit est de 16°C, mais il faut mettre quelques couches d'habits (pentalon, veste, k-way, chaussettes) et être isolé du sol.

Le sac de couchage reste essentiel jusqu'à des températures minimales nocturnes de 22/23°C.

A noter que la température la plus basse de la nuit arrive juste avant le lever du soleil, donc vers 5h/6h du matin.

# 2023-07-11 Annuler un évènement organisé

#lifeprotips 

Si un évènement organisé ne peut pas se produire (imprévu, météo ou autre), il vaut mieux l'annuler complément plutôt que d'essayer de re-organiser un évènement partiel avec une partie des gens ou avec des conditions qui ne conviennent pas à tous.

Les personnes exclues du nouvel événement seront forcément très frustrées. Annuler complètement est compréhensible par tout le monde.

# 2023-07-12 Le recrutement chez Google

#technologie #algorithme 

Le processus de recrutement chez Google est similaire est assez long. On contient généralement les étapes suivantes :
- Candidature en ligne, avec l'envoi du CV et d'une lettre de motivation
- Screening par téléphone
- Une journée d'entretiens (5 à 6 entretiens dans la journée). Cette journée d'entretiens comporte des entretiens techniques, mais aussi des behavioral interviews
- Résultats et négociations

# 2023-07-13 Les signes d'urgence en montagne

#lifeprotips #nature 

Pour indiquer de loin si on a besoin d'aide ou non (en montagne ou autre), il existe deux signes principaux:

- Pour demander de l'aide, lever les deux bras pour former un Y. "Yes, I need help"
- Sinon lever un seul bras et garder l'autre en bas, ce qui forme une sorte de N. "No, I don't need help"

En Europe, le numéro d'urgence est le 112.

# 2023-07-14 Les Binary Tree

#algorithme #code 

Les Binary Tree sont une structure de données où chaque node contient deux sous branches (appelés `left` et `right`). L'origine du Binary Tree est appelé `root`

Un Binary Tree peut avoir plusieurs propriétés :
- `balalanced`, si tous les étages du tree sont remplis (sauf le dernier, qui peut être rempli que partiellement)
- `completed`, si il est balanced et que les nodes du dernier étage sont le plus a gauche possible 
- `full` (`proper` ou `plane` ou `strict`), si chaque node a 0 ou 2 enfants (pas 1 seul enfant)
- `perfect`, si il est complètement rempli, jusqu'à un certain étage 
- `degenerated`, si tous les nodes n'ont qu'un seul enfant. Dans ce cas, le Binary Free se comporte comme une Linked List

Avec h la hauteur du tree (la plus grande distance entre le `root` et `leaf` le plus éloigné)
- Un perfect Binary Tree possède `2^(h+1) - 1` nodes 
- Un tree avec seulement le root a une hauteur de 0

# 2023-07-15 Insertion Sort

#algorithme #code 

Insertion Sort est un algorithme de tri très simple, qui réalise le sort in-place.

En prenant les valeurs de gauche à droite, la valeur est "shiftée" vers la gauche jusqu'à ce que la slice de gauche soit triée.

Insertion Sort est en O(n^2) en worst case et average case. Dans le best case, le tableau est déjà trié et dans ce cas c'est en O(n)

# 2023-07-16 Abeilles, guepes, frelons et bourdons

#nature 

- Les abeilles sont petites, trapues et velues et d'une couleurs généralement allant vers le brun foncé.
- Les guêpes sont plus grandes que les abeilles, sans poils et jaune vif avec des rayures noires. Les guêpes ont l'abdomen bien séparé du reste du corps
- Le frelons est la plus grande des guêpes en Europe et peut mesurer jusqu'à 35mm. Elle est légèrement plus velue et sa tête plus orangée
- Les bourdons sont assez ronds (trapus) et mesurent environ de 2cm

# 2023-07-17 Merge Sort

#code #algorithme 

Le Merge Sort est un algorithme récursif utilisant la stratégie "divide and conquer". Pour merge-sort un Array, couper cet Array en deux et appeler merge-sort sur chaque moitié. Il faut ensuite "merger" les deux subarrays, qui sont en ordre, en prenant le plus petit élément des deux subarrays jusqu'à ce que les deux soient vide.

# 2023-07-18 Le petit théorème de Fermat

#mathematiques 

Le petit théorème de Fermat stipule que

- pour un nombre premier p
- pour un nombre entier a, tel que `a % p != 0`  (p n'est pas un multiple de a)

Alors `a^(p-1) % p = 1`

Ce théorème peut être utilisé pour créer un algorithme probabiliste de recherche de nombre premier.

Soit un nombre q, pour lequel on souhaite déterminer si il est premier ou non.

- Si on trouve un nombre a tel que `a^(q-1) % q != 1`, alors q n'est pas premier (Fermat Witness)
- Attention: La contraposée n'est pas vraie : il est possible que `a^(q-1) % q = 1`, sans que q soit premier (Fermat Liar). Pour un nombre a < p, cela n'arrive qu'avec un probabilité de 1/2. Donc il suffit d'essayer aléatoirement d'autres valeurs de a pour augmenter les chances que q soit effectives premier.

# 2023-07-19 Le pape François 1er

#religion 

Le pape François I, ancien archevêque de Buenos Aires est devenu pape en 2013, après la renonciation de Benoît XVI à 85 ans (il est décède 10 ans plus tard, en 2022).

Note: Le pape n'est reconnu que par le courant catholique du [[2023-06-24 Le christianisme|christianisme]]

François I est argentin et issue de la Compagnie de Jésus (les jésuites) et est le premier pape issu du continent américain.

Ses opinions et idées sont :
- Fort engagement pour la justice sociale, notamment contre la pauvreté 
- Décrie l'euthanasie et considère l'avortement est un problème éthique: l'humain existe dès la formation de son code génétique
- S'oppose au relations homosexuelles et à la transsexualité, mais appelle également à la tolérance et au dialogue 
- Agit pour le dialogue inter-religieux, notamment avec l'islam et le judaïsme
- Adhère à la théorie du Big Bang et à l'évolution, mais considère que ce n'est pas contradictoire avec l'existence d'un Dieu
- Il condamne l'individualisme et le "progrès matériel sans limite"
- Il soutien une meilleure protection de l'environnement, au travers ce qu'il appelle la Sauvegarde de la Création

# 2023-07-20 Haut Conseil pour le Climat

#ecologie #france #politique

Le Haut Conseil pour le Climat est un organisme neutre et indépendant, créé en 2019 (inscrit dans la loi) qui évalue l'action publique en terme de :
- la trajectoire de réduction des émissions
- les puits de carbone
- les émissions de carbone
- l'adaptation
- l'impact socio-économique 

Le Haut Conseil pour le climat compte 14 membres : climatologues, chercheurs, experts et produit un rapport annuel concernant l'action du gouvernement.

# 2023-07-21 L'histoire des peuples en France

#histoire #france 

L'histoire des peuples en France peut être résumé en quelques grandes étapes.

- D'abord de nombreux peuples celtes et gaulois se sont installés en France. Il utilisait l'alphabet grec
- Les romains ont ensuite conquis le territoire et l'ont réorganisé, en désignant Lugdunum (Lyon) comme capitale
- A la chute de [[2023-08-27 L'empire Romain]], diverses invasions barbares ont divisé le territoire, dont les francs (venu du rhin)
- Clovis, grâce à ses victoires militaires et sa conversion au christianisme, parvint a créer une certaine unité, sous le règne des mérovingiens
- Ce son cependant les Carolingiens, tributaire d'un fief, qui reprirent le pouvoir, notamment en repoussant les sarrasins à Poitiers

# 2023-07-22 Oppenheimer

#sciences #histoire 

Julius Robert Oppenheimer un physicien américain né en 1904. Il a dirigé la partie scientifique du Manhattan Project, visant à créer une bombe atomique.

Après l'explosion des deux bombes atomiques Little Boy et Fat Man à Hiroshima et Nagasaki au Japon, il s'est opposé à la prolifération des armes nucléaires et notamment aux recherches sur la bombe à hydrogène. Il fut banni des recherches pour cette raison et pour son penchant pour l'idéologie communiste.

Après les explosions, Oppenheimer a cité un écrit en Sanscrit "I am become death, the destroyer of worlds".

# 2023-07-23 Donner son sang

#france #biologie

En France il est possible de donner son sang toutes les 8 semaines. La collecte est volontaire et est organisée par l'EFS (Établissement Français du sang). Cela permet de soigner 1 million de malades, notamment pour de hémorragies, des maladies du sang, ces cas de cancer, de Drepanocytose, accouchements etc.

Entre 420 et 480 ml sont prélevés selon le poids et l'âge du donneur.

# 2023-07-24 Rendement moteur thermique et électrique

#energie #technologie 

L'efficacité d'un moteur thermique d'une voiture est autour de 36% pour une voiture a essence.

L'efficacité d'un moteur électrique est d'environ 90% dans une voiture électrique, donc une perte d'environ 10%. Cependant, à ces pertes s'ajoutent 
- les pertes de la centrale à la station de charge: 10%
- les pertes lors de la charge: 12%
- les pertes supplémentaires de rendement, lorsque le moteur est froid

Le rendement depuis la centrale est d'environ 60% à 70%

# 2023-07-25 Un kilowattheure

Un kilowattheure correspond à 1000 wattheure, ce qui est une quantité d'énergie.

Un Joule est l'énergie dépensée par une puissance de 1 watt pendant 1 seconde. Autrement dit : 1 watt-seconde = 1 joule ou encore 1 kWh = 3 600 kJ

Avec 1 kWh, il est possible de faire 
- 3 jour d'utilisation d'un téléphone 
- 1.5 jour d'utilisation d'un ordinateur portable
- 1 jour de fonctionnement d'un réfrigérateur
- 5h de télévision
- 6h de climatisation
- 1 heure de chauffage
- 2km avec une smart électrique

# 2023-07-26 Le funnel AAARRR

#product 

Le funnel AAARRR, aussi appelé funnel pirate, décrit les principales étapes de l'utilisation d'un produit par un utilisateur :

- Awareness - combien de personnes connaissent le produit ?
- Acquisition - combien de personnes se rendent sur le produit (site web, application)
- Activation - combien de personnes font une étape importante dans le produit ?
- Retention - combien de personnes retournent sur le produit ?
- Revenue - combien de personnes dépensent de l'argent ?
- Referral - combien de personnes recommandent le produit à d'autres personnes ?

# 2023-07-27 Le Fichier National Unique des Cycles Identifiés

#france #legal 

Le marquage de tous les vélos neufs est obligatoire depuis le 1er janvier 2021, ainsi que les vélos d'occasion vendu par des commerçants. Cette identification a pour objectif de lutter contre les vols de vélo (300000 par an en France).

Les données sont centralisés dans le Fichier National Unique des Cycles Identifiés (FNUCI), géré par l'APIC (Association de Promotion et d'Identification des Cycles). Ce fichier permet de générer des identifiants uniques ainsi que de traces les évènements du cycle de vie d'un vélo (vente, revente, vol, restitution, destruction)

L'Identifiant est un code de 10 caractères, délivré par un opérateur agréé par l'État, qui enregistre et inscrit cet identifiant sur le vélo (gravure, autocollant, capsule)

# 2023-07-29 Les tissus textiles

#sciences 

Il existe 4 grandes familles de textiles :
- les tissus textiles d'origine végétale (coton, lin, chanvre)
- les tissus textiles d'origine animale (laine, soie, cuir)
- les tissus textiles artificiels (viscose, lyocell, acétate)
- les tissus textiles synthétiques (polyester, nylon, acrylique, elasthanne)

# 2023-07-30 Le Hardergrater Trail

#europe #nature #sport 

A proximité de Berne, en [[2023-08-12 La Suisse]], le Hardergrater Trail est une randonnée sur une crête de 24km, longeant le lac de Brienz. Avec ses 3000m de dénivelés positifs, aucune ombre et aucune source d'eau, il fait partie des parcours les plus difficiles de Suisse.

Au départ, un funiculaire permet de monter directement jusqu'à Harder Kulm et d'éviter une partie de la montée. Sur le retour, une gare à Brienz permet de rentrer facilement jusqu'à Interlaken.

# 2023-07-31 Prolifération Nucléaire

#energie #politique #geographie 

En 2019, huit États souverains ont testé avec succès la détonation d'armes nucléaires:
- La France
- Les Royaumes Unis
- Les États-Unis
- La Chine
- La Russie
- L'Inde
- Le Pakistan
- La Corée du Nord

L'Israel est soupçonné de posséder des armes nucléaires, sans le reconnaître publiquement.

Le Traité de Non-Prolifération (TNP), entré en vigueur en 1970, est le principal outil permettant de limiter la prolifération de l'arme nucléaire. L'Inde, le Pakistan et la Corée du Nord ont développé l'arme nucléaire malgré ce traité. L'Iran est soupçonné d'avoir un programme nucléaire.

# 2023-08-01 DPE de l'appartement

*File not found.*

# 2023-08-03 Rotation d'un array

#algorithme #code 

Il est possible de faire une rotation d'array de k places en O(n), sans utiliser de mémoire supplémentaire.

Exemple d'une rotation de 2 vers la droite donne:
- avant : 1 2 3 4 5 6
- après : 5 6 1 2 3 4

Pour cela: 
1. inverser les `n-k` premiers éléments
2. inverser le reste
3. inverser le tout 

Dans l'exemple (n=6 et k=2)
- après 1) 4 3 2 1 5 6
- après 2) 4 3 2 1 6 5
- après 3) 5 6 1 2 3 4

# 2023-08-04 Ailefroide

#escalade #escalade 

Ailefroide est un hameau dans la vallée de Vallouise, dans le massif des écrins, en Hautes Alpes.

Le hameau est situé à 1500m d'altitude et est entouré de falaise avec des grandes voies, des couennes et des blocs de tout niveau. Frais en été, le camping d'Ailefroide peut accueillir jusqu'à 1000 grimpeurs.

Note: la connexion 4G est presque inexistante dans la vallée, notamment chez SFR et Free

# 2023-08-05 Oeuf frais et extra frais

Un oeuf est extra frais les 9 jours après la ponte
Un oeuf est frais 28 jours après la ponte

# 2023-08-06 Scanner tous les ports d'internet

#technologie 

Tenter une connexion (envoyer un packet SYN sur le port 80) à toutes les machines d'internet (4 milliards d'adresse IPv4) ne prends que quelques avec ~100 machines Google Cloud.

De David

# 2023-08-07 Composition de la bière

#nutrition 

La bière est principalement composée d'eau, de malt d'orge (=orge fermenté), de levure et de houblon. Le houblon donne son goût et son amertume à la bière, et agit comme conservateur. Certaines bières ont aussi du malt de blé (bières blanches) ou du malt de riz.

Il existe deux types de fermentation : haute (pour les bières de type "Ale") et basse (pour les bières de type "Lager").

Les différentes bières dépendent principalement de 3 facteurs:
- le type de levure utilisée (ainsi que le type de fermentation)
- le type et dosage de malt d'orge
- le type et dosage de houblon

De Thomas

# 2023-08-08 Le Zeugma

#littérature 

Le zeugma ou zeugme est une figure de style qui consiste à relier au même mot deux termes ou expressions incompatibles.

Par exemple :
Il prend son café et son courage à deux mains 

De Pauline

# 2023-08-09 Comètes, astéroïdes, météores et météorites

#espace #sciences 

Une comète est une masse de glace, de roche et de poussière qui traverse l'espace.

Un astéroïde est un corps de roche métallique ou non, de quelques centimètres à plusieurs km de diamètre, en orbite autour du Soleil, généralement dans la ceinture d'astéroïdes entre Mars et Jupiter.

Un météor est un corps qui se consume en traversant l'atmosphère de la Terre, ce qui produit une traînée lumineuse (étoile filante)

Une météorite est un météor qui ne se consomme pas complètement en traversant l'atmosphère, et un fragment arrive au sol.

# 2023-08-10 La reine des Alpes

#nature #france 

La Reine des Alpes, aussi appelé Chardon Bleu, est une espèce protegée de plante qui pousse dans les hauteurs des Alpes.

![[Pasted image 20230815105055.png]]

# 2023-08-11 Boulder et Lead aux JO

#escalade #competition 

Pour les Jeux Olympique de Paris 2024, deux médailles seront décernées pour l'escalade : une pour l'épreuve Speed Climbing et une pour l'épreuve combinée Boulder and Lead.

Pour l'épreuve B&L, 20 athlètes de chaque sexe vont être sélectionnés entre 2023 et 2024. 
- 3 places pour le top 3 de la finale des World Championship B&L à Bern
- 5 places pour les qualifications de chaque région (Afrique, Asie, Europe, Océanie et Amerique)
- 10 places sur la Olympic Qualifier Séries (2024)
- 1 place pour le pays hôte 
- 1 place pour les universités

# 2023-08-12 La Suisse

#europe #geographie #politique 

La Suisse est une confédération de 26 cantons, dont Bern est la capitale de facto. La Suisse est donc un État fédéral, dirigé par un conseil fédéral composé de X membres élus (et, la plupart du temps, réélus).

Presque 60% du pays est couvert par les Alpes, le reste était appelé "le plateau Suisse". Le pays compte près de 9 millions d'habitants.

Elle est fortement liée aux organisations internationales, comme la Croix-Rouge et héberge plusieurs sièges 
- Office des Nations Unies (ONU) à Genève 
- La banque des règlements internationaux à Bâle 
- Le siège de l'Organisation Mondiale de la Santé à Pregny-Chambésy
- Le comité international olympique à Lausanne 

La Suisse est membre fondateur de l'Association Européenne de Libre échange et membre de l'espace Schengen. Cependant, elle est ni membre de l'Union Européenne, ni de l'Espace Économique Européen.

Le pays possède l'Indice de Développement Humain le plus élevé du monde, et le 2eme PIB nominal le plus élevé.

# 2023-08-13 Les 4000 des Alpes

#geographie #europe 

L'Union Internationale des Associations d'Alpinistes (UIAA) a dressé en 1994 une liste "officielle" des sommets de plus de 4000m dans les Alpes, prenant en compte la topographie (prédominance, isolation), la morphologie (aspect) et l'intérêt alpinistique (qualité des voies).

Cette liste retient 82 sommets en France, Italie et Suisse, dont :
- 41 dans les Alpes Pennines (Italie)
- 28 dans le massif du Mont-Blanc
- 9 dans les Alpes Bernoises (Suisse)
- 2 dans le massif des Écrins 
- 1 dans la massif du Grand Paradis (Italie)
- 1 dans la chaîne de Bernina (Suisse/Italie)

D'autres listes ont étés constituées plus tôt, mais sans critères bien défini, notamment par l'autrichienne Karl Blodig, le premier alpiniste ayant gravi tous les 4000 (de sa liste).

# 2023-08-14 L'Office National des Forêts

#nature #france 

L'Office National des Forêts (ONF) est une Entreprise Public à caractère Industriel et Commercial (EPIC) chargé de la gestion des forêts publiques en France (dont DOM/TOM). Il y a prêt de 5 millions d'hectares en métropole ainsi que 6 millions en outre-mer.

Ses missions principales sont:
- La sylviculture (production et vente de bois)
- La surveillance et protection des forêts (prévention incendie, dunes du littoral, terrain érodables en montagne)
- La police de la nature : les forestiers sont habilités à verbaliser selon le code forestier, le code de l'environnement mais également selon le code pénal.
- L'aménagement des forêts (chemins, aires) et animations de visites

# 2023-08-16 Demission et chômage

#legal #finance 

Dans le cas général, il n'est pas possible de bénéficier du chômage suite à une démission. Il existe cependant quelques exceptions :
1) Cas de démission "légitime" (17 cas différents, dont rupture conventionnelle, changement de lieu de résidence pour suivre le conjoint, échec de création d'entreprise, etc)
2) Reconversion professionnelle
3) Abandon de poste pour raison légitime (médicale, droit de grève, modification du contrat de travail par l'employeur, etc.)
4) Réexamen de la situation après 4 mois de chômage par l'Instance Paritaire Régionale (IPR) (121 jours après la démission). Il faut montrer les efforts déployés pour retrouver un emploi durant ces 4 mois.

# 2023-08-17 Les Arrays et Slices en Go

#code #go

Les arrays en Go ont une taille fixe non modifiable. Un array de 8 ints est déclaré de cette façon:

```
var myArray = [8]int
```

Une slice en Go décrit une partie d'un array. C'est l'array sous-jacent qui possède réellement la donnée, le slice ne fait que pointer vers une partie de cet array.

La structure de donnée d'un slice est généralement appelée le "slice header" et comporte : 
- un pointeur vers un élément de l'array (l'index 0 du slice)
- la length du slice
- la capacity restante de l'array sous jacent (car un slice ne peux pas dépasser la taille de l'array sous-jacent)

Pour créer un slice à partir d'un array:

```
var mySlice = array[start:end]
```

A noter que `start` est inclusif et `end` est exclusif. Pour s'en souvenir, il suffit de retenir que :

```
mySlice == mySlice[0:len(mySlice)]
```

Les slices passées en paramètre de fonction sont passées *par valeur*, donc modifier la slice en elle même dans une fonction n'aura pas d'impact en dehors. Cependant, l'array sous-jacent est accédé via un pointeur, et modifier ses valeurs aura un impact en dehors de la fonction.

La fonction `make` permet de déclarer un slice et un array en une seule ligne:

```
var mySlice = make([]int, length, capacity)
```

La fonction `copy` permet de copier une slice dans une autre, et retourne le nombre d'éléments copiés :

```
var count = copy(newSlice, oldSlice)
```

La fonction `append` permet d'ajouter un ou plusieurs éléments à la fin d'une slice, et si la capacité de l'array sous-jacent est dépassée, de créer un nouvel array de plus grande taille

```
mySlice = append(mySlice, 5)
mySlice = append(mySlice, 5, 3, 2, 4)
mySlice = append(mySlice, otherSlice...)
```

En Go, une `string` est une slice de bytes qui est read-only.

# 2023-08-18 Les refuges FFCAM

#nature #france #escalade 

La Fédération Française des Clubs Alpins et Montagne (FFCAM) possède 120 refuges et chalets en France (principalement Alpes, Pyrenees, Jura, Vosges) et 5 au Maroc. Ces refuges ont des périodes de gardiennage (généralement l'été) mais sont normalement accessible en dehors de ces périodes.

Les tarifs de nuitée pour 1 adulte varient d'environ 16€ pour les refuges faciles d'accès à 50€ ou plus pour les plus isolés. Les adhérents FFCAM ont une réduction de 50%. Hors période de gardiennage, le tarif est généralement de 8€/nuit, avec paiement et réservation en ligne.

Par un accord de réciprocité, les adhérents des clubs montagnards d'autres pays d'Europe (Suisse, Italie, Autriche, Allemagne, etc) bénéficient aussi de cette réduction de 50%.

Dans la plupart des cas, il est possible de bivouaquer à proximité d'un refuge, mais c'est parfois payant.

# 2023-08-19 Les projets en bâtiment

#urbanisme

Un projet de construction ou modification d'un bâtiment comporte quelques rôles clés:
- La **foncière** est l'entreprise qui possède le terrain et fait construire un bâtiment sur ce dernier.
- Le **promoteur immobilier** fait la promotion du projet pour trouver un propriétaire (soit def l'espace entier, soit de certains lots), que le projet soit terminé ou non
- Le **maitre d'ouvrage** est le commanditaire du projet immobilier. Si c'est un acteur public, il doit faire un appel d'offre public avec un cahier des charges du projet. Si il est privé, il n'est pas obligé de faire un appel d'offre.
- Le **maitre d'oeuvre** conçoit un projet qui correspond au cahier des charges. Il comprend généralement un cabinet d'architects
- Les **artisans** ou l'**entreprise générale** sont les prestataires qui font les travaux planifiés par le maitre d'oeuvre
- Le **propriétaire** de l'espace est un bailleur si il propose un bail, c'est à dire qu'il loue son espace à des locataires. Le propriétaire peut également être occupant de l'espace.

# 2023-08-20 Les chargeurs compatibles

#hardware #energie 

Les chargeurs (téléphone, ordinateur, etc) délivrent un certain voltage et ampérage. Par exemple 20V/3A, ce qui fait une puissance délivrée de 60W. 

Charger un appareil avec son chargeur d'origine est idéal, car permet d'utiliser la technologie Quick Charge si disponible.

Charger un appareil avec un chargeur moins puissant fonctionne, mais la recharge est plus lente. 

Charger un appareil avec un chargeur plus puissant ne pose pas de problème car le device est muni d'une puce permettant de tirer uniquement la puissance nécessaire (à vérifier).

Dans tous les cas, il est préféré d'utiliser des chargeurs qualité, qui possèdent des sécurités supplémentaires.

# 2023-08-21 La culture du chanvre en France

#biologie #legal #france 

Le chanvre est un espèce de plante issue du cannabis sativa ayant une teneur en THC inférieure à 0.2%.

La culture du chanvre est autorisée en France, pour une 50aines de variétés sélectionnées. Le chanvre est cultivé 
- pour sa tige, transformée en paille de chanvre (papiers, isolation, plastiques biosourcés, usage en bâtiment)
- pour ses graines, le chènevis (huiles, cosmétiques)
- et pour sa fleur (usage pharmaceutique)

La culture de chanvre existe depuis l'histoire, mais sa culture a failli disparaître au XXe avec l'arrivée de la pétrochimie. Elle redevient intéressante ces dernières années notamment pour des questions écologiques (production locale, sans phytosanitaires et sans OGM).

# 2023-08-22 La taxonomie européenne

#europe #finance #ecologie 

La taxonomie européenne classifie les activités selon leur impact sur l'environment, dans l'objectif de neutralité carbone en 2050. Une première taxonomie a été adoptée par l'UE en 2020, puis en 2022 une modification ajoute le gaz et le nucléaire comme activités dites "vertes".

Pour qu'une activité soit classée verte, elle doit contribuer positivement à un ou plusieurs des 6 critères, sans être néfastes pour les autres :
- Atténuation du changement climatique (réduction des émission de CO2, etc.)
- Adaptation au changement climatique
- Utilisation durable et protection des milieux aquatiques et marins
- Transition vers [[2023-08-23 L'économie circulaire]]
- Protection de la biodiversité
- Réduction de la pollution

Cette taxonomie concerne :
- Les entreprises, qui doivent indiquer quelle part de leur chiffre d'affaire, de leurs investissements et dépenses correspond à des activités durables
- Les états membres qui mettent en place des mesures publiques, des normes, des labels ou des produits financiers ou obligations
- Les acteurs financiers (banques, banques centrales, companies d'assurance)

https://www.vie-publique.fr/questions-reponses/283166-neutralite-carbone-la-taxonomie-europeenne-en-six-questions

# 2023-08-23 L'économie circulaire

#economie #ecologie #legal 

L'économie circulaire s'oppose à l'économie dite linéaire, qui consiste à produire, utiliser et jeter. Dans ce nouveau modèle économique, l'objectif est de produire des biens et des services de manière durable en limitant le gaspillage des ressources et la production de déchets.

Pour cela, l'économie circulaire promeut le réemploi, la vente des produits d'occasion, le recyclage. Ce modèle s'articule sur 3 axes:
- Réduction de l'impact environmental de l'extraction et exploitation des ressources naturelles, éco-conception (fabrication durable), mutualisation des infrastructures, localité
- Demande de comportement des consommateurs, allongement des durées de vie (réparation), information sur l'impact environmental de la consommation du produit
- Gestion des déchets (recyclage)

Des lois en France en 2015, 2020 et 2021 pour favoriser la transition vers une économie circulaire sont en cours d'application. Un plan d'action proposé par la Commission Européenne promeut également cette transition au niveau européen.

# 2023-08-24 Équinoxe et solstice

#espace #sciences 

Le solstice (sol statum: Soleil immobile) est un évènement astronomique où le Soleil vu de la Terre atteint sont extrême méridional (position la plus au Sud) ou septentrional (position la plus au Nord). Le solstice d'été correspond au jour le plus long (dans l'hémisphère nord) et le solstice d'hiver est le jour le plus court.

Les deux équinoxes (equi nox: égal nuit) de l'année (égal nuit) se produisent entre les deux et solstices et correspondant au moment où le Soleil traverse le plan équatorial terrestre. La durée du jour est égal à la durée de la nuit.

Dates approximatives, correspondant au passage des [[2023-03-15 Les Saisons]].
- 20 mars : équinoxe (printemps)
- 20 juin : solstice (été)
- 20 septembre : équinoxe (automne)
- 20 décembre : solstice (hiver)

# 2023-08-25 Préparer un cheval

#equitation

Pour préparer un cheval avant de le monter, il faut :

Brosser le cheval
- décoller les poils avec l'étrille
- brosser pour retirer les poils avec le bouchon
- curer les pieds avec le cure-pied

Sur le dos du cheval
- mettre la peau de mouton sur le garrot (à l'envers)
- mettre le tapis
- mettre l'amortisseur
- mettre la selle
- dégarotter (relever la peau de mouton au max)
- attacher la sangle sur les contre-sanglons et la passer dans le tapis

Sur la tête du cheval
- mettre le filet avec les mors + attacher la sous-gorge + la muserolle
- éventuellement mettre l'élastique (entre la sangle et les anneaux des mords)
- passer les reines au dessus de la tête du cheval 

Éventuellement ajouter des guetres sur les pieds du cheval, pour protéger ses pieds (pour faire de saut)

Équipement de base du cavalier
- la bombe
- le pentalon d'équitation + chaussettes hautes
- les bottines 
- les schaps
- éventuellement des gants 

Avant de monter sur le cheval, penser à vérifier la selle et resserrer la sangle (le ventre du cheval dégonfle)

# 2023-08-26 Assainir la seine

#france #urbanisme 

Il est officiellement interdit de se baigner dans la seine depuis maintenant 100 ans. Outre les déchets jetés dans l'eau chaque année, ce sont les déchets fécales et bactéries associés qui rendent l'eau impropre.

En temps normal, l'eau qui sort du système d'égout est évidemment traitée avant d'être relâchée dans la seine en aval. Cependant les égouts à Paris ont un défaut: les eaux usées des maisons sont mélangées avec l'eau de pluie. 

Lorsqu'il pleut en trop grande quantité, les évacuations d'eau sont saturés et le surplus est directement relâché dans la Seine, sans aucun traitement. Cela arrive environ 12 fois par an.

Paris investi depuis des années pour :
1) Construire un énorme cuve qui pourrait stocker l'eau en cas de précipitation importante et éviter de relâcher dans la Seine
2) Forcer toutes les habitations et notamment les 100+ péniches à se raccorder aux égouts

# 2023-08-27 L'empire Romain

#histoire #europe 

La ville de Rome est fondée au VIII AD. La légende latine dit que Romulus pose la première pierre sur le mont Palatin. En -509, la ville, jusqu'à sous forme de Royaume, devient une République.

Rome s'étend dans la Méditerranée Occidentale, notamment en dominant Carthage, une cité nord-africaine en pleine expansion : ce sont les guerres les guerres puniques. En -200, Hannibal, fils d'un général Carthaginois, chercher à prendre sa revanche. Il mène une expedition depuis les colonies d'Espagne, traverse les Pyrénées, la Gaule, les Alpes et sème la terreur et la destruction à Rome. Mais Rome résiste et Carthage fini par capituler.

Dans les siècles suivants, la République Romaine envahit l'Espagne et le Portugal (-133), la Gaule Narbonnaise (e.g le sud de la France en -122), la Macédoine (e.g la Grèce en -148), la Gaule (en -51 par Jules César) puis une partie de l'Afrique du Nord en -46.

A partir de -27, Octave, le neveu de César, impose la Paix Romaine : c'est le début de la période impériale. L'empire continue sa politique expansionniste et atteint son expansion maximale en 117, puis passe à une politique de consolidation (Empereur Adrien).

Au IIIe siècle, des invasions barbares par l'est font entrave à la Paix Romaine.

En 380, l'Edit de Thessalonique, promulguée par Theodose Ier, fait passer le [[2023-06-24 Le christianisme]] au rang de "catholique", donc religion officielle de l'Empire.

En 395, l'Empire est partagé entre les deux fils de Theodose et devient l'Empire Romain d'Occident et d'Orient.

L'Empire d'Occident ne parvient pas à repousser les invasions barbares et, en 476, le dernier Empereur d'Occident déposé.

L'Empire d'Orient, nommé Empire Byzantin par les historiens, existe jusqu'en 1453, lors de la prise de Constantinople par les Turcs.

# 2023-08-28 Les languages Lisp

#code

LISP (List Processing) est un langage de programmation inventé en 1958 par John McCarthy (un an après Fortran). Le langage a été standardisé par l'American National Standards Institute (ANSI) en 1994 pour devenir le Common Lisp. C'est aujourd'hui une famille de langages, dont Scheme et Common Lisp et plus récemment Clojure sont les représentants les plus importants.

Lisp est un langage de programmation fonctionnel, où le code source est écrit sous forme de liste qui sont interprétées (ou compilées) par le Lisp interpreter. La particularité de ce langage est que le code source en lui même est une structure de donnée qui peut être manipulée par le programme, donnant une grande liberté d'automatisation et de meta-programming.

# 2023-08-29 Les streams de Casino

#web #legal 

Les streams Twitch de casino ont explosés ces dernières années, notamment le streaming de paris en ligne sur stake.com et autre.

Certains streams ont finalement été interdit sur Twitch, ce qui a poussé les fondateurs de stake.com de créer une plateforme de streaming concurrente : kick.com

Pour le casino, le modèle économique est que:
1) les streameurs avec une forte audience/influence reçoivent des credits gratuits sur les sites de paris
2) ils streament leurs jeux et empochent les gains
3) L'audience joue sur le site de paris et perdent des grandes quantités d'argent, ce qui rembourse l'argent initialement "offert" au streameur

# 2023-08-30 Résiliation en ligne

#legal #france 

A partir de 1er septembre 2023, tout abonnement souscrit en ligne doit pouvoir être résilié en ligne. (Loi de protection du pouvoir d'achat de 2022)

https://www.francenum.gouv.fr/guides-et-conseils/developpement-commercial/gestion-de-la-relation-client/resiliation-en-3-clics

# 2023-09-01 Les randonnées Norvégiennes

#voyage #nature 

Les montagnes les plus connues en Norvège sont :

Le Preikestolen (Le rocher de la Chaire à Songesand) est un promontoire de pierre qui culmine au dessus du Lysefjord. La randonnée dure environ 4h pour l'aller-retour depuis le refuge. https://www.visitnorway.fr/destinations-norvege/region-fjords/ryfylke/lysefjord/randonnee-preikestolen/

Trolltunga (la langue du troll) est un rocher surplombant de 700m le lac Ringedalsvatnet, situé sur le plateau montagneux de Hardangervidda. La randonnée est difficile et dure 10h-12h (20km avec 800m à 2000m de D+, selon le point de départ) https://www.visitnorway.fr/destinations-norvege/region-fjords/hardangerfjord/la-randonnee-de-trolltunga/

La crête de Romsdalseggen est une randonnée vertigineuse et difficile de 7-8 heures. La crête se trouve dans le Andalsnes. https://www.visitnorway.fr/destinations-norvege/region-fjords/nord-ouest/randonnee-romsdalseggen/

La randonnée du Kjerag est l'ascension du sommet Kjerag, plus haut sommet du Lysefjord (à 1084m), qui est aussi un spot de base-jump et d'escalade. La randonnée du 6-10h  (11km avec 800m D+) https://www.visitnorway.fr/destinations-norvege/region-fjords/ryfylke/lysefjord/randonnee-du-kjerag/

# 2023-09-02 Le stress

#biologie #psychologie 

Le stress est une réponse physiologique, biologique et psychologique face à des pressions et contraintes (stresseurs). En réaction, le corps humain produit deux hormones : l'adrénaline et le cortisol, ce qui augmente temporairement la concentration et les capacité physiques.

Le cortisol a plusieurs actions :
- Il augmente le taux de glycémie (concentration de glucose dans le sang)
- inhibe certaines réponse du système immunitaire
- régule les lipides, protéines et glucides

L'action du stress peut être positif et nous pousser à réussir: on parle de stress positif (eustress), mais au contraire, elle peut aussi être néfaste et nous déstabiliser complètement: le stress négatif (distress). Le stress est différent de l'anxiété, qui est une émotion, tandis que le stress est un mécanisme qui peut amener à différentes émotions, dont l'anxiété.

# 2023-09-03 Le CAC 40

#finance #france 

Le CAC 40 (Cotation Assistée en Continu) est un indice boursier actualisé toutes les 15 secondes entre 9h et 17h30. Il est composé de 40 des 100 plus grosses entreprises françaises par capitalisation. Les 40 sont choisies par un comité scientifique se réunissant tous les 3 mois, avec deux critères principaux:
- le volume de transaction sur le titre de l'entreprise
- la représentativité des secteurs d'activité (agro-alimentaire, énergie, automobile, industrie, luxe, pharmacie, services)

Le CAC 40 est mesuré en points, avec un point de départ de 1000 points pour 1000€ investis en 1987, date de la création de l'index. Les entreprises avec plus de capitalisation et plus de flottant (titres disponibles à la vente) ont un plus gros poids dans le CAC 40, avec un max à 15%.

Quelques dates:
- Avant la crise de la bulle internet en 2000, le CAC 40 avait près de 7000 points (puis descente à 2500)
- Avant la crise des subprimes en 2008, le CAC 40 avait près de 6000 points (puis descente à 3000)
- Avant COVID en 2020: 6000 points
- Aujoutd'hui 7300 points

# 2023-09-04 La détection de cycles

#code #algorithme 

Il y a deux principaux algorithmes pour détecter un cycle (dans une linked list, par exemple) :
- Avec un hashset
- Avec l'algorithme de Floyd (Tortoise and Hare Alorithm)

Avec un hashset, il suffit de noter les nodes visitées dans le hashset. Si une nouvelle node est déjà présente dans le hashset, alors il y a un cycle. Cet algorithme est 0(n) en temps et en espace.

Avec l'algorithme de Floyd, deux "pointeurs" sont utilisés et parcourent la liste. Un pointeur (le lapin) avance deux fois plus vite que le deuxième (la tortue). Si il y a un cycle, alors la tortue et le lapin finiront par se retrouver au même point. Il est ensuite possible de calculer le point de départ du cycle, si besoin.

# 2023-09-05 L'anarchie

#philosophie #politique 

Du point de vue étymologique, "anarchie" signifie "sans commandement". Souvent utilisé comme terme péjoratif pour désigner le chaos et un manque d'organisation, une société anarchique est en réalité une société auto-gérée. Autrement dit, c'est un style de vie et un modèle de société rejetant toute autorité ou hiérarchie, et cherche à promouvoir la liberté et l'égalité. 

D'un point de vue politique, le mouvement anarchique est un mouvement contestataire, rejetant l'Etat et le gouvernement, notamment via du militantisme et un rejet des élections legislatives (ex: vote blanc).

Certains éco-villages (ex: [[2023-03-31 Eco-village Emmaüs Lescar Pau]]) peuvent être considérés comme anarchiques.

# 2023-09-06 Dart 3.0

#code #dart

La mise a jour de Dart en version 3.0 (mai 2023) ajoute des fonctionnalités de programmation fonctionnelle au langage, avec entre autre :
- les records, qui permettent aux fonctions de renvoyer plusieurs variables
```
var hello = (42, "ok", name: "Hello")
```

- les patterns, qui permettent de "matcher" et de "déstructurer" une variable, notamment dans un `switch`
- les class modifiers (interface, base et final, puis sealed sur Dart 3.1)

# 2023-09-07 Principes de la thermodynamique

Le principe zéro (principe d'équilibre thermique, 1931)
> Si deux systèmes sont en équilibre thermique stable avec un troisième, ils sont eux-même en équilibre thermique stable

Le 1er principe (principe d'équivalence)
> Au cours d'une transformation quelconque d'un système fermé, la variation de son énergie est égale à la quantité d'énergie échangée avec le milieu extérieur, par transfert thermique et transfert mécanique

Le 2eme principe (principe de Carnot, 1824)
> Toute transformation d'un système thermodynamique s'effectue avec une augmentation de l'entropie globale incluant l'entropie du système et du milieu extérieur.

Le 3eme principe (principe de Nernst, 1906) :
> La valeur de l'entropie de tout corps pur dans l'état de cristal parfait est nulle à la température de 0 kelvin.

# 2023-09-08 Le système de retraite

#france #legal 

En France, le système de retraite est un système par répartition: les cotisations de la population active est directement utilisé pour payer les retraites la même année. À inverse, dans un système de retraite par capitalisation, chaque personne cotise pour sa propre retraite.

Il existe 3 régimes différents de retraite :
- la retraite de base, obligatoire (qui utilise un calcul par annuités)
- la retraite complémentaire, obligatoire également (qui utilise le calcul par points)
- la retraite supplémentaire, facultative (par capitalisation)

Il existe plusieurs caisses de retraite, en fonction des catégories socio-professionnelle (ex: caisse pour les salariés du privé, pour les fonctionnaires, pour les commerçants et artisans, etc). Si il y avait prêt de 100 caisses différentes dans les années 70, il n'en existe déjà plus que 20, suite à différentes fusions pour simplifier le système.

Les deux caisses à connaître pour les salariés du public ET pour les indépendants (depuis 2021) :
- L'Assurance Retraite (= régime général) pour la retraite de base, qui utilise donc un calcul par annuités (=par trimestres)
- L'Agirc-Accro pour la retraite complémentaire, qui utilise donc un calcul par points 

Pour la retraite de base le calcul de la pension utilise :
- le salaire moyen sur les 25 meilleures années de salaire 
- un taux de max 50% qui diminue en fonction du nombre de trimestres manquants
- le ration trimestres validés / trimestres requis

Dans tous les cas, il n'est possible de partir à la retraite avant 64 ans. Et à l'inverse, dans tous les cas, pour un départ à la retraite à partir de 67 ans, le taux est plein (=50%)

Pour la retraite de base, il est possible de racheter certains trimestres non-cotisés (le prix varie en fonction de l'âge et du salaire). Racheter ces trimestres tôt peut être intéressant financièrement pour partir à "taux plein".

# 2023-09-09 Le Danemark

#europe #geographie #politique 

Le Danemark est un pays Scandinave de 6 millions d'habitants. Le régime politique est une monarchie constitutionnelle (avec un parlement), dont la Reine Margarethe II depuis 1972.

Note: la Scandinavie est constituée uniquement du Danemark, de la Suède et de la Norvège.

Le territoire du Danemark est le Danemark métropolitain (péninsule + 443 îles), les îles Féroé (atlantique) et le Groenland (qui est politiquement indépendant).

Comme la [[2023-08-12 La Suisse]], le Danemark a un Produit Intérieur Brut (PIB) très élevé, tout comme son Indice de Développement Humain (IDH): la population danoise est souvent cité comme la plus heureuse au monde.

Le Danemark est un membre fondateur de [[2023-04-10 L'OTAN]], des Nations Unies, de l'Union Européenne et se l'espace Schengen.

Malgré ses faibles ressources naturelles, le Danemark a une économie forte, notamment via l'innovation et des entreprises internationales tel que The Lego Group, Carlsberg et Velux. Son activité est principalement dans le secteur tertiaire.

# 2023-09-10 Multi-Audio Output sur macOS

#tools #musique 

Sur macOS Ventura, il est possible d'utiliser plusieurs sorties audio. Par exemple, deux personnes peuvent regarder un film sans utiliser les hauts parleurs: une sur des écouteurs filaires et l'autre des écouteurs bluetooth.

Pour configurer la sortie audio:
1. Sur l'application "Audio MIDI Setup", créer un Aggregated Audio Device (cliquer sur le +) et sélectionner les sorties voulues
2. Dans les settings audio du mac, sélectionner "Aggregated Audio Device" qui a été créé à l'étape précédente

# 2023-09-11 Christiana à Copenhague

#voyage #sociologie 

Christiania est un quartier de Copenhague au [[2023-09-09 Le Danemark|Danemark]], construit sur une ancienne caserne. Entre 1971 et 2013, le quartier a fonctionné comme une commune auto-gérée sous une certaine forme [[2023-09-05 L'anarchie|d'anarchie]]. 

La communauté a été fondée par des squatteurs, chômeurs et hippies. La communauté possédait sa propre monnaie, ses activités sportives et culturelles, ainsi qu'un terrain agricole.

Aujourd'hui, la vente de canabis est tolérée, mais les drogues dures sont interdites. C'est également une attraction touristique qui attire jusqu'à 1 million de visiteurs par an.

# 2023-09-12 Température d'un frigo

#appartement 

La température d'un frigo est idéalement comprise entre 2°C et 8°C. Pour un congélateur, c'est entre -6°C et -18°C.

A noter que la température n'est généralement pas homogène dans un frigo: la porte est souvent plus chaude et selon les frigos, certaines parties se refroidissement mieux que d'autres.

# 2023-09-13 Voiture automatique

#voyage 

Les boîtes de vitesses de voiture automatique ont généralement 4 modes :
- P pour Parking: utilisé à l'arrêt, pour éviter que la voiture ne bouge. C'est la position de départ pour démarrer
- R pour Reverse
- N pour Neutral, a utiliser lors d'un feu rouge par exemple. En pratique, ce n'est jamais utilisé
- D pour Drive, c'est la position pour avancer

# 2023-09-14 Les aurores boréales

#sciences #espace 

Les aurores polaires (boréales dans le nord et australes dans le sud) sont provoquées par des éruptions solaires. C'est éruptions projettent des particules chargées (électrons, protons, ions) vers la Terre.

La Terre est protégée de ces particules grâce au bouclier magnétique de la Terre : la magnétosphère. Cependant, lors d'une éruption solaire, une partie de ces particules est "guidé" par le champ magnétique vers le pôle côté nuit.

Les particules chargées interagissent avec la ionosphère (entre 80 en 1000km d'altitude) et excitent (donc ionisent) des atomes d'oxygène, d'azote et d'hydrogène. En se déchargeant, les atomes émettent des photons dans le vert, le rouge, le bleu.

# 2023-09-15 La Hanshelleren cave

#escalade #voyage 

La Hanshelleren cave est une cave de [[2023-06-10 Les types de roche|granite]] à proximité de Flatanger, en Norvège. Cette cave est devenue célèbre à la suite de l'ascension d'Adam Ondra (grimpeur Tchèque) de la voie Silence en septembre 2017. 

Silence est la 1ere voie cotée [[2023-03-12 Les notations en Escalade|9c]]. Elle mesure 45m, avec 3 problèmes de bloc après 20m d'ascension et de multiples knee-bars. Pour l'instant, elle n'a jamais été répétée, mais le grimpeur italien Stefano Ghisolfi l'a travaillé en 2022 et souhaite y retourner.

La seule autre voie 9c à ce jour a été grimpée par Seb Bouin : DNA, dans la cave "La Ramicole", dans les gorges du Verdon, en France.

En 2013, dans la Hanshelleren cave, Adam Ondra a réalisé l'ascension de Change (9b+) et a équipé Silence.

Aujourd'hui, Adam Ondra travaillerait sur un autre projet dans la cave, nommé Projet Big.

# 2023-09-16 Composition de l'atmosphère

#sciences #nature 

L'atmosphère est principalement composée de :
- Diazote (N2) à 78%
- Dioxygène (O2) à 21%
- Des gars rares, comme l'Aragon (<1%)
- une très faible quantité de Dioxyde de carbone (CO2) à 0.04%
- Des traces d'helium, d'hydrogène, de krypton, de méthane, de xénon, d'ozone
- De la vapeur d'eau dans les couches basses (nuages) entre 0% et 4%

# 2023-09-17 Next40 et FrenchTech 120

#entreprenariat #france 

Depuis 2019, la FrenchTech accompagne les startups françaises avec le plus fort potentiel, via un programme d'État d'1 an. Chaque année une nouvelle promotion est choisie, sur des critères de performance économique (levée de fonds, hypercroissance des revenus).

Le Next40 est le top des 120 startups sélectionnées dans le programme. On y retrouve entre autre : Alan, BackMarket, BlaBlaCar, Lydia, Payfit, Doctolib, Voodoo, sorare.

https://lafrenchtech.com/fr/la-france-aide-les-startup/french-tech-120-2/

# 2023-09-18 L'Union Européenne

#europe #legal #voyage #politique 

L'Union Européenne (UE) est une union politique et économique entre pays de l'Europe. L'UE est régie par un traité fondateur: le traité de Maastricht (Pays-Bas) en 1992.  En 2023, l'UE compte 27 État-membres. 
- Les Etat-membres payent une contribution vers l'UE (ex: la France paye environ 27 Mrd/an)
- Les Etat-membres votent des traités européens qui s'appliquent à tous les membres
- Les citoyens des Etats-membres sont citoyens européens, et peuvent vivre et travailler dans n'importe quel pays de l'UE

L'Islande, il la Norvège et le Lichtenstein ne font pas partie de l'UE, mais de l'European Economic Area (EEA), qui englobe également l'UE. Ils payent une contribution à l'UE et doivent suivre certaines lois européennes (sauf sur la pêche, l'agriculture, la justice, les taxes) et, en échange, les citoyens de ces pays peuvent vivre n'importe où en UE, et inversement.

Une majorité de l'EEA et la Suisse font partie de l'Espace Schengen, qui est un espace de libre circulation des personnes: il n'y a pas de contrôle aux frontière intérieures. Les Royaumes-Unis et l'Irlande n'en font pas partie.

L'Euro est utilisé par une majorité (mais pas tous) les membres de l'UE: l'est Euro Zone.

# 2023-09-19 La pluviométrie

#sciences #nature 

La pluviométrie est la mesure quantitative des précipitations en un lieu donné. Le pluviomètre permet de donner une mesure en mm : la hauteur d'eau récupérée dans un tube.

Il faut noter que cette hauteur ne dépend pas de du diamètre du tube, car un tube plus grand récupère plus d'eau, mais il faut également plus d'eau pour le remplir. C'est pourquoi la mesure est indépendante de la surface au sol utilisée.

Cependant, on peut utiliser cette mesure pour calculer le volume d'eau tombé. Par exemple si il est tombé 10mm un jour donné, alors, sur 1 m2, il est tombé `10 mm * 1 m2 = 0.01 m3 = 10 Litres

En Europe, les précipitations sont autour de 400 mm/an dans certaines parties de la région méditerranéenne et certaines plaines centrales européennes. Certaines parties du sud de l'Europe ne dépassent pas 50 mm/an de précipitations effectives, tandis que les côtes atlantiques de l'Espagne et de la Norvège atteignent 1000 mm/an.

# 2023-09-20 La couleur des glaciers

#sciences #nature 

Les glaciers jeunes sont généralement pleins de bulles d'air, et apparaissent en couleur blanche (couleur de la neige). Cependant, des glaciers plus anciens ont été compactés, parfois sous une forte pression. Dans ce cas, la glace forme des cristaux qui n'absorbent pas les longueurs d'ondes bleues de la lumière du soleil. Ces ondes bleues sont donc réfléchies, ce qui donne sa couleur bleue.

# 2023-09-21 Ouverture a la concurrence des mobilités

#france #europe #legal

La loi européenne "Organisation et Régulation du Transport Ferroviaire" (ORTF) de 2009 fixe des dates limites d'ouverture à la concurrence des services de transport en commun.

Les Autorités Organisatrice (AO) des territoires (Île de France Mobilité, par ex) doivent faire un appel d'offre public pour choisir un opérateur, ce qui est fin au monopole de l'exploitation de certaines lignes (par la [[2023-07-08 RATP et SNCF|RATP]], entre autre).

Les dates limites sont :
- 2025 pour les bus 
- 2030 pour les tramways 
- 2040 pour les métros et RER

Cette ouverture à la concurrence est déjà effective pour les nouvelles lignes et certains bus (grande couronne)

https://askip.ratpgroup.com/question/pourquoi-ouvrir-a-la-concurrence/

# 2023-09-22 Organiser un module Go

#code #go 

Un projet Go peut être une librairie (à importer dans du code), un cli (à exécuter directement) ou une combinaison des deux.

1) Pour une librairie simple
```
folder/
  go.mod
  modname.go
  modname_test.go
```

Note: Il n'y a pas de fonction `main` dans `modname.go`. Dans `go.mod`, il y a la ligne `github.com/username/modname` et le package dans `modname.go` est `package modname`.

2) Pour un CLI simple 
```
folder/
  go.mod
  main.go
  main_test.go
```

Il y a une fonction `main` dans `main.go`
Ce CLI peut être installé avec `go install github.com/username/modname@latest`

3) Une librairie contenant des packages internes
Pour que des sous-packages ne puissent pas être importés directement par un client, il faut les placer dans un dossier `internal`

```
folder/
  internal/
    auth/
      auth.go
    hash/
      hash.go
  modname.go
  go.mod
```

4) Un CLI avec plusieurs commandes
Dans ce cas, les commandes sont placées dans des dossiers séparés, dans dossier `cmd`.

```
folder/
  go.mod
  modname.go
  cmd/
    prog1/
      main.go
    prog2/
      main.go
```

Les commandes sont installées avec `go install gitHub.com/username/modname/cmd/prog1@latest`

# 2023-09-23 La fin de l'ère des startups

#idee #startup #user-experience #ecologie 

Les 10 dernières années ont vu une explosion du nombre de de startups et plus largements d'entrepreneurs, notamment en France autour de The Family, Station F, etc.

Avec cela, de nombreux secteurs ont été révolutionnés, la principale valeur ajoutée étant de placer l'Expérience Utilisateur au coeur du produit (ex: Uber vs les taxis, Doctolib vs le téléphone, Alan, Lydia, etc). De plus, certains usages de location/partage n' étaient pas vraiment démocratisés, et qui sont devenus accessibles en retirant une grosse partie de la friction (covoiturage avec BlaBlaCar, location courte durée avec Airbnb, etc).

Dans la tech aussi, la principale valeur ajoutée est la simplicité d'utilisation et l'Experience Utilisateur (Figma vs Illustrator, Heroku vs les servers on-premise, Linear vs Jira, etc).

Suite à des succès fulgurants de certaines startups dans les 10 dernières années(notamment américaines), les Venture Capitals (VC) se sont intéressés de très près à ce type de placement financier: le capital risque, très intéressant car très intéressant sur le Return On Investment, même sur quelques années. Cependant, les VCs en série A n'investissent que s'ils s'attendent à une série B. Les VCs en série B attendent une série C, etc. Et finalement, les derniers s'attendent à un Initial Public Offering (IPO).

Mon avis est que les plus gros marchés à révolutionner (en terme d'UX, donc) ont déjà vu un bon nombre de startups s'y essayer, et parfois réussir. Il reste encore beaucoup de place pour de nouvelles startups, certes, mais sur des marchés plus petits, plus spécialisés et donc globalement à plus faible retour sur investissement. Si une startup n'a pas vocation finale à faire une IPO, cela peut décourager les VCs.

Ma prédiction est que si sur les 10 dernières années, l'entreprenariat était concentré sur les sur la création de startup pour faire de l'hypercroissance, les 10 prochaines années se tourneront vers de l'entreprenariat lié à l'écologie, mais pas forcément sous format de startup, ni de levée de fonds. Cela peut prendre d'autres formes : 
- entreprise classiques à visée écologique (donc non dirigé vers de l'hypercroissance)
- engagement politique
- assos
- nouvelles structures sociales type auto-gestion

# 2023-09-24 Les algorithmes de caching

#algorithme #code 

Les "cache remplacement policies" ou "cache algorithms" sont des algorithmes utilisés pour la gestion du cache.

Le cache est de la mémoire à accès rapide permettant d'accélérer les programmes. Cependant, dès que le cache est plein, il faut supprimer des paires clé-valeur, en essayant de conserver les plus utiles, c'est à dire les plus susceptibles d'être réutilisées par la suite. Cela permet d'augmenter le hit rate (proportion de fois où la clé recherchée est présente dans le cache)

Les principales cache eviction policies sont :
- Random Replacement (RR), utilisé dans certains processeurs ARM
- First In First Out (FIFO), une queue 
- Last In First Out (LIFO), une stack
- Least Recently Used (LRU): la valeur accédée depuis le plus longtemps est retirée (parfois en approximation)
- Most Recently Used (MRU): la valeur ajoutée en dernière est la première à être retirée (rarement utilisé)
- Least Frequently Used: la valeur accédée le moins souvent (avec un compteur pour chaque accès)

# 2023-09-25 La Norvège

#europe #voyage 

La Norvège est une monarchie parlementaire (comme le [[2023-09-09 Le Danemark|Danemark]]), dont le roi est Harald V. Ce pays scandinave compte 5.5M d'habitants et est le pays le moins densément peuplé d'Europe après l'Islande et la Russie.

Le pays a connu une expansion rapide après la 2nd guerre mondiale, et est aujourd'hui un des pays les plus riche du monde, avec une politique sociale très développée. La Norvège est classée parmi les 1ere sur l'Indice de Développement Humain (IDH), avec la [[2023-08-12 La Suisse|Suisse]].

La Norvège ne fait pas partie de l'[[2023-09-18 L'Union Européenne|UE]] (refus par referendum à deux reprises), mais paye tout de même des contributions afin de faire partie de l'European Economic Area (EEA), permettant aux Norvégiens de vivre et travailler partout en UE, et inversement. Elle fait également partie de l'Espace Schengen. La Norvège utilise sa propre monnaie, la couronne norvégienne (NOK).

C'est un des pays avec la plus forte [[2023-09-19 La pluviométrie|pluviométrie]], en raison de sa longue côte atlantique et de ses montages. Le pays compte près de 110 Fjords, longue crevasses d'eau de mer, creusées durant la dernière ère glacière par les glaciers.

La Norvège unifiée à partie du IXe siècle a été dominée par le Danemark dans un premier temps, puis par la Suède. Ce n'est qu'en 1905 que la Norvège redevient indépendante, sans effusion de sang.

Le pays possède d'importantes ressources de gaz et de pétrole et d'hydroélecticité. Le pétrole rapporte près de 150 Mrd de dollars/an au pays (en 2021). Ces revenus ont été investi à partir de 1990 l'étranger pour une somme totale d'environ 1000+ Mrd de dollars: c'est le Petroleum Fund (devenu Government Pension Fund-Global en 2006)

# 2023-09-26 Les catégories de chiens

#chiens

Les principales catégories des chiens sont :
- Groupe 1) Les chiens de bergers (berger allemand, berger des Pyrénées, berger australien, border collie). Facile à éduquer, intelligent, dynamique
- Groupe 2) Les chiens de garde (dogue allemand, bulldog, rottweiler, Saint-Bernard, doberman)
- Groupe 3) Les terriers (Jack Russell, Yorkshire, fox terrier). Capacité de traquer et débusquer du gibier, indépendant, fidèles, vivaces
- Groupe 4) Les teckels. petite pattes avec un long corps, vivacité d'esprit et dynamique, ni pereux ni agressif
- Groupe 5) Les spitz et primitifs: chiens nordiques ou chiens de traineaux. Plutot poil long. Instinct sauvage, fier et indépendant 
- Groupe 6) Les chiens courants (dalmatien, beagle). Capacité à pister et poursuivre le gibier. Endurance et courage. Utilisés pour la chasse à courre
- Groupe 7) Les chiens d'arrêt (epagneuls, braques, griffons). Flair très développé, marquent l'arrêt (se braquent) lorsque le gibier est repéré
- Groupe 8) Chiens rapporteurs de gibier (retriever, labrador, cockers). Sont utilisés pour ramener le gibier, très bon chiens de guide. Aiment rendre service, calme intelligent.
- Groupe 9) Chiens d'agrément (caniche, chihuahua, bichon, bouledogue français). Petite taille, affectueux, complices, infatigables
- Groupe 10) Lévriers (whippet, grey hound, saluki) chiens très rapides

# 2023-09-27 Unités de force, d'énergie et de puissance

#sciences #energie 

En mécanique, l'unité de la Force est le Newton, `N`, qui correspond à des `g.m/s^2`. Pour le retrouver, il suffit d'utiliser le Principe Fondamental de la Dynamique (PFD): 
```
masse (en g) * accélération (en m/s^2) = somme des forces (en N)
```

En mécanique, le travail (W) est l'application d'une Force sur une distance. Un travail correspond à une quantité d'énergie, en Joules (`J`).

```
travail (en J) = force (en N) * distance (en m)
```

Autrement dit, en mécanique, l'énergie en Joules (`J`) a pour unités.

```
Unité d'Energie: J = g.m^2/s^2
```

L'énergie peut être vue comme un stock, donc une quantité finie. La puissance (P en Watt) mécanique est un flux: c'est la réalisation d'un `travail` par unité de temps.

```
puissance (en W) = travail (en J) / temps (s)
```

L'unité de la puissance peut être exprimée de cette façon :

```
Unités de Puissance: W = g.m^2/s^3
```

On peut aussi dire que l'énergie est une puissance pendant un temps donné :

```
Energie (J) = Puissance (W) * temps (s)
```

Une métaphore pratique est celle de la rivière : Si le débit de la rivière correspond à la puissance (gros débit = grosse puissance), alors l'énergie correspond à une quantité d'eau qui passe par une section de la rivière, pendant un temps donné.

L'énergie et la puissance en mécanique ont leur équivalent en électricité: Puissance en `watt`
et Energie en `watt.heure` (ou [[2023-07-25 Un kilowattheure|kilowattheure]])

# 2023-09-28 Les saunas

#histoire #europe 

Les saunas sont des traditions familiales depuis plus de 2000 ans en Finlande et Estonie. Ils sont secs (3 à 30% d'humidité) avec des températures allant de 70 à 100°C. 

Traditionnellement, il faut être nu dans le sauna, pour des raisons d'hygiène, avec une serviette pour se protéger du bois chaud. Dans certains pays du nord et germaniques, le port du maillot est souvent interdit.

La pratique du sauna a des vertus : elle apaise les douleurs articulaires, diminue les risques cardiovasculaires, etc.

En Finlande, il y a 2 à 3 millions de saunas, pour seulement 5 millions d'habitants. Ce pays organisait également des compétitions de saunas entre 1999 et 2010, année où un finaliste est mort ce qui a mis un terme à ces compétitions.

# 2023-09-29 L'exploration polaire

#histoire #nature #geographie 

Les expeditions polaires, au début du XXe siècle, avaient pour but d'atteindre pour la première fois les pôles Nord et Sud.

En 1908, l'américain Frédéric Cook (avec deux Inuits) annonce avoir atteint le pôle nord, mais cela a été remis en doute.

En 1909, Robert Peary et Matthew Henson (avec 4 Inuits) atteignent réellement le pôle nord.

Le norvégien Roald Amundsen, qui préparait une expédition vers le pôle nord, change d'objectif et prend la direction du pôle sud. Il annonce le changement de direction à son équipage seulement après avoir quitté le port. Il atteint le pôle sud en décembre 1911

Cinq semaines plus tard, Robert Falcon Peary et son équipe atteignent également le pôle sud, mais décèdent sur le trajet du retour.

Amundsen survola le pôle nord en 1926 en dirigeable. Il est donc le premier à avoir atteint les deux pôles.

# 2023-09-30 Quicksort

#algorithme #code 

Quicksort est un des meilleurs algorithmes de tri. Il a de meilleures performances que [[2023-07-15 Insertion Sort|Insertion Sort]] et que Sélection Sort, et n'a pas besoin de mémoire supplémentaire, comme [[2023-07-17 Merge Sort|Merge Sort]].

Quicksort a un average case en `O(nlog(n))`, lorsque l'array est aléatoire. Son worst-case arrive si l'array est déjà trié, et dans ce cas il est `O(n^2)` (il est donc conseillé de Shuffle l'array avant de faire un Quicksort). Il est `O(n)` en espace.

Pour faire un Quicksort :
1. Choisir le dernier élément l'array comme pivot
2. Pour chaque élément de l'array, le comparer au pivot. Si c'est plus petit ou égal, placer à gauche. Sinon placer à droite. (Bien garder les indexes de séparation de chaque partition)
3. Appliquer Quicksort récursivement sur chaque partition (gauche et droite)

# 2023-10-01 La Brandenburger Tor

#histoire #europe 

La Brandenburger Tor est un monument à Berlin construit vers l'an 1790. La construction en grès est réalisée dans un style Néoclassique, utilisant l'[[2023-07-01 L'Acropole d'Athènes|Acropole]] comme inspiration. Elle est couronné d'un quadrige figurant la déesse de la Victoire.

Le monument a été commandité par Frédéric-Guillaume II, le roi de Prusse.

En 1806, Napoléon Bonaparte emporte la statue vers Paris. Mais suite à la chute du 1er Empire, la statue est restitué sur la Brandenburger Tor, restauré et agrémenté de l'aigle prussien.

De 1961 à 1989, la Brandenburger Tor est un point de passage du [[2023-10-02 Le Mur de Berlin|Mur de Berlin]], se situant au milieu du No Man's land de la zone est de Berlin.

# 2023-10-02 Le Mur de Berlin

#histoire #europe 

Le Mur de Berlin a été érigé pendant la nuit du 12 au 13 août 1961 en plein centre de Berlin, pour mettre fin a l'exode de la population berlinoise vers Berlin Ouest, constitué des secteurs américains, britanniques et français. Cette séparation est la cristallisation de la guerre froide entre les deux super puissances de l'époque : le bloc soviétique et les US.

Le Mur a séparé la ville jusqu'à sa chute le 9 novembre 1989, suite à l'affaiblissement de l'Union Soviétique, alors dirigée par Gorbatchev.

Quelques points de passage, tel que [[2023-10-01 La Brandenburger Tor|la Brandenburger Tor]] et Checkpoint Charlie permettait aux Allemands de l'ouest de se rendre à l'est, mais l'inverse n'était généralement pas permis.

A Checkpoint Charlie, en 1961, des tanks soviétique et américains se sont fait face pendant 24h avant de se retirer sans aucun tir.

![[Pasted image 20231005083619.png]]

# 2023-10-03 L'absinthe

#nutrition #legal #histoire 

L'absinthe (aussi appelée la "fée verte") a été inventée en Suisse a la fin du XVIIIe siècle. Il était initialement utilisé comme remède, avant que la recette soit achetée par la famille Pernod pour la commercialiser en apéritif.

Les ingrédients principaux sont :
- l'anis
- le fenouil
- la mélisse
- l'hysope

Ces ingrédients sont macérés puis distillés pour obtenir l'absinthe blanche, puis macérés de nouveau pour obtenir l'absinthe verte.

L'absinthe a été interdite en [[2023-08-12 La Suisse|Suisse]] en 1910, suite a une épidémie d'alcoolisme et au lobbying de certains intellectuels et associations anti-alcool (la croix bleue). Suite à cette interdiction, le pastis marseillais est inventé.

En 2001, l'absinthe est de nouveau autorisé en France, mais a été fortement reglementé (taux d'alcoolémie, taux de thujone). Les distilleries principales se trouvent a Pontarlier.

En [[2023-10-04 La République Tchèque|République Tchèque]], le taux de thujone n'est pas règlementé.

# 2023-10-04 La République Tchèque

#europe #histoire 

La République Tchèque (en forme courte Tchéquie) est un pays de l'[[2023-09-18 L'Union Européenne|UE]], dont la capitale est Prague. Elle est composée principalement de la région de la Bohême et de la Moravie.

Au cours de son histoire, la Tchéquie n'a que rarement été indépendante.
- En 895, elle devient État vassal de la Francie Occidentale
- En 1002, elle devient État impérial du Saint-Empire Germanique 
- Après la défaite de la Révolte de Bohême (où il s'est produit des défenestration), elle est intégrée à la monarchie des Habsbourgs.
- Après la dissolution de du Saint Empire (180X ?), elle fait partie de l'empire d'Autriche, puis d'Autriche-Hongrie.
- Après la 1ere guerre mondiale, elle est fusionnée avec la Slovaquie, pour devenir la Tchécoslovaquie, et connaît une courte période démocratique entre les deux guerres mondiales.
- Elle est occupée par les Nazis entre 1938 et 1945
- En 1948, elle devient une république communiste.
- En 1989, la Révolution de Velours restore une démocratie dans le pays.
- La Tchécoslovaquie est dissolue en 1993 et la République Tchèque contemporaine est créé.

# 2023-10-05 Sissi l'impératrice

#histoire #europe 

Elisabeth de Wittelsbach (surnommé Sissi), nait duchesse en Bavière et est élevée à l'écart de la court dans une famille aimante et proche de la nature. Elle devint impératrice d'Autriche suite à son mariage à l'âge de 16 ans avec François-Joseph 1er, héritier de la dynastie des Habsbourgs.

C'était initialement sa sœur ainée, Hélène, qui était promise au jeune empereur. Mais subjugué par la beauté et charme d'Élisabeth, il choisi finalement de se marier à cette dernière.

N'arrivant pas à s'adapter à la vie à Vienne, Sissi passe une grande partie de sa vie à voyager et à faire du sport (équitation, gymnastique, marche, etc). Elle apprend plusieurs langues, collectionne l'art antique, rédige des poèmes... Anorexique, elle pesait 45kg pour 1m72. Mélancolique, voire dépressive, Sissi a écrit beaucoup de poèmes ce sujet.

Son assassinat (par un [[2023-09-05 L'anarchie|anarchiste]]) en 1898 a contribué à la légende du personnage, représenté dans le littérature et le cinéma. Malgré son manque d'intérêt pour sa vie de couple et la court, elle était profondément aimée par François-Joseph.

# 2023-10-06 les CRDTs

#algorithme #technologie 

Un Conflit-Free Data Type est une structure de donnée répliquée sur plusieurs machines via un réseau, avec quelques caractéristiques :
- une update locale est toujours possible, sans se coordonner avec les autres replicas
- un algorithme résoud toutes les inconsistances de données entre les replicas 
- les replicas convergent toujours vers un même état

Les CRDTs sont utilisés pour des outils collaboratifs temps réel, tel que Google Docs ou Figma.

https://jakelazaroff.com/words/an-interactive-intro-to-crdts/

# 2023-10-07 La Hongrie

#histoire #europe 

La Hongrie est un pays de l'[[2023-09-18 L'Union Européenne|UE]] donta capitale est Budapest.

La Hongrie fait partie de l'empire d'Autriche jusqu'en 1867, où la signature du Compromis transforme l'empire d'Autriche en Empire Austro-hongrois, où une certaine autonomie est permise sur le territoire hongrois.

L'empire Austro-hongrois est disloqué a la fin de la 1ere guerre mondiale, en 1918.

# 2023-10-08 Topological Sorting

#algorithme #code 

Un Directed Acyclic Graph (DAG) est une structure de donnée en graph, où chaque "vertex" possède une ou plusieurs "edges" unidirectionnels vers un autre vertex. Il n'y a pas de cycle dans un DAG.

Les vertices d'un DAG peuvent être triés avec l'algorithme Topological Sort. Cet algorithme ordonne les vertices tel que tous les vertices qui ont une edge vers un vertex `v` se trouvent avant `v`.

L'algorithme fonctionne de cette manière :
1. Créer un array `in_degree` qui compte le nombre de edges qui entrent dans chaque vertex `v`
2. Prendre n'importe quel vertex `v` tel que  `in_degree[v] = 0`, et l'ajouter au résultat.
3. Pour chaque edge sortant de `v`, décrémenter le `in_degree` de tous les vertices rencontrés

A la fin, `result` contient tous les vertices du graph, en topological order.

# 2023-10-09 Le democracy index

#sociologie #politique 

Le Democracy Index est un indicateur mesurant la démocratie dans 160 pays du monde, compilé par the Economist Intelligence Unit (UK).

Cet indicateur est créé à partir de 60 questions sur la démocratie, dont les réponses sont données par des experts et des questionnaires à la population.

Le Democratie Index catégorise les pays en:
- Full démocratie ([[2023-09-25 La Norvège|Norvège]], [[2023-08-12 La Suisse|Suisse]], France, ...)
- Flawed Démocratie (USA, [[2023-10-04 La République Tchèque|Tchèque]], [[2023-10-07 La Hongrie|Hongrie]])
- Hybrid Régimes (Tunisia, Ukraine, Mexico)
- Authoritarian régimes (China, Russia, Afghanistan, Myanmar)

# 2023-10-10 Les Révolutions Russes

#histoire 

Les révolutions russes de 1905 et 1917 sont parvenues à mettre fin au règne du Tsar Nicolas II (empereur de Russie).

Dans un contexte économique très difficile et un régime autoritaire, la population se revolt contre le régime en 1905. Une constitution libérale fut octroyée, mais en pratique, elle ne changea rien à la condition du peuple.

En 1917, une seconde révolution abouti à l'abdication du Tsar et la chute de l'empire impérial. Cette année marque aussi le début de la guerre civile, opposant
- l'armée rouge des bolcheviks de Lenin et Trotsky
- et l'armée blanche, rassemblant tous les opposants aux bolcheviks.

Cette guerre civile abouti sur une victoire bolchevik, sur une grande famine et une économie russe particulier fragile. Cette période marque la naissance du communisme et la création de l'URSS.

Joseph Staline, membre du mouvement bolchevik, parvient à s'accaparer tous les pouvoirs après la mort de Lenin en 1924 et à instaurer un régime totalitaire à partir des années 1920. Un régime bien éloigné des idéaux communistes de Lenin, issues des idées de "Das Kapital" de Karl Marx.

# 2023-10-11 Vitesse des chevaux

#equitation 

La vitesse moyenne d'un cheval est de
- 7 km/h au pas
- 15 km/h au trot
- 20 à 30 km/h au galop

# 2023-10-12 Les livres sacrés

#histoire #religion 
### La Torah
La Torah est, selon la tradition juive, les enseignements de Dieu transmis à Moïse sur le Mont Sinaï.

Elle est composée de 5 livres : la Genèse, l'Exode, le Lévitique, les Nombres et le Deutéronome.

La Torah est la première partie du Tanakh, qui regroupe la Torah, le Nevi'im et le Ketouvim.
### La Bible
La Bible chrétienne est composée de l'Ancien Testament (VIIIe au IIe BC) et du Nouveau Testament (fin du Ier siècle).

L'Ancien Testament reprend nombreux textes du Tanakh (bible hébraïque) et traite de tout ce qu'il s'est passé avant Jésus Christ : la création du monde, le Dieu unique, les 10 commandements.

Les textes du Nouveau Testament ont été rédigés par les apôtres et relèvent de la vie du Christ et de ses enseignements.
### Le Coran
Le Coran est le texte sacré de l'islam écrit entre 610 et 632. Selon la tradition musulmane, il regroupe la parole d'Allah mot pour mot, telle que transmise au dernier prophète Mahomet par l'ange Gabriel.

Le Coran reconnaît et reprend de nombreux textes de la Torah et de la Bible, mais considère que les versions récentes ont été altérées et falsifiées par l'homme.

# 2023-10-13 Algorithme de Dijkstra

#code #algorithme 

L'algorithme de Djisktra permet de trouver le chemin le plus court depuis une source S d'un weighted directional graph.

Pour fonctionner efficacement, l'algorithme de Djisktra nécessite une priority queue.

Les étapes sont :
1. Créer un array `shortest`, où `shortest[v]` est la distance jusqu'à `v`.
2. Mettre toutes les distances à +infini, sauf la source S, qui est à 0.
3. Dans une loop, à l'aide de la priority queue, prendre le vertex le plus proche, le sortir de la priority queue, et mettre à jour tous ses voisins.

Avec `n` le nombre de vertices et `m` le nombre de edges dans le graph, on dit que :
- Le graph est "dense" si chaque vertices est connectée à presque tous les autres vertices. Dans ce cas, `m` est proche de `n^2`
- Sinon, le graph est "sparse": `m` est du même ordre de grandeur de `n`.

Il existe plusieurs implémentations possibles pour une priority queue, et la meilleure pour l'algorithme de Djisktra dépend du graph : si il est dense ou sparse.

Pour un graph dense, un priority queue qui utilise un simple array, non trié sera en `O(n^2)` au total: l'opération "extract min" prend n opérations, dans un loop qui est effectuée n fois.

Pour un graph sparse, une priority queue implémentée en `binary heap` sera en `O(m*log(n))`, donc `O(n*log(n))`, car les opérations sur un binary heap prennent `log(n)` en moyenne.

Pourquoi ne pas utiliser un binary heap pour un dense graph ? Car dans ce cas, l'étape qui diminue les shortest de tous les voisins du vertex sélectionné prend `O(m*log(n)` au total, soit `O(n^2*log(n))`. Ce qui est pire que le `O(n^2)`d'un simple array.

# 2023-10-14 Les Heaps

#code #algorithme 

Les heaps sont des structures de données en tree satisfaisant la "heap property" : chaque node parent est plus petit ou égal à ses nodes enfants.

Il est possible de représenter une heap avec un Binary Tree, ou avec un array. Dans l'array:
- le 1er élément de l'array est le root
- les enfants d'un node `i` sont `2i` et `2i + 1`
- le parent d'un node `i` est `i/2`

Une Heap permet d'implémenter un Priority Queue dont les opérations sont en `O(log(n))`, utile notamment pour l'[[2023-10-13 Algorithme de Dijkstra|l'algorithme de Dijkstra]].

Cela permet également de réaliser un "heap sort", un sort en `O(n*log(n))` qui n'utilise pas de mémoire supplémentaire.

# 2023-10-15 La Bulgarie

#europe #histoire 

La Bulgarie est un pays de l'UE (depuis 2007), mais qui ne fait pas partie de l'espace Schengen. La population de 6.5M d'habitants ne cesse de diminuer depuis 1990, date de sa sortie du bloc soviétique.

L'alphabet officiel est le cyrillique, qui prend ses racines sur ce territoire et la religion la plus commune est le christianisme orthodoxe (70%) et l'islam (10%).

Le régime politique actuelle est une république, reconnue pour parmi les plus corrompues d'Europe.

L'histoire de la Bulgarie connaît différentes phases.
- Dans la période antique, le territoire est occupé par les Tharciens
- Les Proto-Bulgares envahissent le territoire et forment le 1er Empire Bulgare au VIIe siècle.
- La Bulgarie est conquise par l'Empire Byzantin au XIe (d'où la prédominance du christianisme orthodoxe)
- Le 2eme Empire Bulgaria nait en XIVe suite à des révoltes
- Au XVIe, le territoire est envahi par l'Empire Ottoman (turcs), apportant l'Islam
- La guerre Russo-Turque de 1878 abouti à la création de la Principauté de Bulgarie (au sein de l'Empire Ottoman), avec comme chef le Prince Alexandre I
- Le traité de Berlin en 1879 divise la Bulgarie en deux territoires, puis, suite à des révoltes le territoire est réunifié (1985)
- En 1908, la Bulgarie déclare son indépendance de l'empire Ottoman
- Lors de la 1ere guerre mondiale, la Bulgarie s'allie à l'Allemagne dans l'espoir de récupérer des territoires
- Lors de la 2nd guerre mondiale, elle rejoint l'Axe mais limite la déportation de sa population Juive.
- En 1946, un coup d'état instauré un régime soviétique, jusqu'à la dissolution du bloc en 1989
- La Bulgarie rejoins l'OTAN puis l'UE en 2007

# 2023-10-16 Les interfaces en Go

#code #go

En Go, les interfaces permettent de définir des actions possibles sur un type. Par exemple :

```
type Greeter interface {
  fun Greet()
}
```

Tout type qui implémente `Greet()` sera automatiquement un `Greeter`, sans avoir à le préciser dans le code.

### Concrete et interface type
Il y a une différence en Go entre 
- un "concrete type", qui est un type qui peut être instancié et qui contient une valeur
- un "interface type" qui est un type abstrait qui ne peut donc pas être instancié directement

### Conversion et assertion 
Il est possible de convertir un type vers un autre avec la syntaxe: `MyType(value)`. 

Il est également possible de faire un "type assertion" avec `v, ok := value.(MyType)`, ou encore un "type switch" avec

```
switch v := value.(type) {
case TypeA:
  (...)
case TypeB:
  (...)
}
```

### Empty interface
Un cas spécial est l'"empty interface" `interface{}` (aka `any`). Elle est implémentée par tous les types (par définition). L'empty interface est une sorte de "value box", permettant d'accepter n'importe quel type, mais cela demande par la suite de faire des toute assertion ou de faire des toute switches pour utiliser la variable.

# 2023-10-17 Le principe de Robustesse

#code #philosophie 

Le principe de robustesse, ou encore Postel's Law est:

> Be liberal in what you accept, and conservative in what you send.

Le principe a été pendant la création du protocole TCP, et a pour objectif de maximiser la compatibilité dans un réseau. Ce principe peut être appliqué à de nombreux autres cas (parsing d'HTML, user input, etc)

Le principe est parfois critiqué car ne permet pas de détecter d'éventuels bugs directement, seulement quand c'est déjà trop tard.

# 2023-10-18 Wave Function Collapse

#code #algorithme #graphics 

Wave Function Collapse est un algorithme procédural permettant notamment de générer des images ou des terrains de jeu en 2d ou 3d. L'algorithme a été inventé par https://twitter.com/exutumno et est utilisé notamment dans Townscaper d'Oskar Stålberg. 

L'algorithme se base sur des tiles, qui peuvent ou non êtres mises côte à côte. Les étapes sont :
1) Créer toutes les règles entre les tiles. Ex: Tile A peut être placée à gauche de Tile B, etc.
2) Sur une grille vide, chaque case peut contenir n'importe quelle Tile. Choisir une case et la fixer sur une Tile prise au hasard.
3) Ce choix réduit les possibilités des tiles voisines : calculer toutes les consequences découlent de ce choix (on dit que la wavefunction collapse ses possibilités)
4) Recommencer en prenant la case avec le plus faible entropy (le moins de possibilités) jusqu'à ce que toutes les cases aient "collapse" (ou qu'il y ait une contraction)

# 2023-10-19 La liberté d'expression

#france #europe #legal 

La liberté d'expression est inscrite dans la Constitution des Droits de l'Homme et du Citoyen de 1789 comme tel:
> Article 11. La libre communication des pensées et des opinions est un des droits les plus précieux de l'Homme : tout Citoyen peut donc parler, écrire, imprimer librement, sauf à répondre de l'abus de cette liberté dans les cas déterminés par la Loi.

Elle est indissociable à la liberté d'opinion:
> Article 10. Nul ne doit être inquiété pour ses opinions, mêmes religieuses, pourvu que leur manifestation ne trouble pas l'ordre public établi par la Loi.

La liberté d'expression a été redéfinie dans la Convention Européenne des Droits de L'Homme. 

Cette liberté comporte cependant des limites, qui sont :
- les propos racistes et la discrimination (homophobie, antisémitisme)
- les propos incitant à la haine, à la violence et au meurtre
- les propos defendant ou justifiant les crimes de guerre, les crimes contre l'humanité ou le terrorisme (ex: le négationisme)
- les propos diffamatoires (càd: porter atteinte à l'honneur ou à la considération d'une personne en affirmant un fait précis, qui n'a pas été vérifié au préalable)

# 2023-10-20 Le nombre d'or

#mathematiques #sciences 

Le nombre d'or est un nombre irrationnel égal à `(1+√5)/2 ~= 1.6180`.

C'est une des deux solutions à l'équation `x = 1 / (1 + x)`

Ce nombre, vu comme un ratio, intervient dans la construction du Pentagone régulier, dans la suite de Fibonacci ainsi que dans la nature (taille des feuilles sur un rameau), dans l'architecture, dansa musique et la peinture.

La première mention de ce ratio remonte à Euclide, dans son traité de mathématiques les Élements.

# 2023-10-21 La Convention Européenne des droits de l'Homme

#legal #europe 

La Convention Européenne des Droits de l'Homme et des libertés fondamentales est une convention adoptée en 1950 par les 46 pays membres du Conseil de l'Europe (y compris la Turquie). Elle garantie les droits et libertés des 700 millions d'habitants des états membres. 

Une institution européenne, la Cour Européenne des Droits de l'Homme, composée de 46 juges indépendants issus des états membres, peut être saisie si une violation aux droits fondamentaux est subie, même par des individus. La Cour siège aujourd'hui à Strasbourg.

# 2023-10-22 Les dates ISO8601

#code #geographie #web 

ISO8601 est un standard pour représenter les dates et le temps, proposé par l'International Standardization Organisation (ISO).

Dans ce format, les unités de temps les plus importantes se trouvent au début, pour qu'un simple tri lexicographique des dates les rangement dans le bon ordre (hors time zones différentes).

Les dates sont représentés sous forme
- extended: `2023-10-22`
- ou simple: `20231022`

Le temps est représenté avec un T au debut et, par défaut, est considéré dans la timezones locale (sauf précision contraire) :
- extended: `T16:13:00` (pour 16h13 et 0 secondes)
- simple: `T161300`

Pour préciser qu'il s'agit d'une date en "Coordinated Universal Time" (UTC) (heure de Londres), il faut rajouter un Z à la fin: `T16:13:00Z` 

Si l'on représente le temps dans une autre timezones, il faut ajouter `+hhmm` ou `-hhmm`. Par exemple
- pour UTC+2 (ex: Athènes) `T16:13:00+0200` 
- pour UTC-8 (ex: San Francisco) `T16:13:00-0800`

Une datetime peut être représentée en collant une date et un temps. Par exemple `2023-10-22T16:13:00+0300` en version extended et `20231022T161300+0300` en version simple.

# 2023-10-23 Les microplastiques

#ecologie #nature 

Les microplastiques sont des particules >5mm de matière plastique dispersées dans l'environnement. Ils s'accumulent dans les sols, les cours d'eau, les océans et certains aliments (ex: poissons). Ils sont retrouvés partout sur le globe: dans les glaces des sommets de montagnes, dans la calotte glacière de l'antarctique, etc. 

Leur effet sur l'homme n'est pas encore bien compris, car ils sont seulement étudiés depuis les années 2000.

Leur production provient de
- la dégradation des produits plastiques (textiles synthétiques, objets plastiques, pneus)
- les microbilles de plastiques utilisées en cosmétiques (>2% de la production)

Une partie importante des microplastiques se retrouvent dans l'environment via l'épandage des boues d'épuration (déchets des stations d'épuration d'eau).

# 2023-10-24 Le backtracking

#code #algorithme 

Le backtracking est un algorithme permettant de trouver un ensemble de solutions parmi un "state-solution tree" et avec un contrainte donnée.

L'algorithme permet de construire incrémentalement une solution en faisant des "choix" puis en dé-faisant ces "choix" (e.g backtrack) plus tard pour trouver d'autres solutions. D'une certaine façon, le backtracking est un Depth First Search (DFS) du solution tree.

L'algorithme a toujours cette forme:

```
void backtrack()
  if (solution is complete)
	  add solution to result
	  return

  for choice in possible solutions
	  if choice is valid
		  make choice
		  backtrack()
		  undo choice
```

# 2023-10-25 La mosquée bleue et l'Ayasofya

#histoire #europe 
## La mosquée bleue 
La mosquée bleue, aussi appelée mosquée Sultanahmet (en turc: Sultanahmet Camii) est une mosquée historique d'Istanbul. Elle fut construite par le Sultan Ahmet Ier entre 1606 et 1616, suite à la signature d'un traité de paix (temporaire) avec l'empire des Hasbourg en 1606.

La mosquée fut construite sur l'ancien site des grands palais des empereurs byzantins, face à la basilique Ayasofya.

## Sainte Sophie
L'Ayasofya, "Sagesse de Dieu", est une ancienne église construite sous l'empereur byzantin Justinien au IVe.
- A la prise de Constantinople en 1543 par les Ottomans, elle est transformée en mosquée
- Près de 400ans plus tard, en 1923, à la déclaration de la république turque, elle est transformée en musé et déclaré par la suite patrimoine mondial de l'UNESCO.
- En 2020, elle est re-devenue une mosqué

# 2023-10-26 Théorie de la relativité

#sciences #histoire 

La relativité restreinte, est une théorie élaborée par Albert Einstein en 1905. Dix ans plus tard, il publie la relativité générale, théorie relativiste dea gravitation qui s'appuie sur les concepts d'espace-temps et des espaces non euclidiens.

## Relativité Restreinte
Elle est basée sur deux principes :

1) Le principe de relativité
2) L'invariance de la vitesse de la lumière

Le principe de relativité indique que les systèmes physiques se comportent de la même manière dans tout référentiel galiléen (=inertiel). Cela signifie qu'il n'y a pas de référentiel "absolu": tout référentiel galiléen est équivalent.

Note: Un référentiel est galiléen si un point qui ne subit aucune force est en translation rectiligne uniforme dans ce référentiel.

L'invariance de la vitesse de la lumière indique que dans tout référentiel galiléen, la vitesse de la lumière est la même. Autrement dit : la vitesse de la lumière ne dépend pas de la vitesse de la source qui l'émet.

A partir de ces deux principes, Einstein explore le concept de simultanéité et arrive à la conclusion qu'il n'existe pas un temps absolu identique dans tout l'univers. Il est découle entre autre le phénomène de "ralentissement des horloges en mouvement".

## Relativité Générale
La relativité restreinte ne parvient pas à modéliser la gravitation. C'est en pensant à une personne en chute libre qu'Einstein eut un éclair de génie : comme une personne en chute libre ne sent pas son propre poids, c'est qu'elle doit être un référentiel galiléen !

Einstein posa alors le principe d'équivalence : un référentiel soumise a l'attraction gravitationnelle est équivalente à un référentiel en accélération "vers le haut". Les deux situations sont indissociables au niveau physique.

La gravitation n'est pas une force qui attire les corps entre eux, mais le résultat de la déformation de l'espace-temps par la masse des corps. A la proximité d'une masse importante (ex: le soleil) l'espace temps est déformé et une trajectoire rectiligne uniforme apparaît comme courbe pour un observateur extérieur.

Les géodésiques sont la trajectoire la plus courte d'un point A à un point B sur une surface donné. Les avions suivent des géodésiques du globales pour aller d'un aéroport a l'autre. Mais d'un point de vue extérieur à la terre, leur trajectoire apparaît comme courbe.

Avec la relativité générale, Einstein a fait une prédiction sur la deviation des rayons lumineux (pourtant sans masse) passant a proximité de corps célestes. Cette déviation, inexplicable par la mécanique newtonienne, a été observée expérimentalement dès 1918, en observant la position des étoiles lors d'une éclipse Solaire.

# 2023-10-27 L'empire byzantin

#histoire #europe 

L'empire byzantin est issue de la scission de l'empire romain en deux entre les deux fils de Théodose Ier en 395.

L'origine officielle de l'empire remonte quelques années plutôt tôt, en 330, lorsque l'empereur Constantin Ier deplace la capitale de l'empire à Constantinople (anciennement Byzance, aujourd'hui Istanbul).

L'empire romain d'occident s'éteint en 476, mais l'empire byzantin continue d'exister jusqu'en 1453, suite à la prise de Constantinople par les Ottomans.

Entre les deux, de nombreux événements marquent cet empire:
- Le règne de Justinien Ier au VIe siècle (avec l'impératrice Theodora), qui marque l'apogée territoriale de l'empire.
- la guerre contre les Perses et les invasions arabes au VIIe- VIIIe siècle.
- En 1054, le schisme de l'église catholique et de l'église orthodoxe.
- En 1204 la 4eme croisade, qui aboutit par le sac de Constantinople par les croisés (bien que initialement missionnés pour reprendre Jérusalem)

# 2023-10-28 Les clés ssh

#tools #web 

Secure Shell, ou SSH, est un protocole permettant la communication sécurisée entre deux machines. Elle utilise une paire de clés public-privé pour créer la connexion.

L'outil `ssh-keygen` est utilisé pour créer une SSH Key, qui est stockée par défaut dans `~/.ssh`. Ensuite, `ssh-copy-id` peut être utilisé pour copier la clé publique sur le serveur, qui se retrouvera dans le fichier `~/.ssh/authorized_keys` du serveur.

Les clés SSH peuvent être protégées par mot de passe, ce qui rajoute un niveau de sécurité supplémentaire. Pour éviter d'avoir à re-entrer le mot de passe a chaque connexion, la key débloquée peut être stockée dans un processus. Pour cela : `eval(key-agent)` puis `key-add` pour ajouter chaque key.

Un fichier `~/.ssh/config` permet de configurer ssh et de faciliter la connexion vers un serveur.

# 2023-10-29 L'empire Ottoman

#histoire #europe 

L'empire Ottoman a ete créé en 1299 en Anatolie, par le premier Sultan Osman I. Les ottomans se sont rapidement étendus, mettant notamment fin à [[2023-10-27 L'empire byzantin|l'empire byzantin]] en 1453 lors de la prise de Constantinople.

L'empire s'est étendu en Afrique du nord, en Grèce, dans les Balkans et en Asie mineure. Il atteint son âge d'or sous Suliman le Magnifique.

L'empire entre ensuite dans une période de déclin, et se fait dépasser par ses ennemis : l'empire de Russie et celui des [[2023-10-05 Sissi l'impératrice|Habsbourgs]]. 

Lors de la 1ere guerre mondiale, il s'allie à l'Allemagne pour éviter l'isolation politique et conduit le génocide Arménien, Assyrian et Grècque. Suite à la victoire alliée,e territoire est divisé et occupé, notamment par les Britanniques.

La guerre d'indépendance turque, menée par Mustafa Kemal Atatürk abouti à la création de la République de Turquie, le 29 octobre 2023, il y a exactement 100 ans.

# 2023-10-30 Cryptographie de courbes elliptiques

#code #mathematiques #cryptographie 

La cryptographie de courbes elliptiques est un algorithme de public key cryptographie basée sur des courbes elliptiques. Cette technique vient remplacer l'algorithme de chiffrement RSA.
## Chiffrement RSA
RSA est un acronyme pout pour "Rivest, Shamir et Adleman", les 3 chercheurs ayant découvert cet algorithme. Dans les grandes lignes, l'algorithme utilise la factorisation de nombres premiers comme fonction irréversible: multiplier deux nombres premiers est facile, mais retrouver les facteurs à partir du résultat est très difficile. 

Son utilisation est aujourd'hui encore très répandue, sur le web, le secteur bancaire, militaire, etc. Cependant, des attaques au chiffrement RSA (autrement dit: des algorithmes meilleurs que le brut force pour décrypter) ont été découvertes. Ce chiffrement n'est considéré comme sécurisé qu'avec des clés relativement longues, ce qui a représente un coût non-négligeable sur les performances.

## Courbes Elliptiques
Une technique plus récente, basée sur les courbes elliptiques, donne un bien meilleur encryptage, même avec des clés plus courtes. D'un point de vue mathématiques, ce sont les courbes du type `y^2 = x^3 + ax + b`. Sur une courbe elliptique, il est possible de définir une opération qui permet de se déplacer sur la courbe de manière très aléatoire, si bien qu'en arrivant sur un point donné, il est (presque) impossible de deviner le chemin parcouru. En utilisant un modulo et en la partie entière des valeurs, cela donne une bonne fonction asymétrique (facile à calculer dans un sens, difficile dans l'autre).

Les courbes elliptiques peuvent également être utilisées pour générer des nombres pseudo-aléatoires, mais leur performances sont bien moindre que d'autres méthodes.

# 2023-10-31 L'affaire Dreyfus

#histoire #france 

L'affaire Dreyfus est une affaire judiciaire et événement marquant de la 3e République. 

En 1894, le contre-espionnage français trouve un bordereau envoyé par un officier français aux Allemands. Il y a donc un traître dans l'armée. Après une courte enquête, l'officier Dreyfus, de confession juive, est jugé à huit-clos par l'armée, puis inculpé et envoyé dans une prison sécurisée (aux conditions très dures).

Cette inculpation est retransmise au public via un journal antisémite et déclanche une vague d'antisémitisme en France. 

Cependant, les proches de Dreyfus, sont convaincu de son innocence. Avec l'aide du nouveau directeur de l'espionnage français, qui découvre a quel point l'enquête a été menée à charge et de l'écrivain Emil Zola, qui rédige une lettre ouverte au président, avec le titre "J'accuse...", ils parviennent à obtenir un nouveau procès. 

La population française est fortement divisée entre les "antidreyfusards" et les "dreyfusards". Cette division affaibli fortement la 3e République.

Après la 2eme condamnation, Dreyfus est gracié puis finalement innocenté en 1906. 

Cet événement inspire le journaliste autrichien Théodore Herzl, qui considère que le peuple juif n'est pas en sécurité en Europe et dit fonder son propre état, idée fondatrice du sionisme.

# 2023-11-01 Empreinte carbone d'un Ferry

#voyage #ecologie 

Prendre un ferry plutôt qu'un avion est globalement moins émetteur en gaz à effet de serre, bien qu'il est difficile d'estimer précisément ces émissions. Elles dépendent de :
- la vitesse du ferry (émission sont au cube de la vitesse)
- le taux de remplissage du ferry
- l'utilisation de cabine (nécessite des ferries plus gros)
- le transport d'un véhicule, type voiture ou camion

Sur le simulateur de [Bon Pote](https://bonpote.com/ferry-calculer-en-1-minute-son-empreinte-carbone/), un trajet d'environ 300km réalisé en ferry (de nuit, sans voiture et sans cabine) plutôt qu'en avion équivaut à environ 10x moins d'émissions de Gaz à Effet de Serre (GES)

# 2023-11-02 L'Illiade d'Homer

L'illiade est une histoire attribuée à Homère, auteur Grec avant vécu 8e siècle avant JC, dont on ne sait pas grand chose. Il a également composé l'Odyssée, ces deux histoires orales ont été posés par écrit deux siècles plus tard.

L'illiade relate des évènements liés à la guerre entre les grecs et la ville de Troy, située en actuelle Turquie. 

Le mythe veut que cette guerre ait été déclenchée par l'enlèvement d'Hélène par un prince troyen, Paris. Hélène était alors la plus belle femme du monde et mariée à un roi un grec, Ménélas. Mais Paris ayant offert une pomme en or (la pomme de la discorde) à Aphrodite, cette dernière lui offrit l'amour d'Hélène.

L'illiade débute au milieu de la guerre, et raconte l'histoire d'Achille et plus largement les combats entre grecs et troyens. Les grecs n'arrivant pas à percer les défenses des troyens, ils mettent en place un stratagème : le cheval de Troy.

Les historiens ne sont pas certain de la guerre ait réellement eu lieu, ou si il s'agit purement d'un mythe. La ville de Troy a elle, bien existée.

# 2023-11-03 Les murs d'escalade en intérieur

#escalade #santé 

Même en intérieur, les prises d'escalade ou certaines parties de module peuvent se briser. Il faut faire particulièrement attention dans les vielles salles.

# 2023-11-04 Self-host et phishing

#web #technologie 

Héberger soit-même un service web (tel que plausible.io ou une instance Bitwarden) est possible, mais est généralement détecté comme "tentative de phishing" par Chrome. Il affichera un écran rouge annonçant que le site est dangereux : c'est la fonctionnalité Safe Browsing de Chrome. En effet, de son point de vue, le site est une copie d'un site connu, mais hébergé sur une adresse différente. C'est louche.

# 2023-11-05 La belote

#coinche

Les règles de la belote sont très proches de la coinche, avec quelques différences sur la phase d'annonce :

D'abord, seules 5 cartes sur 8 sont distribués aux joueurs, et la prochaine carte du paquet est retournée et annonce l'atout. Tour à tour, les joueurs annoncent si il prennent ou non. Celui qui prend récupère la carte d'atout retournée.

Si personne ne prend, un deuxième tour d'annonce laisse les joueurs choisir leur atout. Le dernier joueur à annoncer au 2eme tour peut voir ses 8 cartes avant d'annoncer son atout (il est obligé de prendre).

L'équipe qui prend doit faire au moins 82 points, ou 92 si la belote est en jeu. 

Si le contrat est fait, les points gagnés sont ajoutés au score total, si le contrat est chuté, c'est 162 pour la défense et 0 pour l'attaque (+20 si belote est en jeu)

# 2023-11-06 L'équilibre d'un navire

#sciences 

L'équilibre d'un navire est régi par deux forces : 
- la gravité, appliquée au centre de gravité du navire (point G)
- la force de flottaison, appliquée au centre de gravité de l'eau déplacée (point B)

Si le navire tangue d'un côté, le point B sort de l'axe de symétrie du navire. Le point d'intersection entre l'axe de symétrie du navire et la verticale du point B est un point M.

La distance GM s'appelle la "metadistance".
- si elle est positive, le navire retourne vers sa position d'équilibre
- si elle est négative, le navire se retourne
- si elle est nulle, le navire reste dans sa position (il ne retourne donc pas a la verticale)

# 2023-11-07 le Risorgimento italien

#histoire #europe 

Le Risorgimento est le mouvement d'unification de l'Italie du 19e siècle qui aboutit à une péninsule italienne unifiée et d'une monarchie constitutionnelle issue de la maison de Savoie.

Jusqu'à cette période, l'Italie était divisée en de nombreux royaumes et duchés, dont certains sous contrôle direct ou indirect des Habsbourgs, d'autres sous le contrôle du Pape (les états pontificaux) ou encore des Bourbons. 

Les conquêtes napoléoniennes ont rebattu les cartes, en annexant une partie du nord de l'Italie à la France en 1812 (Toscane, Piémont, Genoa), suivi du redécoupage de ces régions lors du congrès de Vienne et traité de Paris (1815, chute de l'Empire).

De 1859 à 1861, c'est la maison de Savoie, du royaume du Piémont-Sardaigne, qui parvient à expulser les Autrichiens, entre autre à l'aide de Napoléon III et d'un républicain Niçois : Garibaldi.

La proclamation du royaume d'Italie date de 1861, suivie quelques années plus tard en 1870 par l'annexion de Rome.

La monarchie constitutionnelle restera en place jusqu'à l'abolition de la monarchie et la naissance de la République italienne à la sortie de la 2nd guerre mondiale, en 1946.

# 2023-11-08 Pompéi

#histoire #europe 

Pompéi est une ville antique située à proximité de Naples et du mont Vésuve. La cité fut fondée 6e siècle avant JC, puis brusquement recouverte de cendres en 79, après l'éruption du volcan Vésuve.

Des fouilles archéologiques font ressurgir la ville, qui se trouve dans un état de conservation inespéré : on y retrouve des villas entières avec leur peintures murales, les rues pavées, fontaines, amphithéâtres, thermes mais aussi des vases, de la nourriture mais également les corps des habitants essayant de fuir la catastrophe ou de se protéger.

La ville ensevelie est doucement oubliée et protégée des pillages, jusqu'à ce qu'elle soit redécouverte au 17e siècle, par hasard lors de la construction d'un canal.

# 2023-11-09 Les noms de domaines les plus chers

#web #finance 

Les noms de domaines les plus cher du monde sont :
- voice.com pour $30M en 2019
- 360.com pour $17M en 2015
- nfts.com pour $15M en 2022

Quelques mentions honorables sont :
- hôtels.com pour $11M en 2001
- tesla.com pour $11M en 2014
- fb.com pour 8.5M en 2010
- icloud.com en 4.5M en 2011

# 2023-11-10 Shift your job

#france #ecologie

Une liste des entreprises ayant un impact positif sur l'environnement a été réalisée par le Shift Projet. La liste est disponible sur shiftyourjob.org.

Les différents secteurs sont:
- énergie
- alimentation, foresterie, utilisation des sols
- bâtiments et construction
- transport et mobilité
- biens et services 
- accompagnement et conseil 
- formation 
- informations et medias
- finance 
- réglementaire et institutions 
- recherche 
- élimination anthropique du CO2

# 2023-11-11 Spots de bloc en Europe

#escalade #europe 

Les meilleurs spots de bloc en Europe sont:
- Fontainebleau, France
- Albarracin, Espagne
- Magic woods, Suisse
- Val di Mello, Italie
- Zillertal, Autriche
- Peak District, UK
- North Wales, UK
- Kjugekull, Suède
- Goteborg, Suède
- Sintra, Portugal

# 2023-11-12 L'heure d'été

#france #legal 

La France est sur deux fuseaux horaires:
- En hiver (l'heure "normale"): UTC+1
- En été, UTC+1

Les changements d'heure se font toujours dans la nuit de samedi à dimanche, le dernier samedi du mois. En mars pour l'heure d'été et en novembre pour l'heure d'hiver.

L'heure d'été a été instauré une première fois de 1916 à 1944, puis a été ré-instaurée en 1975 (choc pétrolier) pour économiser du pétrole les soirs d'été, qui était alors utilisé pour l'éclairage. Cette mesure devait être provisoire mais est toujours appliquée.

Après une consultation publique à l'échelle de l'UE, plus de 80% des réponses étaient favorable à la suppression du changement d'heure. Une directive européenne devait être appliquée par le Conseil européen fin 2020... mais a été repoussé pour cause de crise sanitaire du Covid et n'est plus à l'ordre du jour depuis.

# 2023-11-13 Bonus Réparation

#ecologie 

Depuis cet automne, un bonus a la réparation des chaussures et des vêtements est disponible chez tous les réparateur labellisés Refashion. La liste est disponible sur https://refashion.fr/citoyen/fr/je-repare-bonus-reparation

# 2023-11-14 Le projet Presqu'ile à vivre

#lyon #ecologie 

La métropole de Lyon travaille sur un projet nommé "Presqu'Ile à vivre" et ayant pour but d'améliorer le cadre de vie au sein de la presqu'île.

Il s'agit à la fois de la création d'une zone a traffic limité ente Croix-Rousse et Perrache, de réaménager les quais du Rhône et les ponts mais également d'ajouter des espaces verts.

Le projet doit arriver à terme aux horizons 2030.

# 2023-11-15 L'Arcom

#france #legal #web 

L'Arcom, ou Autorité de Régulation de la COMunication audiovisuelle et numérique à été créé en 1989. Elle portait alors le nom du Conseil Supérieur de l'Audiovisuel (CSA), mais a été rebaptisée en 2020 pour refléter son rôle élargi danse domaine du numérique.

L'Arcom possède diverses missions, incluant :
- régulation des réseaux sociaux et medias numériques contre la désinformation, l'incitation à la haine, protection de la liberté expression, protection des mineurs 
- la protection des créations audiovisuelles françaises (quotas télévision et radio, lutte contre le piratage)
- la protection de la langue française
- la représentativité de la diversité de la population française, le droit des femmes, la lutte contre les discriminations

# 2023-11-16 Les Licoornes

#entreprenariat #france 

Les Licoornes est un groupe de coopérative française. Elles ont pour ambition de transformer radicalement l'économie au travers d'un idéal écologique, démocratique et solidaire.

Les principales sont les suivantes:
- Enercoop, coopérative d'électricité verte
- Telecoop, opérateur télécom
- mobicoop, mobilité partagée
- Commown, électronique responsable 
- Coop Circuits, circuits courts d'aliments et produits 
- [[2023-06-12 La NEF|La NEF]], banque éthique 
- Citiz, réseau d'autopartage
- Railcoop, opérateur ferroviaire 
- Label Emmaüs, E-Shop militant 
- Ethikdo, plateforme d'avantages et cartes cadeau 
- Windcoop, compagnie maritime

# 2023-11-17 Internet Relay Chat

#web #technologie 

Internet Relay Chat, ou IRC, est un protocol textuel pour la communication de groupe sous forme de channels. Créé en 1988 par Jarkko Oikarinen, IRC était très utilisé jusqu'au début des années 2010. En 2016, une standardisation des fonctionnalités modernes d'IRC a été réalisé sour le nom d'IRCv3, mais n'est pas implémenté par les principaux serveurs IRC.

# 2023-11-18 Asahi Linux

#technologie #hardware #unix 

Asahi Linux est un projet qui vise à porter Linux sur Apple Silicon M1. Ce projet est difficile car, bien que le dual-boot soit autorisé sur macbook (pas besoin de jailbreak), le hardware ne comporte aucune documentation. Cette équipe de bénévoles doit donc reverse-engineer l'ensemble des composants Apple pour faire un port Linux stable. La distribution linux supportée sera Arch Linux ARM.

L'équipe qui travaille sur Asahi Linux est composée entre autre de 
- Hector Martin "marcan", qui a déjà fait des port de LInux sur PS3 et PS4 et a contribué au Homebrew Channel sur la Wii.
- Alyssa Rosenzweig, GPU lead
- Asahi Lina, GPU kernel

# 2023-11-19 Loi anti-gaspillage

#legal #france #ecologie 

La loi anti-gaspillage est composée de 130 articles ayant pour objectif de lutter contre toute forme de gaspillage et renforcer la transition vers une [[2023-08-23 L'économie circulaire|économie circulaire]]. Elle cherche entre-autre à sortir du plastique jetable, renforcer le réemploi, lutter contre l'obsolescence programmée.

La loi prévoit la fin des emballages plastiques d'ici 2040 et inclue des restrictions tel que:
- 2023: interdiction de la vaisselle jetable dans les fast food
- 2024: interdiction de vendre de dispositif médicaux contenant des microplastiques
- 2025: les laves linges neufs doivent retenir les micro-plastiques via des filtres
- 2025: réduction de 20% des emballages plastiques à usage unique
- 2026: interdiction de vendre des cosmétiques contenant des microplastiques
- 2030: réduction de 50% des bouteilles en plastiques

# 2023-11-20 La guerre d'Espagne

#histoire #europe 

La guerre d'Espagne est une guerre civile qui s'est déroulé entre 1936 et 1939, opposants les forces républicaines anti-fasciste et les forces nationalistes, dirigées par Francesco Franco.

En plus d'une guerre de territoires, cette guerre s'est passée au niveau idéologique et politique. Le côté républicain regroupait des milices diversifiées tel que le parti travailliste, des trotskistes, des anarchistes, etc. Le côté nationalité fasciste a été soutenu par l'Allemagne nazie et l'Italie de Mussolini.

En 1939, Franco déclare sa victoire et installe une dictature fasciste jusqu'à sa mort, en 1975.

# 2023-11-21 Le rayonnement ultraviolet

#sciences #santé 

Les rayonnements ultraviolets sont des rayonnements électromagnétiques de longueur d'onde comprise entre 400nm et 100nm. Il se situent donc entre la bande de lumière visible et les rayons x. Les ultraviolets sont catégorisés en 3 bandes, de plus en plus énergétique.
- UV-A, ultraviolets proches: de 400 à 315nm
- UV-B, ultraviolets moyens: de 315 à 280nm
- UV-C, les plus énergétiques, de 280 à 100nm

Le soleil émet des ultraviolets (prêt de 5% de son énergie électromagnétique), mais la plupart sont filtrés par la couche d'ozone de l'atmosphère. 95% des ultraviolets atteignant la surface de la Terra sont des UV-A. Les UV ne sont pas bloqués par les nuages et ils ne procurent pas de sensation de chaleur sur la peau (qui provient des infrarouges). Les UV atteignant le sol sont en partie réfléchi, notamment par le sable, l'eau et la neige.

Une faible exposition aux UV-B permet à l'organisme de synthétiser la vitamine D. Cependant, une exposition prolongée est nocive: elle provoque un dessèchement et vieillissement de la peau et peut entrainer des tumeurs (cancer). L'exposition aux UV plus énergétiques (UV-B et UV-C) est évidemment plus dangereuse.

# 2023-11-22 MapLibre

#open-source #technologie 

MapLibre is an open-source fork of Mapbox, available on Github. It has a JS map renderer, a Native C++ renderer and an experimental Rust and WebAssembly renderer. It also contains a project called Martin, which is a vector tile server.

# 2023-11-23 OPA Hostile

#finance #entreprenariat 

Une OPA, Offre Publique d'Achat (ou "takeover"), est une offre d'achat d'une entreprise, pour en prendre le contrôle. Elle est dite "hostile" si elle faite contre la volonté du management. Une OPA hostile peut être réalisée de 3 manières :
- En achetant suffisamment d'actions en bourse, si l'entreprise est côtée 
- En faisant une offre publique d'achat des actions des actionnaires, à un prix supérieur au prix du marché (tender offer)
- En faisant une "proxy fight", qui vise à convaincre suffisamment de membres du board pour voter l'exclusion des membres reticent au takeover

Les OPA hostiles ont  peu de chances de succès et se terminent généralement par un accord entre les partis ou par un échec du takeover.

Quelques exemples d'OPA hostiles réussies :
	- Le rachat de Genzyme par Sanofi (2011)
	- Le rachat de Cadbury par Kraft foods (2011)
	- Le rachat de Twitter par Elon Musk (2022)

# 2023-11-24 L'usage raisonnable du forfait

#legal #europe #france 

L'Union Européenne permet aux opérateur telecom d'appliquer une politique d'utilisation "raisonnable" du roaming à l'étranger. C'est à dire que les opérateurs peuvent facturer davantage la data, les SMS et appels depuis d'autres pays de l'UE lorsque certaines limites sont dépassées.

Pour SFR, au bout de 3 mois passés à l'étranger, un quota du temps passé en France vs à l'étranger doit être rempli. Dans le cas contraire, des frais supplémentaires sont appliqués. Un SMS est envoyé 15 jours avant l'application de ce surcout.

# 2023-11-25 Période de nidification

#ecologie 

La période de nidification des oiseaux s'étend de mi-mars à mi-juillet. La taille des haies est interdite durant cette période.

Pour les grimpeurs, il est conseillé de bien choisir les falaises à escalader, pour s'assurer de ne pas déranger des nids d'oiseaux.

# 2023-11-26 UK vs US settings

#technologie #lifeprotips 

Il y a plusieurs différences entre le settings "UK" et "US" pour le langage du téléphone.

Le setting US implique:
- la semaine commence dimanche
- le format de date est MM/DD/YYYY
- les unités impériales sont utilisées (Fahrenheit, inches, feet, pounds, gallons)

Le setting UK implique:
- la semaine complète lundi
- le format de date est DD/MM/YYYY
- les unités sont le système métrique

# 2023-11-27 Swift et SwiftUI

#code #technologie 

Swift est un langage moderne et multi-paradigmes (procédural, orienté objet et fonctionnel) permettant de créer des applications sur l'écosystème Apple. Il est devenu le langage officiel des plateformes Apple, pour remplacer Objectve-C.

Le langage a été rendu open-source en 2015 sous une license Apache 2.0. Swift peut être compilé sur macOS, mais aussi sous Windows et Linux.

SwiftUI est un framework déclaratif qui vient remplacer l'ancien UIKit et AppKit (même si UIKit et AppKit sont encore supporté).

# 2023-11-28 Albarracin

#escalade #geographie 

Albarracin est un village d'Espagne, à mi-chemin entre Valencia et Zaragoza très connu pour ses blocs en grès rouge (red sandstone). La meilleure saison est de l'automne au milieu du printemps, car l'été il fait généralement trop chaud, malgré le fait que le spot se trouve à 1000m d'altitude.

La plupart des blocs sont entre 6A et 7b, avec quelques classiques tel que :
- El Jacuzzi (6a)
- Boulder King (6a)
- Minigruyere (6a)
- Minichakra (6a)
- El Chimpanzé (6b)
- Cazo Tanks (6b)
- El Crocodilo (6b)
- Seicerraro (6b)
- Tech Don Pepo (7a)

Pour dormir Albarracin, les meilleurs adresses sont 
- la Sandstone Guesthouse
- Cino Capicol
- le camping (tente ou bungalow)
- le Airbnb Don Pepo

# 2023-11-29 Ajax et Fetch API

#web #javascript #technologie 

Asynchronous JavaScript And XML (AJAX) est une technique permettant de modifier une page web sans avoir à la recharger complètement. Pour cela, une `XmlHttpRequest` est créée dans le javascript de la page et permet d'effectuer des requêtes supplémentaires au serveur, et éventuellement modifier le Document Object Model (DOM). Ces requêtes sont asynchrones et ne bloquent dont pas le reste de l'execution du Javascript.

AJAX est apparu à partir des années 1998 sur les navigateurs et depuis de nombreux frameworks ont été développés par dessus, tel que jQuery, Dojo ou plus récemment Axios.

Le AJAX fonctionne toujours aujourd'hui, mais il est remplacé petit à petit par la nouvelle norme Fetch API, plus moderne.

# 2023-11-30 Les permis de conduire

#legal #france 

Il existe différentes catégories de permis de conduire en France :
- catégorie A pour les motos et quads
- catégorie B pour les voitures
- catégorie C pour les poids lourd et transport de marchandises
- catégorie D pour le transport de personnes
- catégorie E pour les remorques

Il existe 3 permis différents pour la catégorie A:
- le permis A1, pour les 2-roues de moins de 125CC et 11kW
- le permis A2, pour les 2-roues de moins de 35kW
- le permis A, pour toutes les motos et quads

Depuis 2020, les permis motos nécessitent de passer un "code moto", et de réussir au moins 35/40 questions. Il y a ensuite une épreuve pratique sur circuit et en circulation.

Une "formation de 7 heures" permet à un détenteur du permis B d'obtenir le permis A1 (2h théorique, 2h sur circuit, 3h en circulation)

# 2023-12-01 La Conférence des Parties

#ecologie #legal 

La Conférence des Parties (COP) sont des assemblées annuelles d'états et organisations internationales pour fixer des règles communes. La COP climat s'intéresse aux questions d'environnement et de biodiversité.

Le protocole de Kyoto (signé en 1997, entré en vigueur en 2005 et l'ccord de Paris (signé et entré en vigueur en 2016) sont issus de la COP.

# 2023-12-02 Les types de boxe

#sport #competition 

Les principaux types de boxe sont :
- Boxe anglaise (le noble art), uniquement les poings et coups au dessus de la ceinture
- Boxe thailandaise (Muay-Thai), poings, pieds, genous, coudes
- Kickboxing (né au japon), issu du karaté et autorise les pieds et les poings
- La boxe française (la savate), avec les pieds et poings, mais pas les tibias
- La boxe américaine (full contact) - pied et poings, coups au dessus de la ceinture

A noter que le pugilat est un ancêtre de la boxe, pratiqué entre autre dans la Grèce antique.

# 2023-12-03 La laïcité

#france #legal 

La laïcité est un des principes fondateur de la République, et est inscrit dans la constitution. La loi de 1905 formalise la séparation de l'église et de l'Etat, condition essentielle. La laïcité garanti :
1) la liberté de conscience
2) l'égalité des citoyens quelle que soit leur croyance
3) la neutralité de l'Etat face à la religion
4) le libre exercice des cultes

La laïcité va de pair avec la liberté religieuse, qui est le droit d'exprimer sa religion, de la pratiquer ou de l'abandonner, dans le respect de l'ordre publique. La liberté religieuse est encadrée, notamment via deux lois qui ont fait débat :

1) Dans les primaires, collèges et lycées publiques, une loi de 2004 interdit tout signe religieux "ostentatoire". Les signes discrets restent tolérés.
2) Dans l'espace publique, [une loi de 2011](https://www.legifrance.gouv.fr/jorf/id/JORFTEXT000022911670) interdit la dissimulation de son visage, interdisant donc le port de la burqa

# 2023-12-04 Faire un espresso

#nutrition 

Pour faire un bon espresso, il faut :
- 7g de café fraîchement moulu, à la bonne taille de grain (plutôt fin), ou 18g pour un double
- un temps d'extraction de 25 secondes pour 30ml de café (60ml pour un double)
- réchauffer les tasses avant l'extraction 

Un espresso est "sous-extrait" si le café coule trop vite et que la couleur est rapidement blanche. Un espresso est "sur-extrait" si il coule au compte-gouttes et a un couleur très foncée.

# 2023-12-05 Arbres en espalier

#nature 

Les arbres en espalier sont des arbres taillés pour pousser "à plat" contre un mur. Cette technique datant du moyen age est utilisée par les horticulteurs permet de décorer les murs et prendre moins de place, notamment pour les arbres fruitiers.

# 2023-12-06 LLVM

#technologie #code 

LLVM (anciennement "Low Level Virtual Machine") est un compilateur modulaire utilisé pour compiler différents langages, tel que le C, C++, Objective-C, le Java, le Swift ou encore le Rust. Issu d'un projet de recherche, LLVM est arrivé en 2000 et est une alternative à `gcc`, entre autre.

LLVM est divisé en sous-projets, dont `clang`, un compilateur C, C++ et Objective-C.

Mis à part son architecture modulaire, LLVM s'est différencié par des features tel que la possibilité de faire une compilation Just In Time (JIT) ou encore sa Représentation Intermédiaire du code (IR), permettant une optimisation spécifique à chaque processeur.

# 2023-12-07 Déglacer

#cuisine

Déglacer en cuisine consiste à verser un liquide (vin blanc, vin rouge, bière ou autre) sur une cuisson (oignons, viande) afin de récupérer les sucs et en faire une sauce.

# 2023-12-08 Le label VertVolt

#energie #ecologie 

Le label VertVolt est géré par l'ADEME.

En Europe, les producteurs d'électricité verte vendent d'un côté leur électricité (~40€ par MWh) mais aussi des certificats (~ 0.5€ par MWh). Les fournisseurs peuvent acheter les certificats sans acheter l'électricité.

Le label VertVolt garanti que le fournisseur a acheté l'électricité a un producteur d'énergie renouvelable en France. Le label a un deuxième niveau (très engagé) si le fournisseur achète 25% de son énergie a des collectivités territoriales avec un gouvernance partagée.

# 2023-12-09 La truite de Schubert

#histoire #musique 

La Truite est une quintette composée par Schubert en 1819 mais publiée seulement après sa mort.

# 2023-12-10 Être contenu dans une forme

#algorithme 

Une des méthodes pour savoir si un point donné est contenu dans une forme (complexe) est de compter le nombre de fois que la frontière est traversée :

Depuis un point extérieur, traverser la frontière
- un nombre impair de fois: on est dedans
- un nombre pair de fois: on est dehors

# 2023-12-11 FUD

#politique #psychologie 

Fear, Uncertainty and Doubt (FUD) est une technique manipulative utilisée en marketing, politique, sales ou autre.

Exemple: Apple annonce que Bleeper Mini (iMessages sur Android) est dangereux pour la sécurité et la privacy des utilisateur d'iPhone et est susceptible aux hacker et techniques d'arnaques. C'est un technique de FUD pour justifier sa volonté d'interdire toute ouverture d'iMessages aux utilisateurs Android.

# 2023-12-12 Les ESNs

#web #entreprenariat 

Une Entreprise de Service du Numérique (ESN) est une société qui accompagne ses clients dans la réalisation de projets numériques, via du conseil, conception, développement, maintenance ou formation.

ESN est le nouveau nom donné aux SSII (Société de Service en Ingénierie Informatique)

# 2023-12-13 La bête du Gévaudan

#histoire #france 

La bête du Gévaudan est le nom donné au responsable d'une centaine d'agressions mortelles dans le Gévaudan (Lozère) entre 1764 et 1767.

Plusieurs pistes ont été émises :
- meutes de loups anthropophages
- un animal exotique (loup garou)
- le diable
- un vagabond
- un comte de la region

# 2023-12-14 Flutter 3.16

#flutter #code 

Flutter 3.16 possède de nouvelles features:
- Material 3 est activé par défaut
- Améliorations de l'édition de texte
- une preview d'Impeller sur Android
- un game toolkit

# 2023-12-15 NixOS

#tools #technologie 

NixOS est une distribution Linux et un système de configuration utilisant une approche déclarative pour configurer à la fois l'OS et les applications. NixOS utilise le Nix Package Manager afin d'installer, mettre à jour et manager du software de manière reproductible.

L'avantage principal est de pouvoir décrire complètement l'état souhaité du système dans un fichier de configuration, et de pouvoir distribuer cette configuration précise sur un grand nombre de machines.

NixOS a une proposition de valeur similaire à Docker, mais utilise une approche différente.

# 2023-12-16 L'amidon des pommes de terre

#cuisine #lifeprotips 

La pomme de terre est un aliment riche en amidon, un sucre contenu dans les végétaux. L'amidon a un pouvoir liant et il est possible de le rincer en passant les pommes de terre coupées dans l'eau.

Pour faire des frites croustillantes ou des pommes de terre sautées il font donc rincer les pommes de terre après les avoir épluchées et coupées. A l'inverse, pour faire un gratin ou une purée, inutile de rincer.

# 2023-12-17 Les algorithmes de pathfinding

#algorithme #code 

### Breadth First Search (BFS)

BFS l'algorithme de pathfinding le plus simple. Il utilise une `queue` pour ajouter les prochaines nodes à visiter, ainsi qu'un `set` pour marquer les nodes déjà visitées.

Si le graph est "unweighted", BFS permet de trouver le chemin le plus court vers une node donnée.
### L'algorithme de Dijkstra

L'algorithme de Dijkstra utilise une PriorityQueue plutôt qu'une simple queue. Il faut aussi prendre en compte le fait qu'une même node peut être visitée plusieurs fois: il faut re-visiter la node seulement si cela diminue la distance totale jusqu'à cette node.
### A-star

L'algorithme A* utilise un "heuristique", qui est une estimation de distance restante jusqu'à la destination. Dans la priority queue, la priorité d'une node donnée est la distance depuis la départ + l'heuristique.

`priority = costSoFar + heuristic(goal, next)`

De cette façon, A* se concentre en priorité sur les paths qui se rapprochent de la destination, contrairement à l'algorithme de Dijkstra qui cherche dans toutes les directions.

# 2023-12-18 Le TGV M

#france #voyage #ecologie 

Le TGV M, ou TGV modulaire, est la dernière génération de Train Grande Vitesse, conçu en partenariat entre la SNCF et Alstom. Ce nouveau TGV a été améliorer pour diminuer les coûts de fabrication et de service, entre autre via une forme plus aérodynamique de la motrice, et pour maximiser le comfort des voyageurs et du nombre de places disponibles.

Il est prévu que sa mise en service commence fin 2024.

# 2023-12-19 Les principaux LLMs en open source

#technologie #open-source 

Les 3 principaux LLM dit "open" sont:
- LlaMA 2 (par Meta)
- Mistral/Mixtral (Par Mistral AI, une startup française)
- BLOOM (BigScience Large Open-science Open-access Multilingual Language Model)

Ces modèles, ainsi que des variations, peuvent être téléchargés et utilisés offline, par exemple via le projet [Ollama](https://ollama.ai/library)

# 2023-12-20 The Shoelace formula

#algorithme #code 

L'algorithme "shoelace formula" est une formule simple permettant de calculer l'aire d'un polygone sur un plan cartésien. Pour que la formule fonctionne, le polygone peut avoir tout type de forme mais aucun segments ne doit se croiser.

La formule est la suivante :

```
aire = 1/2 * sum(x_(i) * y_(i+1) - x_(i+1) * y_(i))
```

Le nom de la formule vient du fait que les coordonnées sont "croisées", tout comme des lacets.

# 2023-12-21 Boyer-Moore algorithm

#code #algorithme 

Boyer-Moore est un algorithme permettant de faire de la recherche de texte rapide, bien plus rapide que la recherche exhaustive, car permet de skipper un bon nombre de comparaisons.

Soit `P` le pattern recherché dans le texte `T`. `T` est parcouru de gauche à droite, cependant la comparaison commence par la fin de `P` (donc de droite à gauche)

L'algorithme est basé sur deux règles permettant de skipper des alignements, et choisira la max entre ces deux règles. Les règles entrent en compte dès qu'un character est mismatched
1. La "bad character rule" permet de skipper des alignements jusqu'à la prochaine occurence du character dans le pattern
2. La "good suffix rule" permet de skipper jusqu'à ce que le même suffit soit trouvé dans le pattern

Pour que ces règles soient rapide à executer, une lookup table est calculée avant la search et permet d'appliquer chaque règles en O(1).

# 2023-12-22 Une attoseconde

#sciences 

Une attoseconde est la plus petite unité de temps mesurable. Par exemple, te temps de rotation d'un électron sur la première orbite Bohr est de l'ordre de 150 attoseconde .

Dans l'ordre du plus grand au plus petit :
- la seconde 10^0
- la milliseconde 10^-3
- la microseconde 10^-6
- la nanoseconde 10^-9
- la picoseconds 10^-12
- la femtoseconde 10^-15
- l'attoseconde 10^-18

# 2023-12-23 Le Gin

#nutrition 

Le gin est un alcool à base d'alcool éthylique d'origine agricole (alcool neutre) et de baies de genévrier (un conifère). Le gin prend ses origines aux pays-bas espagnols au cours du 17e siècle. Il en existe deux grandes variétés:
- Le London Dry Gin (gin anglais), peu aromatisé, donc facilement utilisé dans les cocktails
- Le Genièvre (gin hollandais ou belge), plus parfumé et rarement utilisé dans les cocktails

# 2023-12-24 Un atoll

#geographie #sciences 

Un atoll est une forme d'ile circulaire avec un lagon intérieur, comme trouvé par exemple en polynésie francaise, sur l'ile d'Hao.

Le processus de formation est long: une ile volcanique s'affaisse peu à peu via l'érosion, tandis que des bancs de sable s'accumulent sur les côtes. Si les bonnes conditions sont réunies, le centre de l'ile peut repasser sous le niveau de la mer, laissant uniquement le contour: c'est un atoll.

# 2023-12-25 Hyperscript

#code #web 

Hyperscript est un langage de scripting créé par le créateur de HTMX. Les scripts sont écrits directement dans les balises html et permettent de créer de l'interactivité, de gérer des évènements et autre.

# 2023-12-26 Oeuvres dans le domaine publique

#legal #europe 

Dans l'UE, les oeuvres entrent dans le domaine publique 70 ans après la mort de l'auteur (ou du dernier auteur survivant, si il s'agit d'une collaboration).

# 2023-12-27 Comment poser sa voix



# 2023-12-28 parser du JSON en Go

#code #go

Pour parser du JSON, il faut utiliser créer une struct avec la bonne forme, si l'on connait la structure du JSON par avance.

```
json.Unmarshal(inputBytes, &myStruct)
```

Si besoin, il est possible d'implémenter la fonction `UnmashalJSON` sur la struct.

Si la forme du JSON n'est pas fixée, il est possible d'utiliser une `map[string]any` puis de type cast les valeurs.

Si le json n'est pas valide, `json.Unmashal` va paniquer. Pour éviter cela, il est possible d'appeler `json.Valid(inputBytes)` auparavant.

# 2023-12-29 Le Commonwealth

#geographie 

Le Commonwealth est une organisation de 56 pays, majoritairement issus de l'empire britannique. Il est formellement constritué par la déclaration de Londres en 1949.

Les états n'ont aucune obligation mais partagent la langue, la culture et l'histoire. Ils ont également une charte qui promeut des valeurs tel que la démocratie, le droit des hommes et l'état de droit.

Les principaux pays membres sont:
- Royaume Unis
- Canada
- Nouvelle Zélande
- Afrique du Sud
- Pakistan 
- Inde
- Sri Lanka
- Australie 
- de nombreux pays d'Afrique

# 2023-12-30 Calculer son IMC

#santé 

L'indice de masse corporelle (IMC) est calculé avec la formule suivante :
```
Imc = poids / taille^2
```

Avec le poids en kg et la taille en mètre.

A partir de
- 25, on parle de surpoids 
- 30, on parle d'obésité
- 35, obésité sévère

# 2023-12-31 Vent onshore ou offshore

#nature 

Le vent offshore souffle depuis la terre en direction de l'océan. Ce vent creuse les vagues et si il est suffisamment puissant, il peut former des tubes

Le vent onshore souffle depuis l'océan vers la terre. Ce vent dégrade la qualité des vagues : il y a du clapot.
