package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./static.netelip.com"))))
	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
