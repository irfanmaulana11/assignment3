package main

import (
	"assignment3/handlers"
	"assignment3/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	go Run()
	go forever()

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
}

func Run() {
	http.HandleFunc("/", handlers.GetWeatherView)
	http.ListenAndServe(":8080", nil)
}

func WriteFileJson() {

	data := models.Weather{
		Status: models.Status{
			Water: GetRandomInt(),
			Wind:  GetRandomInt(),
		},
	}

	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
	}

	_ = ioutil.WriteFile("weather.json", file, 0644)
}

func GetRandomInt() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	randInt := rand.Intn(max-min) + min
	return randInt
}

func forever() {
	for {
		WriteFileJson()
		fmt.Printf("%s+\n", time.Now())
		time.Sleep(time.Second)
	}
}
