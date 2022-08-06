package integers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdder(t *testing.T) {
    sum := Add(2,2)
    expected := 4

    if sum != expected {
        t.Errorf("expected '%d' but got '%d'", expected, sum)
    }
}

func ExampleAdd() {
    sum := Add(1,5)
    fmt.Println(sum)
    // Output: 6
}

func TestSumAll(t *testing.T) {
    got := SumAll([]int{1, 2}, []int{0, 9})
    want := []int{3, 9}

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}
