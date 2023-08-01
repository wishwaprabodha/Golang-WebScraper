package main

import (
	"github.com/wishwaprabodha/go-webscraper/interfaces/router"
	"log"
	"net/http"
)

func main() {
	log.Println(http.ListenAndServe(":8080", router.StartRoutes()))
	log.Println("Server Started")
}
