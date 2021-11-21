package main

import (
	"log"
	"net/http"

	"strconv"
)

const url = "http://www.omdbapi.com/"

func searchMovie(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	api_key := query.Get("apikey")
	searchWord := query.Get("s")
	pagination := query.Get("page")
	paginationInt, err := strconv.Atoi(pagination)
	if err != nil {
		log.Panic(err)
	}
	resp, err := searchPagination(url, searchWord, api_key, paginationInt)
	go logDb(searchWord, paginationInt)
	if err != nil {
		log.Panic(err)
	}
	_, err = w.Write(resp)
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(`{"message": "Failed fetching data"}`))
		if err != nil {
			log.Panic(err)
		}
	}
}

func getMovieDetail(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	api_key := query.Get("apikey")
	movieTitle := query.Get("t")
	resp, err := movieDetail(url, movieTitle, api_key)
	if err != nil {
		log.Panic(err)
	}
	_, err = w.Write(resp)

	if err != nil {
		w.Header().Set("Content-type", "application/json")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(`{"message": "Failed fetching data"}`))
		if err != nil {
			log.Panic(err)
		}
	}
}

func main() {

	http.HandleFunc("/search", searchMovie)
	http.HandleFunc("/detail", getMovieDetail)

	http.ListenAndServe(":8090", nil)
}
