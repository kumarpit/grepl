package main

import (
	"fmt"
	"strings"
	"os"
	"log"
	"bufio"
	"sync"

	"github.com/kumarpit/grepl/regex2fsm"
	"github.com/kumarpit/grepl/pfiles"

	// "golang.org/x/sync/errgroup"
)

var g sync.WaitGroup
var converter *regex2fsm.Parser
var c = make(chan string, 100)

func main() {
	pattern := os.Args[1]
	root := os.Args[2]

	converter = regex2fsm.New()

	paths := pfiles.GetPaths(root)
	fmt.Println(paths)

	search(paths, pattern)
}

func routine(path string, pattern string) {
	defer g.Done()
	fmt.Println(path)

	// create a new machine for each thread
	machine, err := converter.Convert(pattern)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path)
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
			fmt.Println(path, line)
			c <- " "
		}
		machine.Reset()
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func search(paths []string, pattern string) {
	g.Add(len(paths))

	for _, path := range paths {
		go routine(path, pattern)
	}
	
	g.Wait()
	fmt.Println(len(c))
}