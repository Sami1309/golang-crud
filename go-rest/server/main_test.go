package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	//remote packages

	//my packagess

	"go-crud-app/middleware"
	"go-crud-app/models"
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
	//err := os.Remove("books.db")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
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
		panic(httpErr)
	}
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Printf("%T", w.Body)
	fmt.Println(w.Body.String())

	type CreateBook struct {
		CreatedBook models.Book `json:"Created book"`
	}

	var book CreateBook

	json.Unmarshal([]byte(w.Body.String()), &book)

	// fmt.Println(json.NewDecoder(w.Body).Decode(target))

	fmt.Println(book)

	fmt.Println(book.CreatedBook.Id)

	createdBookID := book.CreatedBook.Id

	getBookRequest := fmt.Sprintf("/books/%d", createdBookID)

	// 200 request
	bookW := httptest.NewRecorder()
	book_req, _ := http.NewRequest("GET", getBookRequest, nil)
	r.ServeHTTP(bookW, book_req)

	assert.Equal(t, 200, bookW.Code)

	getBookRequest = fmt.Sprintf("/books/%d", createdBookID+1)
	//400 request
	bookBadW := httptest.NewRecorder()
	book_bad_req, _ := http.NewRequest("GET", getBookRequest, nil)
	r.ServeHTTP(bookBadW, book_bad_req)

	assert.Equal(t, 404, bookBadW.Code)

}

//Test Get 200, 400

//Test Create 200, 400
//use eturrned id to subsequently get and check for a values
//make sure all internal values are the same (check struct structure)
func TestCreateBookRoute(t *testing.T) {

	//200 test
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
		panic(httpErr)
	}
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Printf("%T", w.Body)
	fmt.Println(w.Body.String())

	type CreateBook struct {
		CreatedBook models.Book `json:"Created book"`
	}

	var book CreateBook

	json.Unmarshal([]byte(w.Body.String()), &book)

	// fmt.Println(json.NewDecoder(w.Body).Decode(target))

	fmt.Println(book)

	fmt.Println(book.CreatedBook.Id)

	//400 test
	var badJSON_1 = []byte(`{
		"title": "Example Book",
		"author": "John Adams",
		"publisher": "Harvard Press",
		"publish date": "1997-04-02",
		"rating": 5,
		"checked out": false
	}`)
	var badJSON_2 = []byte(`{
		"title": "Example Book",
		"author": "John Adams",
		"publisher": "Harvard Press",
		"publish date": "1997-04-02",
		"checked out": false
	}`)
	var badJSON_3 = []byte(`{
		"title": "Example Book",
		"author": "John Adams",
		"publisher": "Harvard Press",
		"publish date": "1997-uu",
		"rating": 2,
		"checked out": false
	}`)

	w = httptest.NewRecorder()
	req, httpErr = http.NewRequest("POST", "/books", bytes.NewBuffer(badJSON_1))
	if httpErr != nil {
		panic(httpErr)
	}
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	fmt.Printf("%T", w.Body)
	fmt.Println(w.Body.String())
	w = httptest.NewRecorder()
	req, httpErr = http.NewRequest("POST", "/books", bytes.NewBuffer(badJSON_2))
	if httpErr != nil {
		panic(httpErr)
	}
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

	w = httptest.NewRecorder()
	req, httpErr = http.NewRequest("POST", "/books", bytes.NewBuffer(badJSON_3))
	if httpErr != nil {
		panic(httpErr)
	}
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	//create function that tests for any inputted json byte array
}

//Test Update 200, 400
//create, get id, check vaues, update, check if changed, etc

func TestUpdateBookRoute(t *testing.T) {

	r := router.Router()
	middleware.ConnectDB()

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
		panic(httpErr)
	}
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Printf("%T", w.Body)
	fmt.Println(w.Body.String())

	type CreateBook struct {
		CreatedBook models.Book `json:"Created book"`
	}

	var book CreateBook

	json.Unmarshal([]byte(w.Body.String()), &book)

	// fmt.Println(json.NewDecoder(w.Body).Decode(target))

	fmt.Println(book)

	fmt.Println(book.CreatedBook.Id)

	createdBookID := book.CreatedBook.Id

	getBookRequest := fmt.Sprintf("/books/%d", createdBookID)

	var updateJSON = []byte(`{
		"title": "Example Book",
		"author": "Alexander Hamilton",
		"checked out": false
	}`)
	// 200 request
	bookW := httptest.NewRecorder()
	book_req, _ := http.NewRequest("PATCH", getBookRequest, bytes.NewBuffer(updateJSON))
	r.ServeHTTP(bookW, book_req)

	assert.Equal(t, 200, bookW.Code)
	//TODO check that update worked

	getBookRequest = fmt.Sprintf("/books/%d", createdBookID+1)

	bookW = httptest.NewRecorder()
	book_req, _ = http.NewRequest("PATCH", getBookRequest, bytes.NewBuffer(updateJSON))
	r.ServeHTTP(bookW, book_req)

	assert.Equal(t, 400, bookW.Code)

	var badJSON = []byte(`{
		"title": "Example Book",
		"author": "'Lex Hamilton",
		"rating": 5
		"checked out": false
	}`)

	getBookRequest = fmt.Sprintf("/books/%d", createdBookID)

	bookW = httptest.NewRecorder()
	book_req, _ = http.NewRequest("PATCH", getBookRequest, bytes.NewBuffer(badJSON))
	r.ServeHTTP(bookW, book_req)

	assert.Equal(t, 400, bookW.Code)

}

func TestDeleteBookRoute(t *testing.T) {
	r := router.Router()
	middleware.ConnectDB()

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
		panic(httpErr)
	}
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Printf("%T", w.Body)
	fmt.Println(w.Body.String())

	type CreateBook struct {
		CreatedBook models.Book `json:"Created book"`
	}

	var book CreateBook

	json.Unmarshal([]byte(w.Body.String()), &book)

	// fmt.Println(json.NewDecoder(w.Body).Decode(target))

	// fmt.Println(book)

	// fmt.Println(book.CreatedBook.Id)

	createdBookID := book.CreatedBook.Id

	getBookRequest := fmt.Sprintf("/books/%d", createdBookID)

	// 200 request
	bookW := httptest.NewRecorder()
	book_req, _ := http.NewRequest("DELETE", getBookRequest, nil)
	r.ServeHTTP(bookW, book_req)

	assert.Equal(t, 200, bookW.Code)

	getBookRequest = fmt.Sprintf("/books/%d", createdBookID+1)

	//400 request
	bookW = httptest.NewRecorder()
	book_req, _ = http.NewRequest("DELETE", getBookRequest, nil)
	r.ServeHTTP(bookW, book_req)

	assert.Equal(t, 404, bookW.Code)
}

//Test Delete
//Crreate and delete and make sure looking up id gets 404
