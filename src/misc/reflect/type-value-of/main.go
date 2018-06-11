package main

import (
	"fmt"
	"reflect"
)

type ControllerInterface interface {
	Init(action string, method string)
}

type Controller struct {
	Action string
	Method string
	Tag    string `json:"tag"`
}

func (c *Controller) Init(action string, method string) {
	c.Action = action
	c.Method = method
	// 增加fmt打印，便于看是否调用
	fmt.Println("Init() is run.")
	fmt.Println("c:",c)
}

//增加一个无参数的Func
func (c *Controller) Test(){
	fmt.Println("Test() is run.")
}

func main() {
	//初始化
	runController := &Controller{
		Action: "Run1",
		Method: "GET",
	}

	//Controller实现了ControllerInterface方法,因此它就实现了ControllerInterface接口
	var i ControllerInterface
	i = runController

	// 得到实际的值,通过v我们获取存储在里面的值,还可以去改变值
	v := reflect.ValueOf(i)
	fmt.Println("value:", v)

	// 得到类型的元数据,通过t我们能获取类型定义里面的所有元素
	t := reflect.TypeOf(i)
	fmt.Println("type:", t)

	// 转化为reflect对象之后我们就可以进行一些操作了,也就是将reflect对象转化成相应的值

	// Elem返回类型的元素类型
	controllerType := t.Elem()
	tag := controllerType.Field(2).Tag //Field(第几个字段,index从0开始)
	fmt.Println("Tag:", tag)

	method, _ := t.MethodByName("Init")
	fmt.Println(method)

	vMethod := v.MethodByName("Init")
	fmt.Println(vMethod)

	// 有输入参数的方法调用
	// 构造输入参数
	args1 := []reflect.Value{reflect.ValueOf("Run2"),reflect.ValueOf("POST")}
	// 通过v进行调用
	v.MethodByName("Init").Call(args1)

	// 无输入参数的方法调用
	// 构造zero value
	args2 := make([]reflect.Value, 0)
	// 通过v进行调用
	v.MethodByName("Test").Call(args2)

}
