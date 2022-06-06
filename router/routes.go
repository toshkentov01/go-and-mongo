package router

import (
	"log"
	"net/http"

	"go-mongo/controllers"

	"github.com/gorilla/mux"
)

func SetUpRoutes(route *mux.Router) {
	s := route.PathPrefix("/api/user").Subrouter() //Base Path
	//Routes
	s.HandleFunc("/create/", controllers.CreateProfile).Methods("POST")
	s.HandleFunc("/get/", controllers.GetUserProfile).Methods("GET")
	s.HandleFunc("/list/", controllers.GetAllUsers).Methods("GET")
	s.HandleFunc("/update/", controllers.UpdateProfile).Methods("PUT")
	s.HandleFunc("/delete/{id}/", controllers.DeleteProfile).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", s)) // Run Server
}
