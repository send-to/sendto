# Sendto

Sendto is a quick way for people to send you encrypted files and folders, without knowing anything about encryption, keys or passwords. 

* No key generation or passwords. Sendto lets users use your PGP public key to encrypt files for you, and uploads them to the server encrypted for you to download. There is no difficult dance of sending keys or passwords on an insecure channel, or complex software for them to master. 
* Files encrypted at all times. Sendto cannot open your files because it only knows about your public key, so it can encrypt but never decrypt. TLS is also used for all connections. 
* Open Source. Sendto is completely open source, so that you can verify what happens to your files, and run it on your own server if you prefer self-hosting. 

### Receive files securely

Just send people a link to your profile, and they can download an app for their platform to send you encrypted files. After that download, on Mac OS X they can send you code just by right clicking a file or folder and choosing Services > Send to YOURNAME, at which point you can see it on the website. Other platforms at present have a command line app, but will have drag and drop. 

Try out sending a file to my profile: sendto.click/kennygrant, or set up your own.

### Team
@kennygrant on twitter, github, keybase.io

### Tech
Go 1.4
PGP Encryption : Go Std library
Web Framework: Fragmenta
Web Server: Caddy
Database: Postgresql

### Try it
https://sendto.click


### Open source on github
https://github.com/gophergala/sendto


This app is open source so that you can build it yourself, and check what it does. If you have Go installed, you can also install the client and server from source with:

`go get github.com/gophergala2016/sendto`

and then use the sendto command:

`sendto help` 

you can host the server yourself if you prefer to have complete control. 