package main

import (
	"errors"
	"fmt"
	//"io"
	"net/http"
)

func main() {
	//aeMain()
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/publishers", getPublishers)
	http.HandleFunc("/dsps", getDsps)

	err := http.ListenAndServe(":4444", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server two closed\n")
	} else if err != nil {
		fmt.Printf("error listening for server two: %s\n", err)
	}
}
