package main

import (
	"encoding/json"
	"log"
	"time"
)

type Person struct {
	Name        string `json:"Имя"`
	Email       string `json:"Почта"`
	DateOfBirth time.Time
}

func main() {
	person := Person{
		Name:        "Aлекс",
		Email:       "alex@yandex.ru",
		DateOfBirth: time.Now(),
	}

	ser, err := json.MarshalIndent(person, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(ser))
}
