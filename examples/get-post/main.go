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

	post, err := t.GetPost(3141055616164096839)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	postJson, err := json.MarshalIndent(post, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)

		return
	}
	fmt.Println(string(postJson))
}
