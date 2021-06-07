package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	//remote packages

	//my packagess

	"go-crud-app/middleware"
	"go-crud-app/router"

	"github.com/stretchr/testify/assert"
)

// func TestGetBooks(t *testing.T) {
// 	req, _ := http.NewRequest("GET", "/books", nil)
// 	response := executeRequest(req)

// 	checkResponseCode(t, http.StatusOK, response.Code)

// 	if body := response.Body.String(); body != "[]" {
// 		t.Errorf("Expected an empty array. Got %s", body)
// 	}
// }

func TestGetBooksRoute(t *testing.T) {
	err := os.Remove("books.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	r := router.Router()
	middleware.ConnectDB()
	middleware.ClearTable()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	assert.Equal(t, w.Body.String(), `{"Books":[]}`)
	middleware.ClearTable()
	os.Remove("books.db")
}

func TestGetBookRoute(t *testing.T) {
	err := os.Remove("books.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	r := router.Router()
	middleware.ConnectDB()
	middleware.ClearTable()

	var testJSON = []byte(`{
		"title": "Example Book",
		"author": "John Adams",
		"publisher": "Harvard Press",
		"publish date": "1997-04-02",
		"rating": 3,
		"checked out": false
	}`)
	w := httptest.NewRecorder()
	req, httpErr := http.NewRequest("POST", "/books", bytes.NewBuffer(testJSON))
	if httpErr != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

}
