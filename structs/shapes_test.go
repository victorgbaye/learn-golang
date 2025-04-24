package main

import (
	"reflect"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("testing perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0
	
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("area of rectangle", func(t *testing.T){
		rectangle := Rectangle{12.0, 6.0}
		got := Area(rectangle)
		want := 72.0

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	
}