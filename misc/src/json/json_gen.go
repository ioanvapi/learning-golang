package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type Server struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"-"` //`json:"serverIP"`
}

type Server1 struct {
	// ID 不会导出到JSON中
	ID int `json:"-"`

	// ServerName 的值会进行二次JSON编码
	ServerName  string `json:"serverName"`
	ServerName2 string `json:"serverName2,string"`

	// 如果 ServerIP 为空，则不输出到JSON串中
	ServerIP string `json:"serverIP,omitempty"`
}

type Serverslice struct {
	Servers []Server `json:"servers"`
}

func errHndlr(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func main() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(string(b))

	s1 := Server1{
		ID:          3,
		ServerName:  `Go "1.0" `,
		ServerName2: `Go "1.0" `,
		ServerIP:    ``,
	}
	b1, err := json.Marshal(s1)
	errHndlr(err)
	os.Stdout.Write(b1)
}
