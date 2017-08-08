package main

import (
  "time"
  "fmt"
)

func main() {
  // defer func() {
  //   // 必须要先声明 defer, 否则不能捕获到 panic 异常
  //   fmt.Println("2")
  //   if err := recover(); err != nil {
  //     fmt.Println(err) // 这里的 err 其实就是 panic 传入的内容, bug
  //   }
  //   fmt.Println("3")
  // }()
  f()
}

func f() {
  defer func() {
    // 必须要先声明 defer, 否则不能捕获到 panic 异常
    fmt.Println("2")
    if err := recover(); err != nil {
      fmt.Println(err) // 这里的 err 其实就是 panic 传入的内容, bug
    }
    fmt.Println("3")
  }()
  for {
    fmt.Println("1")
    a := []string{"a", "b"}
    fmt.Println(a[3]) // 这里 slice 越界异常了
    // panic("bug")
    fmt.Println("4") // 不会运行
    time.Sleep(1 * time.Second)
  }
}
