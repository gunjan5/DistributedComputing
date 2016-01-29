build:
	go build server.go

run:  #runs the program with default options that is current dir as the doc_root and port 8080
	./server

test:
	telnet localhost 8080
	GET / HTTP/1.0
	$(\r\n)

