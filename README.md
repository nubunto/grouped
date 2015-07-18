# Grouped

This is a web app, that streams a video using a webpage.
The intent of this app is to be served using [ngrok](http://ngrok.com) from a computer acting as the server.
The server is the user with the video file, located in the `static` folder.

# What else?

The web app also features a (rather primitive) chat application. Other than that, it provides a "pause/unpause" functionality, so that all parties watching
the video can pause it or unpause it.

# What can I do with it?

As is, you can download it (although it is not the most user friendly thing in the entire world) and install some additional tools to make it work, and watch
a simple video with your friends, over the internet, with no hosting whatsoever.
If you are really interested, you can hack it up and improve it (and please, if you do, feel free to send me a pull request).

# Instalation

You need an external tool, called *ngrok* (you can download it [here](http://ngrok.com)). It exposes through a secure tunnel a web server running on your localhost.
If you download ngrok and run it without signin up, you won't be able to use a custom subdomain, which means you will have to type a weird URL to see your app on the web.
If that's not your thing, sign up in the ngrok website, and follow the instructions there to accomodate that.

After that, all you need to do is:
* Install [Go](http://golang.org)
* Using the command line, `cd` in the exported directory and compile the app with `go build`.
* A new file will pop up. If you're on linux, run `./file-that-just-popped-up`.
* The app will be listening on port `8080`. Leave it running in this command line tab, open another one, and run **ngrok** with the port the app is listening in: `ngrok http 8080`.
* The app serves the file called `video.webm` inside the `static` folder. Make sure the video is in the correct MIME type supported by web browsers. A handy tool here is [Cloud Convert](http://cloudconvert.com).
* Go have fun with your friends, as they join in and watch a video together.

I hope you enjoy this simple project. Go is a awesome platform, and I intend to continue learning it.

## Some side notes

The code in this projects is a proof of concept. It is by no means intended to actual production use. In other words, I am aware it sucks. It forces the user to configure stuff by hand, changing the video file, etc. That may (or may not) be addressed, so feel free to use this only as a learning tool. 

Also, this code here is actually borrowed from [this article](http://gary.burd.info/go-websocket-chat). I only studied it and added the pause/unpause functionality. Thanks Gary, you're awesome. And sorry for not forking it. I only now realize your code is on github.
