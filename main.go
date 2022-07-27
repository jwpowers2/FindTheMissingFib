package main

import (
	"fmt"
)

func main(){
	f := NewFibGame()
	f.Fibs = f.MakeFib(21)
	fmt.Println(f.Fibs)
	fmt.Println(f.FibMap)
	f.RemoveNumber(2)
	fmt.Println(f.Fibs)
	f.FindMissingFib(f.Fibs)
}