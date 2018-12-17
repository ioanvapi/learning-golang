package main

import (
	"fmt"
	"log"
	"math/big"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// Account 银行账户
type Account struct {
	Id      int64
	Balance string `xorm:"Decimal(32,16)"`
	Version int    `xorm:"version"` // 乐观锁
}

var x *xorm.Engine

func init() {
	var err error
	x, err = xorm.NewEngine("mysql", "root:password@tcp(10.10.0.122)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Fail to create engine: %v\n", err)
	}

	// 同步结构体与数据表
	if err = x.Sync(new(Account)); err != nil {
		log.Fatalf("Fail to sync database: %v\n", err)
	}
}

func main() {
	// insert
	n := new(big.Int)
	n, ok := n.SetString("1000", 10)
	if !ok {
		fmt.Println("SetString: error")
		return
	}
	_, err := x.Insert(&Account{Balance: n.String()})
	if err != nil {
		log.Fatalf("insert failed, err: %v", err)
	}

	// query
	a := &Account{}
	has, err := x.ID(1).Get(a)
	if err != nil {
		log.Fatal(err)
	} else if !has {
		log.Fatal("Account does not exist")
	}

	fmt.Printf("account balance: %v\n", a.Balance)
}
