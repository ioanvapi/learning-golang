package main

import (
	"io"
	"log"

	"fmt"
	"os"

	"github.com/linkedin/goavro"
)

var (
	schema string
	codec  goavro.Codec
)

func init() {
	schema = `{
		"type" : "record",
		"name" : "Weather",
		"namespace" : "test",
		"doc" : "A weather reading.",
		"fields": [{
			"name": "station",
			"type": "string"
		}, {
			"name": "time",
			"type": "string",
		}, {
			"name": "temp",
			"type": "int"
		}]
	}`

	var err error
	// If you want speed, create the codec one time for each
	// schema and reuse it to create multiple Writer instances.
	codec, err = goavro.NewCodec(schema)
	if err != nil {
		log.Fatalln(err)
	}
}

func dumpWriter(w io.Writer, codec goavro.Codec) {
	fw, err := codec.NewWriter(
		// goavro.Compression(goavro.CompressionDeflate),
		goavro.Compression(goavro.CompressionSnappy),
		goavro.ToWriter(w),
	)

	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		err := fw.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	raw := []map[string]interface{}{
		{"station": "011990-99999", "time": int64(-619524000000), "temp": int32(0)},
		{"station": "011990-99999", "time": int64(-619506000000), "temp": int32(22)},
		{"station": "011990-99999", "time": int64(-619484400000), "temp": int32(-11)},
		{"station": "012650-99999", "time": int64(-655531200000), "temp": int32(111)},
		{"station": "012650-99999", "time": int64(-655509600000), "temp": int32(78)},
	}
	for _, rec := range raw {
		record, err := goavro.NewRecord(goavro.RecordSchema(schema))
		if err != nil {
			log.Fatal(err)
		}
		for k, v := range rec {
			record.Set(k, v)
		}
		fw.Write(record)
	}
}

func main() {
	switch len(os.Args) {
	case 1:
		dumpWriter(os.Stdout, codec)
	case 2:
		fh, err := os.Create(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		dumpWriter(fh, codec)
		fh.Close()
	default:
		fmt.Fprintf(os.Stderr, "usage: %s [filename]\n", os.Args[0])
		os.Exit(2)
	}
}
