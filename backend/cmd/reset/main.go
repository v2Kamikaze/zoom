package main

import (
	"os"
	"path/filepath"
)

func main() {
	patterns := []string{"./assets/go-*", "./assets/zero-*", "./assets/lena-*", "./assets/car-*"}

	for _, pattern := range patterns {
		files, err := filepath.Glob(pattern)
		if err != nil {
			panic(err)
		}
		for _, f := range files {
			if err := os.Remove(f); err != nil {
				panic(err)
			}
		}
	}
}
