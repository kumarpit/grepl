package main

import (
	"fmt"
	"github.com/kumarpit/grepl/regex2fsm"
	"github.com/kumarpit/grepl/fsm"
	"strings"
)

func main() {
	fmt.Println("This is grep(l)")

	pattern := os.Args[1]
	text := os.Args[2]

	converter := regex2fsm.New()
	machine, err := converter.convert(pattern)
	if err != nil {
		log.Fatal(err)
	}
	
	result := machine.Run(strings.split(text, ""))

	fmt.Println("%d", result)
}	