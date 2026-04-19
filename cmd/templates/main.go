package main

import (
	"log"
	"net/http"
	"os"

	"github.com/lukemassa/jclubtakeaways/internal/templater"
)

func main() {

	t := templater.New("src/templates")
	if len(os.Args) > 1 {
		if os.Args[1] == "--server" {
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, "/index.html", http.StatusMovedPermanently)
			})
			http.Handle("/{page}", t)
			log.Print("Listening on :8080")
			log.Fatal(http.ListenAndServe(":8080", nil))
		}
		log.Fatal("Usage: [--server]")
	}
	err := t.Write("docs")
	if err != nil {
		log.Fatal(err)
	}
}
