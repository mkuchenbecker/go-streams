package main

import "cmp"


func LessThan[T cmp.Ordered](t T, k T) bool {
	return t < k
}

func GreaterThan[T cmp.Ordered](t T, k T) bool {
	return t > k
}