package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(os.Getenv("SECOND_URL"))
	if err != nil {
		fmt.Fprintf(w, "Error making request to second server: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading response from second server: %v\n", err)
		return
	}

	fmt.Fprintf(w, "Hello from first server!\n")
	fmt.Fprintf(w, "Response from second server: %s\n", string(body))
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Starting first server at port 80")
	if err := http.ListenAndServe("0.0.0.0:80", nil); err != nil {
		panic(err)
	}
}
