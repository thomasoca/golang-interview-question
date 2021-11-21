package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func apiCall(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)

		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return body, nil
}

func searchPagination(url string, searchWord string, api_key string, pagination int) ([]byte, error) {

	url_req := fmt.Sprintf("%s/?apikey=%s&s=%s&page=%d", url, api_key, searchWord, pagination)
	resp, err := apiCall(url_req)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return resp, nil
}

func movieDetail(url string, title string, api_key string) ([]byte, error) {
	url_req := fmt.Sprintf("%s/?apikey=%s&i=%s", url, api_key, title)
	resp, err := apiCall(url_req)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return resp, nil
}
