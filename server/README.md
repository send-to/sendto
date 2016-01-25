<img src="https://raw.githubusercontent.com/gophergala2016/sendto/master/images/logo.png">

# Sendto server

Sendto is a quick way for people to send you encrypted files and folders, without knowing anything about encryption, keys or passwords. 

![Sendto](https://raw.githubusercontent.com/gophergala2016/sendto/master/images/sendto.png?s=600)


### Running the server

To run the server, cd to the server directory, and build or run the server.go file:

`go run server.go`

which will launch a local instance. This can then be used to test against, or deployed to a host in order to run your own server instance. You can also use the fragmenta command line tool (https://github.com/fragmenta/fragmenta) to migrate databases, run the server, build cross platform and deploy it.

At present the server database backup is not available, so you'll have to set up the database manually and run migrations to create the database.

### Requirements

* Database: Postgresql

### Try it
https://sendto.click

