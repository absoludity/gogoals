package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
    var port int
	flag.IntVar(&port, "port", 8000, "The port to which the server listens.")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        defer log.Println(r.URL.Path)
		fmt.Fprint(w, "Serving away on port ", port)
	})

    log.Println("Serving OpenGoalTracker on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), nil))
}
