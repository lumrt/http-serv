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

var keyServerAddr = "serverAdress"
// PATH handler Root
func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	fmt.Printf("%s : got / request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "This is my Webserv !\n")
	//debug print LocalIP
	ip := GetLocalIP()
	if ip == "" {
		fmt.Printf("No IP adress found")
	} else {
		fmt.Printf("local IP adress : %s\n\n", ip)
	}
}

// PATH handler test
func getTest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s : got /test resquest\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "Hi , HTTP !\n")
	//debug print LocalIP
	ip := GetLocalIP()
	if ip == "" {
		fmt.Printf("No IP adress found")
	} else {
		fmt.Printf("local IP adress : %s\n\n", ip)
	}
}

// PATH handler HTML
func getHTML(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s : go /HTML request !\n",  ctx.Value(keyServerAddr))
	io.WriteString(w, "Welcome om the main HTML index !\n")
	//debug print LocalIP
	ip := GetLocalIP()
	if ip == "" {
		fmt.Printf("No IP adress found")
	} else {
		fmt.Printf("local IP adress : %s\n\n", ip)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/test", getTest)
	mux.HandleFunc("/HTML", getHTML)
	
	ctx, cancelCtx := context.WithCancel(context.Background())
	serverOne := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	serverTwo := &http.Server{
		Addr:    ":4444",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}
	go func() {
		err := serverOne.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server one closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
		cancelCtx()
	}()
	go func() {
		err := serverTwo.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server two closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server two: %s\n", err)
		}
		cancelCtx()
	}()

	<-ctx.Done()
	// si l'erreur correspond au code erreur d'un serveur http ferme
	// if errors.Is(err, http.ErrServerClosed) {
	// 	fmt.Printf("server closed\n")
	// } else if err != nil {
	// 	fmt.Printf("error starting server: %s\n", err)
	// }
}	