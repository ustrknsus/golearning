// HTTP clients, servers, and timeouts
// Exercise: Add a second route and decode a JSON request body.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

func run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"message": "Hello from Go"})
	})
	// An ephemeral loopback server keeps this lesson self-contained.
	server := httptest.NewServer(mux)
	defer server.Close()
	client := &http.Client{Timeout: 3 * time.Second}
	response, err := client.Get(server.URL + "/hello")
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP status: %s", response.Status)
	}
	var data map[string]string
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return err
	}
	fmt.Println(data["message"])
	return nil
}
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
