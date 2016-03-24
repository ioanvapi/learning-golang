package main

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

func main() {
	jsonParsedObj, _ := gabs.ParseJSON([]byte(`{
  "outter":{
      "inner":{
          "value1":10,
          "value2":22
      },
      "alsoInner":{
          "value1":20
      }
  }
}`))

	jsonOutput := jsonParsedObj.Search("outter").String()
	fmt.Println(jsonOutput)
}
