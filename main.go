package main

import (
	"fmt"
)

// return append(Map(proc, Tail(arr)), makeSlice(proc(Head(arr)))...)
// return Concat(proc(Head(arr)), Map(proc, Tail(arr)))
func main() {
	myarr := []int{1, 2, 3, 4}

	s := Map(AddOne, myarr)
	fmt.Println(s)
	ls := Filter(biggerThanTwo, myarr)
	fmt.Println(ls)
	mq := Reduce(addTogether, 0, myarr)
	fmt.Println(mq)
}

func addTogether(a, b int) int {
	return a + b
}

func AddOne(x int) int {
	return x + 1
}

func Head[T any](x []T) T {
	return x[0]
}

func Tail[A any](x []A) []A {
	return x[1:]
}

func Concat[T any](elem T, arr []T) []T {
	arr = append(arr, elem)
	return arr
}

func biggerThanTwo(elem int) bool {
	return elem > 2
}

func makeSlice[T any](elem T) []T {
	arrs := []T{elem}
	return arrs
}

// (a -> Bool) -> [a] -> [a]
func Filter[A any](predicate func(A) bool, arr []A) []A {
	if len(arr) == 0 {
		return nil
	} else if predicate(Head(arr)) == false {
		return Filter(predicate, Tail(arr))
	} else {
		return append(makeSlice(Head(arr)), Filter(predicate, Tail(arr))...)
	}
}

// Map :: (a -> b) -> [a] -> [b]
func Map[A, B any](proc func(A) B, arr []A) []B {
	if len(arr) == 0 {
		return nil
	} else {
		return append(makeSlice(proc(Head(arr))), Map(proc, Tail(arr))...)
	}
}

// Reduce :: (a -> b -> b) -> b -> [a] -> b
func Reduce[A, B any](op func(A, B) B, initial B, arr []A) B {
	if len(arr) == 0 {
		return initial
	} else {
		return op(Head(arr), Reduce(op, initial, Tail(arr)))
	}
}
