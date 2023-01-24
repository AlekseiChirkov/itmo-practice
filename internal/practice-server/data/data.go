package data

import (
	"embed"
	_ "embed"
	"encoding/json"
	"log"
)

var Glossary map[string]string
var MindMap [][]string

//go:embed glossary.json
var glossaryFS embed.FS

//go:embed mind-map.json
var mindMapFS embed.FS

func Init() {
	glossaryJson, _ := glossaryFS.ReadFile("glossary.json")
	if err := json.Unmarshal(glossaryJson, &Glossary); err != nil {
		log.Fatalln(err)
	}

	mindMapJson, _ := mindMapFS.ReadFile("mind-map.json")
	if err := json.Unmarshal(mindMapJson, &MindMap); err != nil {
		log.Fatalln(err)
	}
}
