package main

// ch <- 42 // send
// variable := <-ch // receive

// * Unbuffered → “sender and receiver must meet.”
// * Buffered → “sender can drop results into the channel and continue.”

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FetchUser(id int, ch chan User, wg *sync.WaitGroup) {
	defer wg.Done()

	url := fmt.Sprintf(
		"https://jsonplaceholder.typicode.com/users/%d",
		id,
	)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	defer resp.Body.Close()

	var user User

	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	ch <- user
}

func main() {
	var wg sync.WaitGroup

	ch := make(chan User, 2)

	wg.Add(2)

	go FetchUser(1, ch, &wg)
	go FetchUser(2, ch, &wg)

	// user1 := <-ch
	// user2 := <-ch

	// fmt.Println(user1)
	// fmt.Println(user2)

	wg.Wait()
	close(ch)

	// close the channel when using range
	// don’t bother closing if you know exactly how many values you’ll receive
	for user := range ch {
		fmt.Println(user)
	}

}
