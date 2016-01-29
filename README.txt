Name: Gunjan Patel
Assignment 1: web server (written in Go)
Date: Jan 28






Test cases: 
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


List of files:

├── Santa Clara University_files
│   ├── -livewhale-plugins-jquery-jquery.lw-widget.js
│   ├── 311_levis.jpg
│   ├── BryanStevenson760480-760x481.jpg
...  ... ...
│   ├── scu.css
│   ├── search-autocomplete.js
│   └── sinskywchild760480.jpg
├── nopermission.html			//set it to no read permission (sudo chmod 200 nopermission.html)
├── index.html
├── minion.gif
├── r2d2.jpg
├── server
└── server.go
script.PDF

1 directory, 42 files



Instructions for running:
-This assignment is written in Google's new programming language Go aka Golang (with professor's prior permission)
-Tested with Go version 1.5
-makefile is provided 
-If you want to test it and don't have Go installed, use this link where I have this project setup with Go installed in the cloud: https://ide.c9.io/gunjan5/coen317-assignment1