package main

import (
	"log"
	"reflect"
)

//Foo 目标接口定义
type Foo interface {
	Bar(int)
}

//CheckFoo 检查是否是Foo接口实现
func CheckFoo(src interface{}) bool {
	dst := (*Foo)(nil)
	dstType := reflect.TypeOf(dst).Elem()

	srcType := reflect.TypeOf(src)
	return srcType.Implements(dstType)
}

func main() {
	if CheckFoo(1) {
		log.Println("Foo OK. ")
	} else {
		log.Println("Foo No. ")
	}

}
