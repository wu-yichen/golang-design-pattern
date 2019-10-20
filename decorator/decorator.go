package main

import (
	"fmt"
	"time"
)

type test func(num int) int
func decoratorFunc(fn test) test {
	return func(num int) int {
		startTime := time.Now()
		result := fn(num)
		endTime := time.Now()
		fmt.Printf("the total running time is %s \n", endTime.Sub(startTime))
		return result
	}
}
func sum(num int) int {
	result := 0
	for i := 0; i < num; i++ {
		result += i
	}
	return result
}
func main() {
	t := decoratorFunc(sum)
	fmt.Println(t(1666))
}
