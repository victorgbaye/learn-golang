package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collectionn of 5 numbers", func(t *testing.T){
		var numbers = []int{1,2,3,4,5}
		got := Sum(numbers)
		want := 15
	
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	t.Run("collection of any size", func(t *testing.T) {
		var number = []int{1,2,3}
		got := Sum(number)
		want := 6

		if got != want {
			t.Errorf("got %d want want %d given, %v", got, want, number)
		}

	})

	t.Run("sum all slices",func(t *testing.T) {
		got := SumAll([]int{1, 4}, []int{2,3})
		want := []int{3, 7}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}