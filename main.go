package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sahilarora535/project-solutions-go/numbers"
)

func main() {
	log.SetPrefix("project-solutions-go: ")
	log.SetFlags(0)

	if len(os.Args) < 2 {
		log.Fatalln("usage: project-solutions-go numDigits")
	}
	numDigits, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	result, err := numbers.DigitsOfPi(numDigits)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)
}
