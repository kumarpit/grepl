package main

import (
	
	
	"fmt"
	
	"os"
	"path/filepath"
	
	// "golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

// run grepl parallely on multiple files
func main() {
	paths := make(chan string, 100)
	root := os.Args[1]
	
	g.Go(func() error {
		defer close(paths)

		return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}

			select {
			case paths <- path:
				return nil
			}
		})
	})

	// c := make(chan string, 100)
	for path := range paths {
		fmt.Println(path)
	}
} 