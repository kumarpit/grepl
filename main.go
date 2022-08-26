package main

import (
	"fmt"
	"github.com/kumarpit/grepl/regex2fsm"
	"strings"
	"os"
	"log"
	"bufio"
)

func main() {
	// fmt.Println("This is grep(l)")

	pattern := os.Args[1]
	filename := os.Args[2]

	converter := regex2fsm.New()
	machine, err := converter.Convert(pattern)
	if err != nil {
		log.Fatal(err)
	}
	
	// search in file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, "")
		result := machine.Run(tokens)
		if result {
			fmt.Println(line)
		}
		machine.Reset()
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}	