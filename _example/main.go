package main

import (
	"fmt"
	"os"

	"github.com/typetalk-gadget/go-typetalk-token-source/source"
)

func main() {
	ts := source.TokenSource{
		ClientID:     os.Getenv("TYPETALK_CLIENT_ID"),
		ClientSecret: os.Getenv("TYPETALK_CLIENT_SECRET"),
		Scope:        "my topic.read",
	}
	t, err := ts.Token()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("token: %#v\n", t)
}
