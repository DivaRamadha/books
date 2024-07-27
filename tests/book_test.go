package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

var createdBookID uint

// login function to get the toke
// createBook function to create a new book
func createBook(t *testing.T, token string) uint {
	body := `{"title": "Test Book", "isbn": "1234567890", "author_id": 39}`
	req, err := http.NewRequest("POST", "http://localhost:8080/books", bytes.NewBuffer([]byte(body)))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)

	var createdBook map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&createdBook)
	require.NoError(t, err)

	bookID, ok := createdBook["ID"].(float64)
	require.True(t, ok, "ID not found in response")

	return uint(bookID)
}

// TestCreateBook tests creating a book and then performs delete
func TestBookAPI(t *testing.T) {
	token := loginFunc(t)
	createdBookID = createBook(t, token) // Store the created book ID for use in other tests
	TestFindAllBooks(t)
	TestFindBookByID(t)
	TestUpdateBook(t)
	TestDeleteBook(t)
}

// TestUpdateBook updates a book and verifies the response
func TestUpdateBook(t *testing.T) {
	token := loginFunc(t)
	body := `{"title": "Updated Test Book", "isbn": "0987654321", "author_id": 39}`
	req, err := http.NewRequest("PUT", fmt.Sprintf("http://localhost:8080/books/%d", createdBookID), bytes.NewBuffer([]byte(body)))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)
}

// TestFindAllBooks finds all books and verifies the response
func TestFindAllBooks(t *testing.T) {
	token := loginFunc(t)
	req, err := http.NewRequest("GET", "http://localhost:8080/books", nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)
}

// TestFindBookByID finds a book by ID and verifies the response
func TestFindBookByID(t *testing.T) {
	token := loginFunc(t)
	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/books/%d", createdBookID), nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

// TestDeleteBook deletes a book and verifies the response
func TestDeleteBook(t *testing.T) {
	token := loginFunc(t)
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8080/books/%d", createdBookID), nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)
}
