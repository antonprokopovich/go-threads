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
		log.Fatal("NewThreads: ", err)
	}

	likers, err := t.GetPostLikers(3141055616164096839)
	if err != nil {
		log.Fatal("GetPostLikers: ", err)
	}

	var likersPretty bytes.Buffer
	if err = json.Indent(&likersPretty, likers, "", "\t"); err != nil {
		log.Fatal("json.Indent: ", err)
	}

	fmt.Println(likersPretty.String())
}
