package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type SqlObj struct {
	Username string `json:"username"`
	Port     string `json:"port"`
}

func main() {
	// os.Setenv("FOO", "1")
	// fmt.Println("FOO:", os.Getenv("FOO"))
	// fmt.Println("BAR:", os.Getenv("BAR"))
	// Use os.Environ to list all key/value pairs in the environment. This returns a slice of strings in the form KEY=value. You can strings.SplitN them to get the key and value. Here we print all the keys.

	fmt.Println()
	e := os.Environ()
	fmt.Println("Reaching Here eavlkvnlse", e)
	for _, e := range os.Environ() {
		fmt.Println("Reaching Here eavlkvnlse", e)
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", handler).Methods("GET")
	router.HandleFunc("/evariables", eventHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func handler(w http.ResponseWriter, r *http.Request) {
	responseJSON(w, http.StatusOK, "listening")
}

func eventHandler(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("MY_USER")
	if username == "" {
		log.Fatal("Environment variable not set")
	}
	port := os.Getenv("MY_PORT")
	if port == "" {
		log.Fatal("Environment variable not set")
	}
	sql := SqlObj{
		Username: username,
		Port:     port,
	}
	responseJSON(w, http.StatusOK, sql)
}
