package complexStruct

/*
   -*- coding: utf-8 -*-
   @Author  : hakusai22
   @Time    : 2023/08/14 14:43
*/

import (
	"fmt"
)

//结构体是值类型
//如果结构体的字段类型是: 指针，slice，和map的零值都是 nil ，即还没有分配空间。如果需要使用这样的字段，需要先make，才能使用.

type Person struct {
	//字段大写表示公有，小写为私有
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int              //指针
	slice  []int             //切片
	map1   map[string]string //map
}

type Monster struct {
	Name string
	Age  int
}

func main() {

	//声明方式1：
	var p1 Person
	fmt.Println(p1)

	if (p1.ptr == nil) && (p1.slice == nil) && (p1.map1 == nil) {
		fmt.Println("指针，slice，和map都是 nil")
	}

	var a int = 10
	p1.ptr = &a

	//使用slice, 一定要make
	p1.slice = make([]int, 10)
	p1.slice[0] = 100

	//使用map, 一定要先make
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "tom~"

	fmt.Println(p1)

	//结构体是值类型
	var monster1 Monster
	monster1.Name = "牛魔王"
	monster1.Age = 500

	monster2 := monster1 //结构体是值类型，默认为值拷贝
	monster2.Name = "青牛精"

	fmt.Println("monster1=", monster1) //monster1= {牛魔王 500}
	fmt.Println("monster2=", monster2) //monster2= {青牛精 500}

}
