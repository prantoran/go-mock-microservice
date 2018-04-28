package main

import (
	"fmt"

	"github.com/prantoran/go-mock-microservice/printer"
)

func main() {
	pr := printer.New("localhost:4242")
	printer.Cur = pr
	res, _ := printer.Cur.Print("Friday")
	fmt.Println("body:", res.Body, " status:", res.Status)
}
