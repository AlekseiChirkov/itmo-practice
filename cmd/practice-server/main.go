package main

import (
	"log"
	"net/http"
	"practice-server/internal/practice-server/data"
)

func main() {
	data.Init()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/glossary", glossary)
	mux.HandleFunc("/mind-map", mindMap)

	log.Println("Запуск сервера на http://127.0.0.1:4000")

	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Fatalln(err)
	}
}
