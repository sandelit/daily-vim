package models

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

type Tip struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Command string `json:"command",omitempty`
	Link    string `json:"link,omitempty"`
}

var tips []Tip

func init() {
	loadTips()
}

func loadTips() {
	file, err := os.ReadFile("data/tips.json")
	if err != nil {
		log.Fatalf("Error reading tips file: %v", err)
	}

	if len(file) == 0 {
		log.Fatal("Tips file is empty")
	}

	err = json.Unmarshal(file, &tips)
	if err != nil {
		log.Fatalf("Error parsing tips JSON: %v\nJSON content: %s", err, string(file))
	}

	log.Printf("Loaded %d tips", len(tips))
}

func GetRandomTip() Tip {
	rand.Seed(time.Now().UnixNano())
	return tips[rand.Intn(len(tips))]
}

func GetAllTips() []Tip {
	return tips
}
