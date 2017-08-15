package main

import (
  "fmt"
)

func f() (ret string) {
  defer func() {
    if r := recover(); r != nil {
      ret = fmt.Sprintf("was panic, recovered value: %v", r)
    }
  }()
  panic("test")
  return "Normal Return Value"
}

func main() {
  fmt.Println("Returned:", f())
}
