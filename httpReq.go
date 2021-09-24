package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = ":8000"
)

func GetRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func PostRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It's a post request!"))
}

func PathVariableHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	w.WriteHeader(404)
	w.Write([]byte("There is no path called " + name + " on this server :)"))
}

func main() {
	router := mux.NewRouter()
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error starting up server: ", err)
	}
	defer logFile.Close()
	router.HandleFunc("/", GetRequestHandler).Methods("GET")
	router.HandleFunc("/post", PostRequestHandler).Methods("POST")
	router.HandleFunc("/hello/{name}", PathVariableHandler).Methods("PUT", "GET")
	http.ListenAndServe(CONN_HOST+CONN_PORT, handlers.LoggingHandler(logFile, router))
}
