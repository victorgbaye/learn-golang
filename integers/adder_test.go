package integers

import (
	"testing"
	"fmt"
	)

func TestAdder (t *testing.T){
	sum:= Add(2,2)
	expected:= 4

	if expected!= sum{
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
}