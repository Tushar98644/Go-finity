package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

var (
	PORT string
    HOST string
)

func initializer() {
	env := godotenv.Load("../.env"); 
	if env != nil {
		log.Print("No .env file found")
	}

	PORT = os.Getenv("PORT")
	HOST = os.Getenv("HOST")

	if PORT == "" {
		PORT = "8080"
	}

	if HOST == "" {
		HOST = "localhost"
	}

	fmt.Println("Server listening on", HOST+":"+PORT)

}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received:", r.URL.Path)
	fmt.Fprintf(w, "hello guys!")
}

func main() {
	initializer()
	http.HandleFunc("/", hello)

	err := http.ListenAndServe(HOST+":"+PORT, nil)

	if err != nil {
		log.Fatal("error starting http server: ", err)
	}
}
