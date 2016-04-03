package main

import (
	"bufio"
	"fmt"
	"os"
)

type Result struct {
	count int
	files []string
}

func main() {
	counts := make(map[string]*Result)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, res := range counts {
		if res.count > 1 {
			fmt.Printf("%d\t%s\t%v\n", res.count, line, res.files)
		}
	}
}

func countLines(f *os.File, counts map[string]*Result) {
	input := bufio.NewScanner(f)
	name := f.Name()
	for input.Scan() {
		text := input.Text()
		if _, ok := counts[text]; !ok {
			counts[text] = new(Result)
		}
		counts[text].count++
		if counts[text].count > 1 {
			counts[text].files = append(counts[text].files, name)
		}
	}
}
