//Name: Gunjan Patel
//COEN317 assignment 1

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	ReadTime  = 60
	WriteTime = 90
)

func main() {
	//Starting a new multiplexer/Router to route requests
	mux := http.NewServeMux()

	//Getting current dir for default directory_root (if not specifies)
	wd, _ := os.Getwd()

	//specifying the flags
	dir := flag.String("document_root", wd, "document root")
	port := flag.Int("port", 8080, "specify server listen port")

	//parsing the flags
	flag.Parse()

	//making sure the document_root specified exists, if not then exit with error
	if _, err := os.Stat(*dir); os.IsNotExist(err) {
		fmt.Println("ERROR: File/path does not exist! Exiting the program ...")
		os.Exit(1)
	}

	//ser file root from the flag
	fileRoot := http.Dir(*dir)
	//redirectHandler:= http.RedirectHandler(":8080/index.html", 307)

	//FileServer will return a handler
	files := http.FileServer(fileRoot)
	//http.HandleFunc("/test", testhandler)

	//specify the handler for root url
	mux.Handle("/", files)
	//fmt.Println(*port)
	srvAddr := ":" + strconv.Itoa(*port)
	//fmt.Printf("%T\n\n\n", srvAddr)
	fmt.Println(srvAddr)
	fmt.Println(fileRoot)

	//Set server properties: address:port, handler, Read/Write timeouts (note these will also defer based on HTTP/1.0 and HTTP/1.1, this is maximum)
	srv := &http.Server{
		Addr:         srvAddr,
		Handler:      mux,
		ReadTimeout:  ReadTime * time.Second,
		WriteTimeout: WriteTime * time.Second,
	}

	//Listen on TCP port specified, if not then default port 8080 and localhost
	ln, err := net.Listen("tcp", srvAddr)
	if err != nil {
		log.Fatalln(err.Error())
	}

	//Close the listener when done
	defer ln.Close()

	//Serve the server once server gets a connection and Listener is returned
	srv.Serve(ln) //Serve method spins threads (goroutines) for every new connections

	//srv.ListenAndServe()
}
