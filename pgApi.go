package main

import (
	"log"
	"net/http"

	"fmt"
	"math/rand"
	"time"

	"github.com/gorilla/mux"
)

func generatePassword(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

	// Create a slice of possible characters to use in the password
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:',.<>?")

	// Choose a random length for the password between 8 and 12 characters
	length := rand.Intn(5) + 8

	// Generate the password by choosing random characters from the slice
	password := ""
	for i := 0; i < length; i++ {
		password += string(chars[rand.Intn(len(chars))])
	}

	// Print the password
	fmt.Fprintln(w, "Your new strong password is:", password)
}

func main() {
	router := mux.NewRouter()

	// Generate a strong password
	router.HandleFunc("/generatePassword", generatePassword).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
