package main

import (
	"github.com/gorilla/mux"

	"go-mongo/router"
)


func main() {
	route := mux.NewRouter()
	router.SetUpRoutes(route)
}
