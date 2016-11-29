package main

import (
	"bytes"
	"log"

	"github.com/linkedin/goavro"
)

func main() {
	recordSchemaJSON := `{
		"type": "record",
		"name": "comments",
		"doc:": "A basic schema for storing blog comments",
		"namespace": "com.example",
		"fields": [
		{
			"doc": "Name of user",
			"type": "string",
			"name": "username"
		},
		{
			"doc": "The content of the user's message",
			"type": "string",
			"name": "comment"
		},
		{
			"doc": "Unix epoch time in milliseconds",
			"type": "long",
			"name": "timestamp"
		}
		]
	}`

	codec, err := goavro.NewCodec(recordSchemaJSON)
	if err != nil {
		log.Fatal(err)
	}

	encoded := []byte("\x0eAquamanPThe Atlantic is oddly cold this morning!\x88\x88\x88\x88\x08")
	bb := bytes.NewBuffer(encoded)
	decoded, err := codec.Decode(bb)

	log.Println(decoded)

	// but direct access to data is provided
	record := decoded.(*goavro.Record)
	log.Println("Record Name:", record.Name)
	log.Println("Record Fields:")

	for i, field := range record.Fields {
		log.Println(" field", i, field.Name, ":", field.Datum)
	}

}
