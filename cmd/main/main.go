package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/korn/learning/book_store/pkg/routes"
)

func main() {
	router := gin.Default()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
