package unittest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := Square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d, expected is %d, actual is %d", inputs[i], expected[i], ret)
		}
	}
}

func TestErrorInCode(t *testing.T) {
	fmt.Println("start")
	t.Error("Error")
	fmt.Println("end")
}

func TestFatalInCode(t *testing.T) {
	fmt.Println("start")
	t.Fatal("Fatal")
	fmt.Println("end")
}

func TestSquareWithAssert(t *testing.T) {
	inputs := [...]int{1, 3, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := Square(inputs[i])
		if ret != expected[i] {
			assert.Equal(t, expected[i], ret)
		}
	}
}
