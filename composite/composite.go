package main

import "fmt"

type human interface {
	eating()
	addHuman(thing human)
}

type humanBeing struct {
	humanList []human
}
type man struct {
}

func (m man) eating() {
	fmt.Println("this is a man eating")
}

func (m man) addHuman(thing human) {
	panic("implement me")
}

func (h *humanBeing) eating() {
	if len(h.humanList) == 0 {
		fmt.Println("no human")
		return
	}
	fmt.Println("I am a composite")
	for _, val := range h.humanList {
		val.eating()
	}
}

func (h *humanBeing) addHuman(thing human) {
	h.humanList = append(h.humanList, thing)
}

func main() {
	var hb human
	hb = &humanBeing{}
	hb.eating()
	child := &humanBeing{}
	hb.addHuman(child)

	m := &man{}
	hb.addHuman(m)
	hb.eating()
}
