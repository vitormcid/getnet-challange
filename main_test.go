// main_test.go
package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
         "github.com/vitormcid/getnet-challange"
	"golang.org/x/crypto/bcrypt"
)

var a main.App

func TestMain(m *testing.M) {

	a = main.App{}
	a.Initialize()

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/users", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentUser(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/user/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "User not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'User not found'. Got '%s'", m["error"])
	}
}

func TestCreateUser(t *testing.T) {
	clearTable()

	var jsonStr = []byte(`{"name":"test user", "email": "user@teste.com", "password": "abcd1!432A"}`)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "test user" {
		t.Errorf("Expected user name to be 'test user'. Got '%v'", m["name"])
	}

	if m["email"] != "user@teste.com" {
		t.Errorf("Expected user email to be user@teste.com. Got '%v'", m["email"])
	}

	password := fmt.Sprint(m["password"])

	err := bcrypt.CompareHashAndPassword([]byte(password), []byte("abcd1!432A"))

	if err != nil {
		t.Errorf("Expected user password to be password. Got '%v'", m["password"])
	}

	if m["id"] != 1.0 {
		t.Errorf("Expected user ID to be '1'. Got '%v'", m["id"])
	}
}

func TestInvalidPassword(t *testing.T) {
	clearTable()

	var jsonStr = []byte(`{"name":"test user", "email": "user@teste.com", "password": "invalidPassword"}`)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	message_error := `{"error":"The password is invalid"}`

	if response.Body.String() != message_error {
		t.Errorf("Expected user body response to be '%s'. Got '%s'", response.Body.String(), message_error)
	}
}

func TestGetUser(t *testing.T) {
	clearTable()
	addUsers(1)

	req, _ := http.NewRequest("GET", "/user/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func addUsers(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO users(name, email, password) VALUES($1, $2, $3)", "User "+strconv.Itoa(i), "User@"+strconv.Itoa(i), "abcd1!432A")
	}
}

func TestUpdateUser(t *testing.T) {

	clearTable()
	addUsers(1)

	req, _ := http.NewRequest("GET", "/user/1", nil)
	response := executeRequest(req)
	var originalUser map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalUser)

	var jsonStr = []byte(`{"name":"test user - updated name", "email": "updateuser@teste.com", "password": "abcd1!432B"}`)
	req, _ = http.NewRequest("PUT", "/user/1", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != originalUser["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalUser["id"], m["id"])
	}

	if m["name"] == originalUser["name"] {
		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalUser["name"], m["name"], m["name"])
	}

	if m["email"] == originalUser["email"] {
		t.Errorf("Expected the email to change from '%v' to '%v'. Got '%v'", originalUser["email"], m["email"], m["email"])
	}

	if m["password"] == originalUser["password"] {
		t.Errorf("Expected the password to change from '%v' to '%v'. Got '%v'", originalUser["password"], m["password"], m["password"])
	}
}

func TestDeleteUser(t *testing.T) {
	clearTable()
	addUsers(1)

	req, _ := http.NewRequest("GET", "/user/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/user/1", nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/user/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
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
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM users")
	a.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS users
(
    id SERIAL,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
	password TEXT NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
)`
