package main

import (
	"fmt"
	"io"
	"os"
	// "path/filepath"
	// "strings"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, patch string, printFiles bool) error {
	return renderTree(out, patch, printFiles, "")
}

func renderTree(out io.Writer, patch string, printFiles bool, prefix string) error {
	files, err := os.ReadDir(patch)
	if err != nil {
		return fmt.Errorf("error read files: %w", err)
	}

	for _, item := range files {
		fmt.Fprintf(out, "├%s\n", item.Name())
	}

	return nil
}