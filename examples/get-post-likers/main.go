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

	likers, err := t.GetPostLikers(3141002295235099165)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	var likersPretty bytes.Buffer
	if err = json.Indent(&likersPretty, likers, "", "\t"); err != nil {
		log.Fatal("JSON parse error: ", err)
	}

	fmt.Println(likersPretty.String())
}
