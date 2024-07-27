package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForLogin(t *testing.T) {
	// Setup a mock server using httptest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/auth/login", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		// Simulating a response
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"token": "mockToken"})
	}))
	defer server.Close()

	// Call the login function
	token := loginFunc(t)

	// Assert that the token is as expected
	assert.NotEmpty(t, token, "The token should not be empty")
}

func TestUpdateAuthor(t *testing.T) {
	token := loginFunc(t)
	body := `{"name": "Updated Test Author"}`
	req, err := http.NewRequest("PUT", "http://localhost:8080/authors/39", bytes.NewBuffer([]byte(body)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDeleteAuthor(t *testing.T) {
	token := loginFunc(t)
	req, err := http.NewRequest("DELETE", "http://localhost:8080/authors/40", nil)
	assert.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCreateAuthor(t *testing.T) {
	token := loginFunc(t)
	body := `{"name": "Test Author", "birth": "2000-01-01"}`
	req, err := http.NewRequest("POST", "http://localhost:8080/authors", bytes.NewBuffer([]byte(body)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestFindAuthorByID(t *testing.T) {
	token := loginFunc(t)
	req, err := http.NewRequest("GET", "http://localhost:8080/authors/39", nil)
	assert.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestFindAllAuthors(t *testing.T) {
	token := loginFunc(t)
	req, err := http.NewRequest("GET", "http://localhost:8080/authors", nil)
	assert.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Read the response body
	body, err := json.Marshal(resp.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)
}

func loginFunc(t *testing.T) string {
	loginBody := `{"username": "username", "password": "passw0rd#"}`
	req, err := http.NewRequest("POST", "http://localhost:8080/auth/login", bytes.NewBuffer([]byte(loginBody)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response map[string]string
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)

	token, ok := response["token"]
	assert.True(t, ok, "Token not found in response")

	return token
}
