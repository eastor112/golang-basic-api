package handler

import (
	"encoding/json"
	"net/http"
	"simple-api-v1/db"
	"simple-api-v1/models"
	"simple-api-v1/utils"
)

func TestHandler(res http.ResponseWriter, req *http.Request) {
	HandlerMessage := []byte(`{
			"success": true,
			"message": "The server is running properly"
		}`)

	utils.ReturnJsonResponse(res, http.StatusOK, HandlerMessage)
}

func GetMovies(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Check your HTTP method: Invalid HTTP method executed"
		}`)

		utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
		return
	}

	var movies []models.Movie

	for _, movie := range db.Moviedb {
		movies = append(movies, movie)
	}

	movieJSON, err := json.Marshal(&movies)

	if err != nil {
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Error parsing the movie data",
		}`)

		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
		return
	}

	utils.ReturnJsonResponse(res, http.StatusOK, movieJSON)
}

func GetMovie(res http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Check your HTTP method: Invalid HTTP method executed",
		}`)

		utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
		return
	}

	if _, ok := req.URL.Query()["id"]; !ok {
		HandlerMessage := []byte(`{
			"success": false,
			"message": "This method requires the movie id",
		}`)

		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
		return
	}

	id := req.URL.Query()["id"][0]

	movie, ok := db.Moviedb[id]
	if !ok {
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Requested movie not found",
		}`)

		utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
		return
	}

	movieJSON, err := json.Marshal(&movie)
	if err != nil {
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Error parsing the movie data",
		}`)

		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
		return
	}

	utils.ReturnJsonResponse(res, http.StatusOK, movieJSON)
}

func AddMovie(res http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Check your HTTP method: Invalid HTTP method executed",
		}`)

		utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
		return
	}

	var movie models.Movie

	payload := req.Body

	defer req.Body.Close()

	err := json.NewDecoder(payload).Decode(&movie)
	if err != nil {
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Error parsing the movie data",
		}`)

		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
		return
	}

	db.Moviedb[movie.ID] = movie

	HandlerMessage := []byte(`{
		"success": true,
		"message": "Movie was successfully created",
	}`)

	utils.ReturnJsonResponse(res, http.StatusCreated, HandlerMessage)
}

func DeleteMovie(res http.ResponseWriter, req *http.Request) {

	if req.Method != "DELETE" {
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Check your HTTP method: Invalid HTTP method executed",
		}`)

		utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
		return
	}

	if _, ok := req.URL.Query()["id"]; !ok {
		HandlerMessage := []byte(`{
			"success": false,
			"message": "This method requires the movie id",
		}`)

		utils.ReturnJsonResponse(res, http.StatusBadRequest, HandlerMessage)
		return
	}

	id := req.URL.Query()["id"][0]
	_, ok := db.Moviedb[id]
	if !ok {
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Requested movie not found",
		}`)

		utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
		return
	}

	delete(db.Moviedb, id)

	HandlerMessage := []byte(`{
			"success": true,
			"message": "Movie deleted successfully"
	}`)
	utils.ReturnJsonResponse(res, http.StatusOK, HandlerMessage)
}
