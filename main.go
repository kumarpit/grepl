package main

import (
	"fmt"
	"github.com/kumarpit/grepl/regex2fsm"
	"strings"
	"os"
	"log"
)

func main() {
	fmt.Println("This is grep(l)")

	pattern := os.Args[1]
	text := os.Args[2]

	converter := regex2fsm.New()
	machine, err := converter.Convert(pattern)
	if err != nil {
		log.Fatal(err)
	}
	
	result := machine.Run(strings.Split(text, ""))

	fmt.Println(result)
}	