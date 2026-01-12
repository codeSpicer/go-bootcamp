package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	const port string = ":8080"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello server!")
	})

	fmt.Printf("Server is listening on port%v\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("error starting server: ", err)
	}

}
