package main

import (
  "bytes"
  "fmt"

  "github.com/eoscanada/eos-go"
  "net/url"
)

func main() {
  eosRPC, _ := url.Parse("http://ortc.io:8888")
	api := eos.New(eosRPC, bytes.Repeat([]byte{0}, 32))
	infoResp, err := api.GetInfo()
	if err != nil {
		log.Fatalf("EOS get info failed, err: %v", err)
	}
	log.Infof("EOS get info: %+v", infoResp)

	tableResp, err := api.GetTableRows(eos.GetTableRowsRequest{
		Scope: "EOS",
		Code: "eosio.token",
		Table: "stat",
		JSON: true,
	})

	tableRows, _ := tableResp.Rows.MarshalJSON()

	log.Infof("table response: %v", string(tableRows))
}
