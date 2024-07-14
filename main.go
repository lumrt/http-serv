package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my Webserv !\n")
}

func getTest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /test resquest\n")
	io.WriteString(w, "Hi , HTTP !\n")
}

func main() {
	http.HandleFunc("/" , getRoot)
	http.HandleFunc("/test", getTest)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}	