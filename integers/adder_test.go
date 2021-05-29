package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	epected := 4

	if sum != epected {
		t.Errorf("expected '%d', but got '%d'", epected, sum)
	}
}

func ExampleAdd() {
	sum := Add(2, 3)
	fmt.Println(sum)
	// output: 5
}
