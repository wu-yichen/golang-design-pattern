package main

import "fmt"

type adaptee struct{}

func (ade adaptee) doSomething() {
	fmt.Println("do something here from adaptee")
}

type adaptor struct {
	ade *adaptee
}

func (adt *adaptor) doSomethingElse() {
	fmt.Println("do something else from adaptor")
	adt.ade.doSomething()
}

func getAdaptor() *adaptor {
	return &adaptor{ade: &adaptee{}}
}

func main(){
	i := getAdaptor()
	i.doSomethingElse()
}
