package main

import (
	"fmt"
	"log"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(3)
}
