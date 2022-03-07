package main

import (
	"fmt"

	"github.com/ggarcia209/portfolio-k8s/service/util/server"
)

func main() {
	// create http server and handler functions
	fmt.Println("initializing http server...")
	server.InitMuxRouter(8081)
}
