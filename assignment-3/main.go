package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Data struct {
	Status Status `json:"status"`
}

type Status struct {
	Wind  int `json:"wind"`
	Water int `json:"water"`
}

func updateJSON() {
	water, wind := rand.Intn(100)+1, rand.Intn(100)+1

	data := Data{
		Status: Status{
			Water: water,
			Wind:  wind,
		},
	}

	dataByte, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error!")
	}

	_ = os.WriteFile("data.json", dataByte, 0644)
	fmt.Println("success!")
}

func getWaterStatus(value int) string {
	if value <= 5 {
		return "aman"
	}
	if value >= 6 && value <= 8 {
		return "siaga"
	}
	return "bahaya"
}

func getWindStatus(value int) string {
	if value <= 6 {
		return "aman"
	}
	if value >= 7 && value <= 15 {
		return "siaga"
	}
	return "bahaya"
}

func index(w http.ResponseWriter, r *http.Request) {
	fileData, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("error!")
		return
	}

	var data Data
	json.Unmarshal(fileData, &data)

	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("error!")
		return
	}
	waterStatus, windStatus := getWaterStatus(data.Status.Water), getWindStatus(data.Status.Wind)
	t.Execute(w, map[string]interface{}{
		"water":       data.Status.Water,
		"waterStatus": waterStatus,
		"wind":        data.Status.Wind,
		"windStatus":  windStatus,
	})
}

func main() {
	http.HandleFunc("/", index)

	ticker := time.NewTicker(15 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				updateJSON()
			}
		}
	}()

	http.ListenAndServe(":8080", nil)
}
