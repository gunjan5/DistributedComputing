Name: Gunjan Patel
Assignment 1: web server (written in Go)
Date: Jan 28


* Test cases [All passing]: 
================
HTTP/1.0, 200 OK
	telnet localhost 8080
	GET / HTTP/1.0

HTTP/1.0, 400 BAD REQUEST
	telnet localhost 8080
	GET bad HTTP/1.0

HTTP/1.0, 404 NOT FOUND
	telnet localhost 8080
	GET /nothing.html HTTP/1.0

HTTP/1.0, 403 FORBIDDEN
	telnet localhost 8080
	GET /nopermission.html HTTP/1.0	

================
HTTP/1.1, 200 OK
	telnet localhost 8080
	GET / HTTP/1.0

HTTP/1.1, 400 BAD REQUEST

	GET bad HTTP/1.0

HTTP/1.1, 404 NOT FOUND

	GET /nothing.html HTTP/1.0

HTTP/1.1, 403 FORBIDDEN

	GET /nopermission.html HTTP/1.0	

=================
Supports HTTP/1.1 and 1.0
Spins new threads (goroutine) for every new connections 



* List of files:
================

.
├── Makefile          //makefile
├── README.txt        //this file :)
├── Santa Clara University_files //default page supporting files directory
│   ├── -livewhale-plugins-jquery-jquery.lw-widget.js
│   ├── 311_levis.jpg
... ... ...
│   ├── scu.css
│   ├── search-autocomplete.js
│   └── sinskywchild760480.jpg
├── index.html         //default page
├── minion.gif         //test GIF file
├── nopermission.html  //set it to no read permission (sudo chmod 200 nopermission.html)
├── r2d2.jpg           //test JPG file
├── script.pdf        //screenshots of some tests
├── server             //compiled program (run `make build` to get it)
└── server.go         //main code

1 directory, 45 files




* Instructions for running:
=========================
-This assignment is written in Google's new programming language Go aka Golang (with professor's prior permission)
-Tested with Go version 1.5
-makefile is provided 
-If you want to test it and don't have Go installed, use this link where I have this project setup with Go installed in the cloud: https://ide.c9.io/gunjan5/coen317-assignment1