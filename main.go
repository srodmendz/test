package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Running on 8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
