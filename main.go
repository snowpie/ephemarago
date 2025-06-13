package main

import (
	"fmt"
     "net/http"
)

var Secrets = make(map[string]func() string)


func Add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(Add(5, 6))
	// Create a map to hold secrets
	// Secrets := make(map[string]func() string)


	http.HandleFunc("/", frontpage)
	http.HandleFunc("/secret", sendmeyoursecret)
	http.HandleFunc("/hello", httphandler)

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

// secret returns a closure that captures the secret string.
// The returned function can be called to retrieve the secret.
func secret(mysecret string) func() string {
	secret := mysecret
	return func() string {
		return secret
	}
}

func sendmeyoursecret(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Get the secret from the form data
	mysecret := r.FormValue("secret")
	if mysecret == "" {
		http.Error(w, "Secret is required", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Your secret is: %s\n", mysecret)
	fmt.Println("Received secret:", mysecret)
	key:=stashsecret(mysecret)
	fmt.Fprintf(w, "Your key is: %s\n", key)

}

func stashsecret(mysecret string) string {

	fmt.Println("Stashing secret:", mysecret)
	Secrets["key"] = secret(mysecret)
	fmt.Println("Secret stored successfully!")
	return "key"
}

func getsecret(key string) string {
	// Retrieve the secret using the key
	if secret, exists := Secrets[key]; exists {
		return secret()
	}
	return ""
}

func httphandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! ")
}

func frontpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><h1>Welcome to the homepage!</h1>")
	fmt.Fprintf(w, "<p>Enter your secret</p>")
	fmt.Fprintf(w, "<form action=\"/secret\" method=\"post\">")
	fmt.Fprintf(w, "<input type=\"text\" name=\"secret\" placeholder=\"Enter your secret\" required>")
	fmt.Fprintf(w, "<input type=\"submit\" value=\"Submit\">")
	fmt.Fprintf(w, "</form>")
	fmt.Fprintf(w, "</body></html>")

	fmt.Println("Homepage accessed")

}