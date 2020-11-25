package main

import (
	"github.com/zerosuxx/go-http-server/go-http-server/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	port := "1234"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	log.Print("Server listening on: http://localhost:" + port)

	http.HandleFunc("/healthcheck", handler.CreateHealthCheckHandler().Handle)
	http.HandleFunc("/cmd", handler.CreateCommandHandler().Handle)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
