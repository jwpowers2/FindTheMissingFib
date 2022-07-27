package main

import (
	"fmt"
	"math"
)

type FibGame struct {
	Fibs []int
	FibMap map[int]int
}

func NewFibGame()*FibGame{
	return &FibGame{
		Fibs: []int{},
		FibMap: map[int]int{},
	}
}

func (f *FibGame) MakeFib(num int) []int{

	count := 0
	fibSlice := []int{0}
    f.FibMap[0] = 0
	for fibSlice[count] < num{
		
		var newInt int
		if fibSlice[count] == 0 {
			newInt = fibSlice[count] + 1
		} else {
		    newInt = fibSlice[count] + fibSlice[count-1]
		}
		fibSlice = append(fibSlice, newInt) 
		f.FibMap[newInt] = count
		count += 1
		
	}
	return fibSlice
}

func (f *FibGame) RemoveNumber(val int){

	index := f.FibMap[val]
	left := index + 1
	right := index + 2
    f.Fibs = append(f.Fibs[:left], f.Fibs[right:]...)
}

func (f *FibGame) FindMissingFib(arr []int) {
	
	if len(arr) > 2 {
		
		var length float64
		length = float64(len(arr)) / float64(2)
		pivot64 := math.Ceil(length)
		pivot := int(pivot64)
		left := append(arr[:pivot])
		right := append(arr[pivot:])
		leftFibSlice := f.MakeFib(arr[pivot-1])
		totalFibSlice := f.MakeFib(arr[len(arr)-1])
		rightFibSlice := append(totalFibSlice[pivot:])
		
		if len(leftFibSlice) != len(left){
			fmt.Println(left)
			f.FindMissingFib(left)

		} else if len(rightFibSlice) != len(right){
			fmt.Println(right)
			f.FindMissingFib(right)
		}
	} else {
		fmt.Println(arr)
	}
}
