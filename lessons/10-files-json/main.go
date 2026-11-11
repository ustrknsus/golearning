// Files and JSON
// Exercise: Add a port field and encode/decode it.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Host    string `json:"host"`
	Workers int    `json:"workers"`
}

func run() error {
	dir, err := os.MkdirTemp("", "go-lesson-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	data, err := json.MarshalIndent(Config{Host: "localhost", Workers: 3}, "", "  ")
	if err != nil {
		return err
	}
	path := filepath.Join(dir, "config.json")
	if err := os.WriteFile(path, data, 0600); err != nil {
		return err
	}
	loaded, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var config Config
	if err := json.Unmarshal(loaded, &config); err != nil {
		return err
	}
	fmt.Printf("Loaded: %+v\n", config)
	return nil
}
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
