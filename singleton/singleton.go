package main

import (
	"fmt"
	"sync"
)

type myClass struct {
	age int
}
var (
	once     sync.Once
	instance *myClass
)
func getMyClass() *myClass {
	once.Do(func() {
		instance = &myClass{age: 32}
	})
	return instance
}

func (mc *myClass) setAge(age int){
	mc.age = age
}
func (mc *myClass)getAge() int{
	return mc.age
}
/*
It should be noted that the singleton pattern is actually considered an anti-pattern because of the introduction of a global state.
This causes a hidden coupling between components and can lead to difficult-to-debug situations.
It should not be overused.
 */
func main() {
	class := getMyClass()
	fmt.Println(class.getAge())
	class.setAge(18)
	class2 := getMyClass()
	fmt.Println(class2.getAge())
}
