package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net"
)
// defalut value too block err flag unused 
var ctx = context.Background()
// loop and get local server IP 
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {// Check if IPV4 addr
			return ipnet.IP.String()
		}
	}
}
return ""
}

var keyServerAddr = GetLocalIP
// PATH handler Root
func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	fmt.Printf("%s : got / request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "This is my Webserv !\n")
}

// PATH handler test
func getTest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s : got /test resquest\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "Hi , HTTP !\n")
}

// PATH handler HTML
func getHTML(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s : go /HTML request !",  ctx.Value(keyServerAddr))
	io.WriteString(w, "Welcome om the main HTML index !\n")
	ip := GetLocalIP()
	if ip == "" {
		fmt.Printf("No IP adress found")
	} else {
		fmt.Printf("local IP adress : ", ip)
	}
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