package main

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/ichtrojan/fuckthepolice/controller"
	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	logger, _ := thoth.Init("log")

	if err := godotenv.Load(); err != nil {
		logger.Log(errors.New("no .env file found"))
		log.Fatal("No .env file found")
	}

	route := mux.NewRouter()

	route.HandleFunc("/incidents/{id}", controller.GetIncident).Methods("GET")
	route.HandleFunc("/incidents", controller.AllIncidents).Methods("GET")
	route.HandleFunc("/incidents", controller.ReportIncident).Methods("POST")

	port, exist := os.LookupEnv("PORT")

	if !exist {
		logger.Log(errors.New("PORT not set in .env"))
		log.Fatal("PORT not set in .env")
	}

	if err := http.ListenAndServe(":"+port, route); err != nil {
		logger.Log(err)
		log.Fatal(err)
	}
}
