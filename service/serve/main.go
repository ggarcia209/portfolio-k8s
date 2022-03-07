package main

import (
	"fmt"
	"log"

	"github.com/ggarcia209/portfolio-k8s/service/util/server"
)

func main() {
	// create http server and handler functions
	fmt.Println("initializing http server...")
	srv := server.InitHTTPServer("localhost:8081")
	fmt.Printf("server address: %v\nread timeout: %v\nwrite timeout: %v\n",
		srv.Addr, srv.ReadTimeout, srv.WriteTimeout)
	fmt.Printf("listening at: '%v'...\n", srv.Addr)
	server.RegisterHandlers()
	log.Fatal(srv.ListenAndServe())
}
