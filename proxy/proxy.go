package main

import "fmt"

type proxyClass struct {
	mc *myClass
}
func (pc *proxyClass) doing(){
	fmt.Println("I can doing something else here perhaps an API call")
	pc.mc.doing()
}

type myClass struct {}
func (m *myClass) doing(){
	fmt.Println("myclass is doing something")
}
func main(){
	pc := &proxyClass{}
	pc.doing()
}
