package main

import (
	"fmt"
	"log"
	"manhua-funage/db"
	"manhua-funage/service"
	"net/http"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	//http.HandleFunc("/", service.IndexHandler)
	//http.HandleFunc("/api/count", service.CounterHandler)
	http.HandleFunc("/", service.ValidateSign)

	log.Fatal(http.ListenAndServe(":80", nil))
}
