package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/antonprokopovich/go-threads"
)

func main() {
	t, err := threads.NewThreads()
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	replies, err := t.GetUserReplies(314216)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	var repliesPretty bytes.Buffer
	if err = json.Indent(&repliesPretty, replies, "", "\t"); err != nil {
		log.Fatal("JSON parse error: ", err)
	}

	fmt.Println(repliesPretty.String())
}
