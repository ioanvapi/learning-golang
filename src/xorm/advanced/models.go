package main

import (
	"errors"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// Account 银行账户
type Account struct {
	Id      int64
	Name    string `xorm:"unique"`
	Balance float64
	Version int `xorm:"version"` // 乐观锁
}

// BeforeInsert before insert
func (a *Account) BeforeInsert() {
	log.Printf("before insert: %s", a.Name)
}

// AfterInsert after insert
func (a *Account) AfterInsert() {
	log.Printf("after insert: %s", a.Name)
}

// ORM 引擎
var x *xorm.Engine

func init() {
	// 创建 ORM 引擎与数据库
	var err error
	x, err = xorm.NewEngine("mysql", "root:password@tcp(10.10.0.122)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Fail to create engine: %v\n", err)
	}

	// 同步结构体与数据表
	if err = x.Sync(new(Account)); err != nil {
		log.Fatalf("Fail to sync database: %v\n", err)
	}

	// 记录日志
	f, err := os.Create("sql.log")
	if err != nil {
		log.Fatalf("Fail to create log file: %v\n", err)
		return
	}
	x.SetLogger(xorm.NewSimpleLogger(f))
	x.ShowSQL()

	// 设置默认 LRU 缓存
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	x.SetDefaultCacher(cacher)
}

// 创建新的账户
func newAccount(name string, balance float64) error {
	// 对未存在记录进行插入
	_, err := x.Insert(&Account{Name: name, Balance: balance})
	return err
}

// 获取账户数量
func getAccountCount() (int64, error) {
	return x.Count(new(Account))
}

// 获取账户信息
func getAccount(id int64) (*Account, error) {
	a := &Account{}
	// 直接操作 ID 的简便方法
	has, err := x.ID(id).Get(a)
	// 判断操作是否发生错误或对象是否存在
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("Account does not exist")
	}
	return a, nil
}
