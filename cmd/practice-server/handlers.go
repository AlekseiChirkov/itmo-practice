package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"practice-server/internal/practice-server/data"
)

type Data struct {
	Glossary map[string]string
	MindMap  [][]string
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, Data{
		Glossary: data.Glossary,
		MindMap:  data.MindMap,
	})

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func glossary(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/glossary" {
		http.NotFound(w, r)
		return
	}

	glossaryJsonBytes, err := json.Marshal(data.Glossary)
	if err != nil {
		fmt.Printf("Error while serializing response %s\n", data.Glossary)
		return
	}

	if _, err := w.Write(glossaryJsonBytes); err != nil {
		log.Printf("Error while responding to request: %s", err)
	}
}

func mindMap(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/mind-map" {
		http.NotFound(w, r)
		return
	}

	mindMapJsonBytes, err := json.Marshal(data.MindMap)
	if err != nil {
		fmt.Printf("Error while serializing response %s\n", data.MindMap)
		return
	}

	if _, err := w.Write(mindMapJsonBytes); err != nil {
		log.Printf("Error while responding to request: %s", err)
	}
}
