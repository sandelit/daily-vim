package models

import (
	"encoding/json"
	"errors"
	"math/rand"
	"os"
	"time"
)

type Tip struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Command string `json:"command,omitempty"`
	Link    string `json:"link,omitempty"`
}

var tips []Tip

func init() {
	file, err := os.ReadFile("data/tips.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &tips)
	if err != nil {
		panic(err)
	}

	for index := range tips {
		tips[index].Id = index + 1
	}
}

func GetTipOfTheDay() Tip {
	// Use the current date as seed for random number generator
	now := time.Now()
	seed := now.Year()*10000 + int(now.Month())*100 + now.Day()
	r := rand.New(rand.NewSource(int64(seed)))

	return tips[r.Intn(len(tips))]
}

func GetAllTips() []Tip {
	return tips
}

func GetTipByID(id int) (Tip, error) {
	id--
	if id < 0 || id >= len(tips) {
		return Tip{}, errors.New("tip not found")
	}
	return tips[id], nil
}
