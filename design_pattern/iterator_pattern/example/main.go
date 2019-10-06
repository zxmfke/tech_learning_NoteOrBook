package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("===========slice iterator===========")
	sliceIterator := NewSliceIterator([]interface{}{"a", "b", "c"})
	for sliceIterator.HasNext() {
		element, err := sliceIterator.Next()
		if err != nil {
			fmt.Println("out of range\r\n")
			return
		}

		fmt.Printf("element : %v | type : %v\r\n", element, reflect.TypeOf(element))
	}

	fmt.Println("===========list iterator===========")
	list := NewList([]interface{}{1,2,3,4,5,6,7})
	listIterator,err := NewListIterator(list)
	if err != nil {
		fmt.Println("list nil\r\n")
		return
	}
	for listIterator.HasNext() {
		element, err := listIterator.Next()
		if err != nil {
			fmt.Println("out of range\r\n")
			return
		}

		fmt.Printf("element : %v | type : %v\r\n", element, reflect.TypeOf(element))
	}
}
