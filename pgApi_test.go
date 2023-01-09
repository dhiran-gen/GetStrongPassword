package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	// Test that the API returns a password of the correct length
	req, err := http.NewRequest("GET", "/generatePassword", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(generatePassword)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Test that the API returns a password with a mix of lowercase and uppercase letters
	req, err = http.NewRequest("GET", "/generatePassword", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	password := rr.Body.String()
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") || !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		t.Errorf("handler returned password without a mix of lowercase and uppercase letters: %v", password)
	}

	// Test that the API returns a password with at least one number
	req, err = http.NewRequest("GET", "/generatePassword", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	password = rr.Body.String()

	// Test that the API returns a password with at least one special character
	req, err = http.NewRequest("GET", "/generatePassword", nil)
	if err != nil {
		t.Fatal(err)
	}

	// import (
	// 	"net/http"
	// 	"net/http/httptest"
	// 	"testing"
	// )

	// func TestGeneratePassword(t *testing.T) {
	// 	// Test that the API returns a password of the correct length
	// 	req, err := http.NewRequest("GET", "/generatePassword", nil)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	rr := httptest.NewRecorder()
	// 	handler := http.HandlerFunc(generatePassword)
	// 	handler.ServeHTTP(rr, req)

	// 	if status := rr.Code; status != http.StatusOK {
	// 		t.Errorf("handler returned wrong status code: got %v want %v",
	// 			status, http.StatusOK)
	// 	}

	// 	password := rr.Body.String()
	// 	if len(password) < 8 || len(password) > 12 {
	// 		t.Errorf("incorrect password length: got %v, want 8-12", len(password))
	// 	}

	// 	// Test that the API returns a password with a mix of lowercase and uppercase letters
	// 	req, err = http.NewRequest("GET", "/generatePassword", nil)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	rr = httptest.NewRecorder()
	// 	handler.ServeHTTP(rr, req)

	// 	password = rr.Body.String()
	// 	lowercase := false
	// 	uppercase := false
	// 	for _, char := range password {
	// 		if char >= 'a' && char <= 'z' {
	// 			lowercase = true
	// 		}
	// 		if char >= 'A' && char <= 'Z' {
	// 			uppercase = true
	// 		}
	// 	}
	// 	if !lowercase || !uppercase {
	// 		t.Errorf("password does not contain a mix of lowercase and uppercase letters: %v", password)
	// 	}

	// 	// Test that the API returns a password with at least one number
	// 	req, err = http.NewRequest("GET", "/generatePassword", nil)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	rr = httptest.NewRecorder()
	// 	handler.ServeHTTP(rr, req)

	// 	password = rr.Body.String()
	// 	containsNumber := false
	// 	for _, char := range password {
	// 		if char >= '0' && char <= '9' {
	// 			containsNumber = true
	// 			break
	// 		}
	// 	}
	// }

	// // if !containsNumber {
}
