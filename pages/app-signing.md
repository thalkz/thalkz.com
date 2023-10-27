[Back Home](/)
# App Signing and Apple Certificates
Created: 9 May 2020

While developing Flutter apps, one of the necessary steps is to get your iOS apps on the AppStore. And boy it is complicated. One of the most confusing things are thoses files you need to generate in order to publish your app. There are .12 files, .cer files, but also .provisioning files, Certification Requests, and so on. I’ve never quite took the time to fully understand the logic behind it, and usually scanned online tutorial in order to solve the problem I was facing.

But today, I will the time to the resarch necessary. Today is the day I understand.

## Pourpose of the signing an app
Okay, first things first. App signing solves two very important problems :

**Problem 1)** As a user of the app, I need to be certain that the app I’ve downloaded was actually published by the correct developer. If I download the Uber app, I need to be sure that it is accually Uber that published this particuliar app I’ve downloaded, and not anyone else (eg. an attacker trying to inpersonate Uber)

**Problem 2)** As a developer of an app, I need to control who can use my app. Selected testers only ? My whole organisation ? Everyone ? If I decide to share my app with a tester and him only, I want to know he will not be able to leak the app to anyone else.

To put it simply, Problem 1) is solved using certificates, while Problem 2) is solved with provisioning files. But for now, let’s focus on the signing process, provisioning files will be research for another time.

## The complete app signing process
Here is the situation. Your app is ready and compiled, and you want to publish it. But first you need to sign it. Here are the steps you need to follow :

**Step 1)** You generate a Certificate Request from your Keychain Access app on your Mac (yes, you need a Mac for that). This creates a .certSigningRequest file.

**Step 2)** You upload your Certificate Request to Apple’s website. Apple thanks you and gives you a “Certificate” —which is a .cer file— in exchange.

**Step 3)** You import your .cer file back into Keychain Access. You can now see that your “Certificate” is available in Keychain Access, along with a private key “attached” to it (We’ll go back to it later).

Congratulations, you now can sign your App in Xcode ! But wait, what if you want another developer on your team to be able to sign and publish the app ? Or your CD (Continuous Deployment) server to de able to, well, sign and deploy the app ? You’ll find out that simply sharing them the .cer file will not be enough. You actually need to export yet another file from Keychain access, a .p12 file. Which means :

**Step 4)** (Optionnal) From Keychain Access, which holds both your Certificate and associated private key, you export a .p12 file that you can share to give signing access to your app.

Note: the .p12 file is still encrypted by a password that you chose, but that’s not a very strong protection (it can be brute-forced). If your .p12 file leaks, you should revoke your Certificate and create another one.

**Step 5)** You can now “archive” your app (which signs it) and upload your app to the AppStore from this Mac, or from a mac that has imported the .p12 file.

**Step 6)** When a user downloads your app, their iPhone will automatically check that the app is correctly signed. In order to do this, it uses the certificate included in the app and check that the rest of the app is signed with this certificate. It it doesn’t, iOS won’t open the app.

Thankfully, once everyting is setup, Keychain Access keeps your Certificate and lets you sign your app anytime you want, withouyt generating a new one. Well up until the certificate expires (1 year) or you revoke or delete it, in which case you’ll have to restart from step 1.

Now, why are those steps necessary and how does it help to solve problem 1 ? To fully understand what happened in this signing process, let’s take a step back and talk about digital signing.

## Digital signing and key pairs
Digital sining is the concept of of signing a message (or file), just like when you sign a document on paper (but usually more secure). Digital signing looks a lot like encryption, but with the distinction that with digital signing, the message is still readable by anyone, whereas in encryption you need “decrypt” the message. Instead of encrypting & decrypting a message, digital signing is use for signing & verifying a message.

A central concept of digital signing is the public-private key pair. It’s a key “pair” because there are two keys, and each key can be represented as a string of characters. Each of these keys has its specific pourpose.

The private key is only stored on your machine and should not be shared with anyone. It is used to “sign” a message (in our case, an app) that can be verified with the public key.

The public key can be shared with anyone and has the ability to verify messages (in our case, apps) that were signed with the private key. It has the ability to say “yep, it’s signed with the private key” or “nope, not signed with the private key”. In all cases, the public key cannot sign an app and cannot be used to generate a private key.

*Note: How exactly “sign with private key” and “verify with public key” works is also research for another time. For now, let’s just assume private and public keys actually work as described.*

You can create key pairs from any computer and use them for signing & verification, but for signing an iOS app, not any public-private key pair will do. We’ll come to that later.

Okay, now let’s go back to the signing process and explain it again, this time armed with our knowledge about key pairs.

## How the app signing process works
In step 1, when generating a .certSigningRequest file, we actually also generated another file that stayed in Keychain Access. This other file is the private key. The .certSigningRequest contains both the generated public key as well as some information about you (name, email, etc). At this point, everything is still offline. Apple has no idea about your keys.

In step 2, you actually gave Apple a copy of your public key (which is ok, since public keys can be shared), and Apple saved it on their servers. You do not send your private key to Apple, and at no point will they receive it. It would completely compromise the process because it would then be impossible for the user to know for sure that it is actually the developper that signed the app. Apple returns to you a .cer file, that also contains your public key and is the proof that you Apple “knows” about your public key.

After step 3, you have both keys in Keychain Access and you can now sign apps.

In step 4, the generated .p12 file is simply a file that hold both your public and private keys together in order to share them if needed.

When publishing your app on the AppStore, Apple uses their copy of the public key to verify your app was signed. But it actually goes one bit further : When a user downloads the app in step 6, the downloaded app actually contains the .cer file (you included it during the Archive process) and the iPhone automatically verifies that a) the app is signed and that b) the developer that signed the app is a “trusted developer”.

Pfiou ! We’ve made it !

Well actually, quite not yet. There is this concept of “trusted developer” that we haven’t covered yet. We’ve seen earlier that anyone can create a valid public-private key pair. Yes, the user knows that the developer who signed the app is the one that generated the keys. But since anyone can generate keys, it doesn’t help us much.

Well, there is one last concept we need to understand.

## Identification and Certificate Authorities
As a developer, if you want your key pair to identify you or your organisation, you will need a Certificate Authority to recognise the keys that you generated. There is a widly use system that solves this problem, and it’s called the PKI. PKI stands for Public Key Infrastructure, which is usually described as a set of “management technologies, people and practices” aimed at dealing with digital certification and verification.

In other words, the PKI is designed to bind a public key with an identity (people or organisations). This “binding” is done by organisations called Certificate Authorities. This resulting “binding” is represented by a Certificate file, which holds both the public key and identifying information of the attached identity.

There are many Certificate Authorities, such as private companies (Google, Apple, Microsoft) but also open-source ones, such as OpenSSL. You can chose to trust (or not) each Certificate Authority for doing a good job at making sure their generated key-value pairs are secure and bound to the correct identity.

*Side Note: There are different levels of “binding” with various levels of identification, that can either be performed automatically or by human verification.*

For iOS apps, the Apple Certificate Authority uses both your Mac and your Developer Account to “bind” you (or your organisation) to your generated public key. Apple guarantees that an app that was signed using this key pair corresponds to the right developer. If as a user, you trust Apple to do a good job at this, than you can trust that the Uber app that contains an Apple Developer Certificate was actually signed by Uber.
