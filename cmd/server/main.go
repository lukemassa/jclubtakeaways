package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lukemassa/jclubtakeaways/internal/templater"
	"github.com/lukemassa/jclubtakeaways/internal/token"
)

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":3000"
}

func getTokener() (*token.Tokener, error) {
	content := os.Getenv("WEB_CLIENT_KEY")
	if content == "" {
		return nil, errors.New("WEB_CLIENT_KEY is not set, skipping token setup")
	}
	tokener, err := token.New(content)
	if err != nil {
		return nil, fmt.Errorf("creating tokener: %w", err)
	}
	return &tokener, nil
}

func main() {

	templater := templater.New("src/templates")
	tokener, err := getTokener()
	if err != nil {
		log.Print(err)
		log.Print("Continuing with setup, but Submit will not work")
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/index.html", http.StatusTemporaryRedirect)
	})
	http.Handle("/{page}", templater)
	if tokener != nil {
		http.Handle("/token", tokener)
	}
	log.Printf("Successfully built templates")
	port := getPort()
	log.Printf("Listening on %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
