package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/v1/time", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		out := Resp{
			CurrentTime: time.Now().UTC(),
		}
		err := json.NewEncoder(w).Encode(out)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Success")

	})
	port := os.Getenv("PORT")
	fmt.Println("Listening on port " + port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), http.DefaultServeMux); err != nil {
		panic(err)
	}
}

type Resp struct {
	CurrentTime time.Time `json:"current_time"`
}
