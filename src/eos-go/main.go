package main

import (
  "bytes"
  "fmt"

  "github.com/eoscanada/eos-go"
  "net/url"
)

func main() {
  api := eos.New(&url.URL{Scheme: "http", Host: "ortc.io:8888"}, bytes.Repeat([]byte{0}, 32))
  infoResp, _ := api.GetInfo()

  fmt.Printf("infoResp: %+v\n", infoResp)
  accountResp, _ := api.GetAccount("eosio")
  fmt.Println("Permission for eosio:", accountResp.Permissions[0].RequiredAuth.Keys)
}
