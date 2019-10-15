package main

import (
	"fmt"
	"golang-book/binarySearch/newArr"
	"golang-book/binarySearch/out"
)
import "golang-book/binarySearch/search_dr"

func main () {
	amountOfElements := 100
	Arr := newArr.NewArr(amountOfElements)

	desiredNum := 5
	element, err := search_dr.Search(desiredNum, Arr)

	fmt.Println(Arr)
	out.Output(element, err)
}