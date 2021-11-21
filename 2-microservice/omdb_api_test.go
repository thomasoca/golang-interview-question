package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchPagination(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "A")
	}))

	defer server.Close()
	r, err := searchPagination(server.URL, "batman", "test123", 2)
	if err != nil {
		t.Errorf("Expected nil, received %v", err)
	}
	fmt.Println(r)

}
