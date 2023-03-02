package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"simple-api-v1/db"
	"simple-api-v1/handler"
	"simple-api-v1/models"
)

func main() {
	db.Moviedb["001"] = models.Movie{ID: "001", Title: "A Space Odyssey", Description: "Science fiction"}
	db.Moviedb["002"] = models.Movie{ID: "002", Title: "Citizen Kane", Description: "Drama"}
	db.Moviedb["003"] = models.Movie{ID: "003", Title: "Raiders of the Lost Ark", Description: "Action and adventure"}
	db.Moviedb["004"] = models.Movie{ID: "004", Title: "66. The General", Description: "Comedy"}

	http.HandleFunc("/", handler.TestHandler)
	http.HandleFunc("/movies", handler.GetMovies)
	http.HandleFunc("/movie", handler.GetMovie)
	http.HandleFunc("/movie/add", handler.AddMovie)
	http.HandleFunc("/movie/delete", handler.DeleteMovie)

	log.Print("The is Server Running on localhost port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
