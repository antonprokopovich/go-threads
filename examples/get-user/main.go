package main

import (
	"encoding/json"
	"fmt"

	"github.com/antonprokopovich/threadsnet"
)

func main() {
	t, err := threads.NewThreads()
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	userID, err := t.GetUserID("zuck")
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	user, err := t.GetUser(userID)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	userBytes, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)

		return
	}
	fmt.Println(string(userBytes))
}
