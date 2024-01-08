package main

import (
	"fmt"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func main() {
	http.HandleFunc("/get", getHandler)
	fmt.Println("Server starting on port 9090...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
