package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T){
	got := Add(2,2)
	want := 4

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}


// Example functions help with accessibility of your cod by creating examples
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}