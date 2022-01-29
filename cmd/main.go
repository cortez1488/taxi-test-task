package main

import (
	"log"
	reqToAPI "taxiTestTask/internal/reqToAPI/JSON"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	err := reqToAPI.RequestJSON()
	if err != nil {
		log.Fatal(err.Error())
	}
}
