package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	main "github.com/theovassiliou/pwrulevalidator"
)

var a main.App

func TestMain(m *testing.M) {
	a.Initialize()
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/validators", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Exppected an empty array. Got %s", body)
	}
}

func TestGetNonExistentValidator(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/validator/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Validator not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Validator not found'. Got '%s'", m["error"])
	}
}
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
func ensureTableExists() {
}

func clearTable() {
}
