package main

import "fmt"

type seller interface {
	cancelInformation(info string) string
}
type seller1 struct{}

func (s seller1) cancelInformation(info string) string {
	return fmt.Sprintf("this is cancel information of seller1 %s \n", info)
}

type seller2 struct{}

func (s seller2) cancelInformation(info string) string {
	return fmt.Sprintf("this is cancel information of seller2 %s \n", info)
}

type reservation struct {
	s seller
}
type premierReservation struct {
	reservation
}

func (r *reservation) cancelReservation() {
	fmt.Println(r.s.cancelInformation(""))
}
func (pr *premierReservation) cancelReservation() {
	fmt.Println(pr.s.cancelInformation(""))
}
func main() {
	r := reservation{s: seller1{}}
	r.cancelReservation()

	r2 := reservation{seller2{}}
	pr := premierReservation{r2}
	pr.cancelReservation()
}
