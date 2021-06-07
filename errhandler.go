package main

import (
	"fmt"
	"log"
	"net/http"
)

func UnableLoad (w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Fprint(w, "Unable to load page.")
		log.Fatal(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}