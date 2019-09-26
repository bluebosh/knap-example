package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

// Calculate pi using Gregory-Leibniz series:   (4/1) - (4/3) + (4/5) - (4/7) + (4/9) - (4/11) + (4/13) - (4/15) ...
func say(word string) string {
    title := "[20190926] Hello World"
    return title + " " + word
}

func handler(w http.ResponseWriter, r *http.Request) {
    log.Print("Say something!")
    word := r.URL.Query()["word"][0]
    fmt.Fprintf(w, "%s", say(word))
}

func main() {
    log.Print("Say something started.")

    http.HandleFunc("/", handler)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
