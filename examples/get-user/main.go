package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

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

	var userPretty bytes.Buffer
	if err = json.Indent(&userPretty, user, "", "\t"); err != nil {
		log.Fatal("JSON parse error: ", err)
	}

	fmt.Println(userPretty.String())
}
