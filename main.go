package main

import (
	"fmt"
	"strings"
	"os"
	"log"
	"bufio"

	"github.com/kumarpit/grepl/fsm"
	"github.com/kumarpit/grepl/regex2fsm"
	"github.com/kumarpit/grepl/pfiles"

	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func main() {
	pattern := os.Args[1]
	root := os.Args[2]

	converter := regex2fsm.New()
	machine, err := converter.Convert(pattern)
	if err != nil {
		log.Fatal(err)
	}

	paths := pfiles.GetPaths(root)
	fmt.Println(paths)

	search(machine, paths)
}	

func search(machine *fsm.StateMachine, paths []string) {
	c := make(chan string, 100)

	for _, path := range paths {
		g.Go(func() error {
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
					fmt.Println(line)
					c <- " "
				}
				machine.Reset()
			}

			err = scanner.Err()
			if err != nil {
				log.Fatal(err)
			}

			return nil
		})
	}

	
	g.Wait()

	fmt.Println(len(c))

	// return g.Wait()
}