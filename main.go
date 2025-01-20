package main

import (
	"devmentor-BE103-golang/infrastructure"
	"devmentor-BE103-golang/route"
	"log"
)

func main() {

	dbErr := infrastructure.InitMySQL()
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	_, err := route.InitGinServer()
	if err != nil {
		log.Fatal(err)
	}
}
