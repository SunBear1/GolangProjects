package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func validateInput(path string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Fatalf("Error. Root path does not exist.")
		}
		if errors.Is(err, os.ErrPermission) {
			log.Printf("Warning. Permission denied of the root path.")
		}
	}
	if !fileInfo.IsDir() {
		log.Fatalf("Error. Root path points to a file.")
	}
}

func buildTree(path string, prefix string) {
	if len(prefix) > 30 {
		log.Fatalf("Error. Heap overload. Escaping...")
	}
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	for i, entry := range dirEntries {
		connector := "├──"
		if i == len(dirEntries)-1 {
			connector = "└──"
		}
		if entry.Type().IsDir() {
			fmt.Printf("%s%s📁 %s\n", prefix, connector, entry.Name())
			dirPath := filepath.Join(path, entry.Name())
			buildTree(dirPath, prefix+"│   ")
		} else {
			fmt.Printf("%s%s📄 %s\n", prefix, connector, entry.Name())
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter root path: ")
	inputPath, _ := reader.ReadString('\n')
	inputPath = strings.TrimSuffix(inputPath, "\n")
	validateInput(inputPath)
	buildTree(inputPath, "")
}
