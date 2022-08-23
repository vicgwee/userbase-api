package main

import (
	"context"
	"log"
	"net/http"

	"userbase-api/server/dal/mongoDB"
	"userbase-api/server/handlers"

	_ "userbase-api/server/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func handleRequests(ctx context.Context) {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(handlers.LoggingHandler)
	r.Use(handlers.RecoverHandler)
	r.HandleFunc("/users", handlers.GetAllUsersHandler).Methods("GET")
	r.HandleFunc("/users", handlers.CreateNewUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/users/{id}", handlers.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.GetSingleUserHandler).Methods("GET")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":10000", r))
}

// @title Userbase API documentation
// @version 1.0.0
// @host localhost:10000
// @BasePath /
func main() {
	ctx := context.Background()
	if err := initialize(ctx); err != nil {
		log.Fatalf("failed to initialize, err=%v", err)
	}
	defer mongoDB.Disconnect(ctx)
	handleRequests(ctx)
}

func initialize(ctx context.Context) error {
	if err := mongoDB.Setup(ctx); err != nil {
		log.Printf("failed mongoDB setup, err=%v", err)
		return err
	}

	if err := mongoDB.InitUserCollection(ctx); err != nil {
		log.Printf("failed mongoDB InitUserCollection, err=%v", err)
		return err
	}

	return nil

}
