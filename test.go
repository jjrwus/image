package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	files, err := ioutil.ReadDir("./img")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}


machine source.developers.google.com login Juan.F.Nadal@gmail.com password 1/GF5l0A7VlHeU-lajAKrk6bDri0A7jL-WXVrlelPdYVM