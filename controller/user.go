package controller

import (
	"fmt"
	"log"
	"net/http"
	"sample/service"

	"sample/docs"
	"sample/exception"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func HandleRequests() {
	router := mux.NewRouter()
	fmt.Println(docs.SwaggerInfo)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.HandleFunc("/users", service.GetUsers).Methods("GET")
	router.HandleFunc("/users", service.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", service.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", service.DeleteUser).Methods("DELETE")

	router.Use(exception.ExceptionAdvisorMiddleware)
	log.Fatal(http.ListenAndServe(":8080", router))

}
