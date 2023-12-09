package main

import (
	"algopack/cmd/algopack"
	"algopack/internal/parse"
	"fmt"
	"log"
	"net/http"
)

func main() {
	if err := algopack.RootCommand.Execute(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		latestTicketDataJSON := parse.ParseTicketData(w, r)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write(latestTicketDataJSON)
		if err != nil {
			return
		}
	})

	fmt.Println("Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
