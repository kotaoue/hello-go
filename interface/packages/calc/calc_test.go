package calc_test

import (
	"testing"

	"github.com/kotaoue/go-interface/packages/calc"
)

func TestCalculator_Calc(t *testing.T) {
	c := &calc.Calculator{Processor: &calc.CALU{}}
	r := c.Calc(1, 2)
	e := 3
	if r != e {
		t.Fatalf("I was expecting %d, but return value was %d.", r, e)
	}
}

type faker struct{}

func (f *faker) Calc(a, b int) int {
	return a - b
}
func TestFaker_Calc(t *testing.T) {
	c := &calc.Calculator{Processor: &faker{}}
	r := c.Calc(1, 2)
	e := -1
	if r != e {
		t.Fatalf("I was expecting %d, but return value was %d.", r, e)
	}
}
