package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // <------------ here
	"github.com/renanmachad/ecommerce-rest/pkg/app/rest"
	"github.com/renanmachad/ecommerce-rest/pkg/di"
	"github.com/renanmachad/ecommerce-rest/pkg/routes"
	"github.com/renanmachad/ecommerce-rest/pkg/utils"
	"log"
	"net/http"
)

func main() {

	// init mux
	router := mux.NewRouter()
	// init db connection
	db, err := di.NewConfig()
	// handle error
	utils.ErrorHandler(err)

	// make default port to test if server is okay
	router.HandleFunc("/", rest.Home)
	
	routes.OrderRoute(router, db)

	log.Println("Starting server on port 8006")
	log.Fatal(http.ListenAndServe(":8006", router))
}
