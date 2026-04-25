package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/lukemassa/jclubtakeaways/internal/token"
)

func main() {
	content := os.Getenv("WEB_CLIENT_KEY")
	decoded, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		log.Fatal(err)
	}
	tokener, err := token.New(string(decoded))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tokener.Get())
}
