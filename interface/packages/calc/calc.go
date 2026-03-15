package calc

// ALU is the interface that wraps the basic calculate method.
type ALU interface {
	Calc(a, b int) int
}

// Calculator is the struct for execute basic calculate.
type Calculator struct {
	Processor ALU
}

// Calc is calculated two numbers.
func (c *Calculator) Calc(a, b int) int {
	return c.Processor.Calc(a, b)
}

// CALU is ALU for Calculator struct.
type CALU struct{}

// Calc is adds up two numbers.
func (c *CALU) Calc(a, b int) int {
	return a + b
}

// NewCalculator returns a Calculator that calculate inputed numbers.
func NewCalculator(a ALU) *Calculator {
	return &Calculator{Processor: a}
}
