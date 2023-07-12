package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	t, err := threads.NewThreads()
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	post, err := t.GetUserThreads(314216)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	var repliesPretty bytes.Buffer
	if err = json.Indent(&repliesPretty, post, "", "\t"); err != nil {
		log.Fatal("JSON parse error: ", err)
	}

	fmt.Println(repliesPretty.String())
}
