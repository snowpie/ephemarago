package main

import (
	"fmt"
     "net/http"
)

func Add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(Add(5, 6))
	http.HandleFunc("/", httphandler)
	go func(port string) {
		fmt.Println("Starting HTTP server...")
		if err := http.ListenAndServe(port, nil); err != nil {
			fmt.Printf("Error starting HTTP server: %v\n", err)
		}
	}(":8080")
	// Wait for the server to start
	fmt.Println("HTTP server started on port 8080")
	// Keep the main goroutine alive
	select {}		
}

func secret(mysecret string) func() string {
	secret := mysecret
	return func() string {
		return secret
	}
}

func httphandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! ")
	fmt.Println("Received request from:", r.RemoteAddr)
	fmt.Println("Request method:", r.Method)
	fmt.Println("Request URL:", r.URL)
	fmt.Println("Request headers:", r.Header)
	fmt.Println("Request body:", r.Body)
	fmt.Println("Request content length:", r.ContentLength)
	fmt.Println("Request host:", r.Host)
	fmt.Println("Request protocol:", r.Proto)
	fmt.Println("Request user agent:", r.UserAgent())
	fmt.Println("Request referer:", r.Referer())
	fmt.Println("Request form:", r.Form)
	fmt.Println("Request post form:", r.PostForm)
	fmt.Println("Request trailer:", r.Trailer)
	fmt.Println("Request context:", r.Context())
	fmt.Println("Request cookies:")
	for _, cookie := range r.Cookies() {
		fmt.Println(" -", cookie.Name, ":", cookie.Value)
	}
	fmt.Println("Request URL path:", r.URL.Path)
	fmt.Println("Request URL query:", r.URL.Query())
}
