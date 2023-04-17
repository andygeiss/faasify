package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/andygeiss/faasify/internal/http/server"
)

func main() {
	// Define the CLI flags
	data := flag.String("data", "", "data used as input")
	host := flag.String("host", "http://127.0.0.1:3000", "host of the functions")
	name := flag.String("name", "status", "name of the function")
	flag.Parse()
	// Create a new client with a timeout
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	// Call the function
	body := bytes.NewReader([]byte(*data))
	req, _ := http.NewRequest("POST", *host+"/function/"+*name, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+server.Token)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// Read the result
	content, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Write it to stdout
	fmt.Printf("%s", content)
}
