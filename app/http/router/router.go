package router

import (
	"github.com/gorilla/mux"
	"hello-world-api/app/http/controllers"
	"net/http"
)

func Init() *mux.Router {

	r := mux.NewRouter()

	nameController := controllers.NewNameController()

	r.HandleFunc("/hello-world", nameController.GetName).Methods(http.MethodGet)

	return r
}
