<h1><img src="https://raw.githubusercontent.com/gophergala2016/sendto/master/images/logo-sml.png" height=32 width=32> Sendto</h1>

Sendto is a quick way for people to send you encrypted files and folders, without knowing anything about encryption, keys or passwords. 

* No key generation or passwords. Sendto lets users use your PGP public key to encrypt files for you, and uploads them to the server encrypted for you to download. There is no difficult dance of sending keys or passwords on an insecure channel, or complex software for them to master. 
* Files encrypted at all times. Sendto cannot open your files because it only knows about your public key, so it can encrypt but never decrypt. TLS is also used for all connections. 
* Open Source. Sendto is completely open source, so that you can verify what happens to your files, and run it on your own server if you prefer self-hosting. 

![Sendto](https://raw.githubusercontent.com/gophergala2016/sendto/master/images/sendto.png?s=600)


### Receive files securely

Just send people a link to your profile, and they can download an app for their platform to send you encrypted files. An example profile is here:

https://sendto.click/users/demo

This shows the public key, and a set of download links. Download the binary for your platform (or install from source if you prefer, see below), and then you can send a file to any sendto account for which you know the username, with a command like:

`sendto demo my/file/or/folder`

this will send the file to the demo account. You won't be able to download it though, as you need your own profile to access uploaded files. Please only send test files to the demo account. 

Once you've tried the demo, if you have a pgp key, or a keybase.io account, try setting up a user. The server can pull keys automatically from keybase.io so setup is easy. No email is required for sign up at this time. 

Once files are uploaded to your account, you're able to view and download them by logging in. Decryption happens on your machine, so that your private keys are never shared with the server. 

![Sendto](https://raw.githubusercontent.com/gophergala2016/sendto/master/images/files.png?s=600)

### Team
@kennygrant on twitter, <a href="https://github.com/kennygrant">github</a>, <a href="https://keybase.io/kennygrant">keybase.io</a>, <a href="https://sendto.click/users/kennygrant">sendto.click</a>

### Tech
* Go 1.4
* PGP Encryption : golang.org/x/crypto/openpgp
* Web Framework: Fragmenta
* Web Server: Caddy
* Database: Postgresql

### Try it
https://sendto.click


### Open source on github
go get github.com/send-to

This app is open source so that you can build it yourself, and check what it does. If you have Go installed, you can also install the client and server from source with:

`go get github.com/send-to/sendto`

and then use the sendto command:

`sendto help` 

you can host the server yourself if you prefer to have complete control. 


### Possible bugs

I haven't had much time to test on Windows, so there may be path issues there. Tested on Mac OS X and linux. The server runs on linux. 

Keys are cached locally for users (at ~/.sendto), so if you change your key you'd have to remove the prefs locally in order to update it. This needs fixed at some point obviously. 
