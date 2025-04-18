package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Step 1: Send an HTTP GET request
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		// Step 2: Handle errors
		fmt.Println("Error:", err)
		return
	}
	// Step 3: Close the response body
	defer resp.Body.Close()

	// Step 4: Check the HTTP status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: HTTP request failed with status", resp.Status)
		return
	}

	// Step 5: Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// Step 6: Handle errors while reading
		fmt.Println("Error reading response body:", err)
		return
	}

	// Step 7: Process the data (convert to string)
	data := string(body)

	// Step 8: Output the data
	fmt.Println("HTTP response body:")
	fmt.Println(data)
}
