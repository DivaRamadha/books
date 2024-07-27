package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {

	body := `{"username": "TestUserss", "password": "QWErty123!"}`
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/auth/login", bytes.NewBuffer([]byte(body)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var response map[string]string
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, http.StatusOK, resp.StatusCode)

	token, ok := response["token"]
	require.True(t, ok, "Token not found in response")

	require.NotEmpty(t, token)
}

func RegisterTestUser(t *testing.T) {
	body := `{"username": "TestUserss", "password": "QWErty123!"}`
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/auth/register", bytes.NewBuffer([]byte(body)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got %d", http.StatusOK, resp.StatusCode)
	}

	require.Equal(t, http.StatusOK, resp.StatusCode)
}
