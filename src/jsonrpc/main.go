package main

import (
	"github.com/ybbus/jsonrpc"
)

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	rpcClient := jsonrpc.NewClient("http://sss:8080/rpc")

	var person *Person
	rpcClient.CallFor(&person, "getPersonById", 4711)

	person.Age = 33
	rpcClient.Call("updatePerson", person)
}
