package main

import (
	"fmt"
	"testing"
)

func TestCountTriplets1(t *testing.T) {
	arr := []int64{1, 2, 2, 4}
	r := int64(2)

	triplets := countTriplets(arr, r)

	if triplets != 2 {
		t.Errorf("got %d instead of 2", triplets)
	}
}

func TestCountTriplets2(t *testing.T) {
	fmt.Println("")
	arr := []int64{1, 3, 9, 9, 27, 81}
	r := int64(3)

	triplets := countTriplets(arr, r)

	if triplets != 6 {
		t.Errorf("got %d instead of 6", triplets)
	}
}

func TestCountTriplets3(t *testing.T) {
	fmt.Println("")
	arr := []int64{1, 5, 5, 25, 125}
	r := int64(5)

	triplets := countTriplets(arr, r)

	if triplets != 4 {
		t.Errorf("got %d instead of 4", triplets)
	}
}
