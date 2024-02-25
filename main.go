package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: password <length> <has special characters: true|false>")
		return
	}

	length, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid length. Please provide an integer.")
		return
	}

	hasSpecial, err := strconv.ParseBool(os.Args[2])
	if err != nil {
		fmt.Println("Invalid value for special characters. Use true or false.")
		return
	}

	fmt.Println(generatePassword(length, hasSpecial))
}

func generatePassword(length int, hasSpecial bool) string {
	var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	if hasSpecial {
		letters += "!@#$%^&*()"
	}

	// Creating a new local random generator, seeded with the current time
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rng.Intn(len(letters))] // Use the local rng
	}
	return string(b)
}
