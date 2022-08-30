package pfiles

import (
	"os"
	"strings"
	"path/filepath"
)

// recursively find all .txt files
func GetPaths(root string) []string {
	paths := make(chan string, 100)
	list := []string{}

	go func() error {
		defer close(paths)

		return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			if !info.IsDir() && !strings.HasSuffix(info.Name(), ".txt") {
				return nil
			}

			path = strings.Replace(path, "\\", "/", -1)

			select {
			case paths <- path:
				return nil
			}
		})
	}()

	for path := range paths {
		list = append(list, path)
	}

	return list
}