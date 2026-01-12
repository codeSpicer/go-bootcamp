package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 1. Make the GET Request
	// http.Get is a blocking call. It sends a request to the URL and
	// waits for the server to send back headers.
	url := "https://jsonplaceholder.typicode.com/posts/1"
	resp, err := http.Get(url)

	if err != nil {
		// This triggers if the URL is invalid or the server is down.
		fmt.Println("Error making request:", err)
		return
	}

	// 2. CRITICAL: Close the Response Body
	// The response body is a 'ReadCloser' stream. If you don't close it,
	// you will leak network connections and eventually crash your app.
	// We use 'defer' to ensure it closes even if the function exits early.
	defer resp.Body.Close()

	// 3. Inspect the Status Code
	// 200 means OK, 404 means Not Found, 500 means Server Error.
	fmt.Println("Status Code:", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
		return
	}

	// 4. Read the Response Body
	// The body is a stream (io.ReadCloser), not a string.
	// io.ReadAll reads the entire stream into a byte slice ([]byte).
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 5. Convert Bytes to String
	// Since JSON is text, we convert the []byte to a string for printing.
	fmt.Println("\n--- Response Body ---")
	fmt.Println(string(body))
}
