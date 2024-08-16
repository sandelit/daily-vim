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
		log.Fatal("Error reading tips file:", err)
	}

	err = json.Unmarshal(file, &tips)
	if err != nil {
		log.Fatal("Error parsing tips JSON:", err)
	}
}

func GetRandomTip() Tip {
	rand.Seed(time.Now().UnixNano())
	return tips[rand.Intn(len(tips))]
}
