package main

import "fmt"

type abstractFactory interface {
	createInvoice() string
	createReservation() string
}

type hotelFactory struct{}

type flightFactory struct{}

func (f flightFactory) createInvoice() string {
	return "create invoice for flight"
}

func (f flightFactory) createReservation() string {
	return "create reservation for flight"
}

func (h hotelFactory) createInvoice() string {
	return "create invoice for hotel"
}

func (h hotelFactory) createReservation() string {
	return "create reservation for hotel"
}

func getFactory(t string) abstractFactory {
	switch t {
	case "hotel":
		return hotelFactory{}
	case "flight":
		return flightFactory{}
	default:
		return nil
	}
}

func main() {
	factory := getFactory("hotel")
	fmt.Println(factory.createInvoice())
	fmt.Println(factory.createReservation())
}
