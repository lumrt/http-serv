package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	// "net"
)

var ctx = context.Background()

const keyServerAddr = "serverAddr"


func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// Vérifiez le type d'adresse et si elle n'est pas bouclante
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil { // Assurez-vous que l'adresse est IPv4
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	fmt.Printf("%s : got / request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "This is my Webserv !\n")
}

func getTest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s : got /test resquest\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "Hi , HTTP !\n")
}

func getHTML(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s : go /HTML request !",  ctx.Value(keyServerAddr))
	io.WriteString(w, "Welcome om the main HTML index !\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/test", getTest)
	mux.HandleFunc("/HTML", getHTML)
	err := http.ListenAndServe(":3333", mux)
	// si l'erreur correspond au code erreur d'un serveur http ferme
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
	}
}	