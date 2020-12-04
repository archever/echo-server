package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, _ := json.Marshal(r.Header)
		w.Write(body)
	})
	fmt.Println("listen on 0.0.0.0:8000")
	http.ListenAndServe(":8000", nil)
}
