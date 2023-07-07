package main

import (
	"encoding/json"
	"fmt"

	"github.com/antonprokopovich/threadsnet"
)

func main() {
	t := threads.NewThreads()

	post, err := t.GetPost(3141055616164096839)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	postJson, _ := json.MarshalIndent(post, "", "  ")
	fmt.Println(string(postJson))
}
