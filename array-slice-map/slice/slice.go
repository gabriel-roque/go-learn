package main

import (
	"fmt"
	"reflect"
)

func main() {
	array := [3]int{1, 2, 3} // array
	slice := []int{1, 2, 3}  // slice

	fmt.Println(array, slice)
	fmt.Println(reflect.TypeOf(array), reflect.TypeOf(slice))

	// Slice não é um array! Slice define um pedaço de um array!
	array2 := [5]int{1, 2, 3, 4, 5}
	slice2 := array[1:3]
	fmt.Println(array2, slice2)

	slice3 := array2[:2]
	fmt.Println(array2, slice3)
}
