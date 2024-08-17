package models

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

type Tip struct {
	Content string `json:"content"`
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
