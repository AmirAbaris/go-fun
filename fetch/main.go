package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/1")

	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		fmt.Println("Unexpected status:", resp.Status)
		return
	}

	var user User

	err = json.NewDecoder(resp.Body).Decode(&user)

	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Name:", user.Name)
	fmt.Println("Email:", user.Email)
}
