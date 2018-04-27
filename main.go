package main

import (
	"fmt"

	"github.com/prantoran/go-mock-microservice/printer"
)

func main() {
	p := printer.New("localhost:4242")
	res, err := p.Print("Friday")
	fmt.Println("res:", res, " err:", err)
}
