package main

import (
	"io"
	"log"
	"os"

	"github.com/linkedin/goavro"
)

func dumpReader(r io.Reader) {
	fr, err := goavro.NewReader(goavro.BufferFromReader(r))
	if err != nil {
		log.Fatalf("goavro new reader failed, err: %v", err)
	}
	defer func() {
		if err := fr.Close(); err != nil {
			log.Fatalf("file close failed, err: %v", err)
		}
	}()

	for fr.Scan() {
		datum, err := fr.Read()
		if err != nil {
			log.Printf("file read failed, err: %v", err)
			continue
		}
		log.Println(datum)
	}
}

func main() {
	if len(os.Args) > 1 {
		for i, arg := range os.Args {
			if i == 0 {
				continue
			}
			fh, err := os.Open(arg)
			if err != nil {
				log.Fatalf("file open failed, err: %v", err)
			}
			dumpReader(fh)
			fh.Close()
		}
	} else {
		dumpReader(os.Stdin)
	}
}
