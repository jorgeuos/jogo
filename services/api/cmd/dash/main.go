package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
    log.Println("Starting server at port 8282")

    // Adds a HandleFunc to the DefaultMux
    http.HandleFunc("/up", func(w http.ResponseWriter, r *http.Request) {
        // Utilize fmt.Fprintf that accepts a io.Writer
    // and write the content of the page
        fmt.Fprintf(w, "Jogo bonito!")
    })

    if err := http.ListenAndServe(":8282", nil); err != nil {
        log.Fatal(err)
    }
}