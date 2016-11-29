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

	someRecord, err := goavro.NewRecord(goavro.RecordSchema(recordSchemaJSON))
	if err != nil {
		log.Fatal(err)
	}

	// identify field name to set datum for
	someRecord.Set("username", "Aquaman")
	someRecord.Set("comment", "The Atlantic is oddly cold this morning!")
	// you can fully qualify the field name
	someRecord.Set("com.example.timestamp", int64(1082196484))

	log.Printf("Some record: %v", someRecord.String())

	codec, err := goavro.NewCodec(recordSchemaJSON)
	if err != nil {
		log.Fatal(err)
	}

	bb := new(bytes.Buffer)
	if err = codec.Encode(bb, someRecord); err != nil {
		log.Fatal(err)
	}

	actual := bb.Bytes()
	expected := []byte("\x0eAquamanPThe Atlantic is oddly cold this morning!\x88\x88\x88\x88\x08")
	if bytes.Compare(actual, expected) != 0 {
		log.Printf("Actual: %#v; Expected: %#v", actual, expected)
	}
}
