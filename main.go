package main

import (
	"log"

	"go-database/entity"
)

func main() {
	var tt = entity.Comment{
		Id: 1,
	}
	log.Println("Test Golang ...", tt)
}
