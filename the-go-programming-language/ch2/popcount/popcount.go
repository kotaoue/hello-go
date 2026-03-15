package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var b byte
	for i := 0; i < +8; i++ {
		b += pc[byte(x>>(i*8))]
	}
	return int(b)
}

func PopCountBitShift(x uint64) int {
	var c int
	for i := 0; i < 64; i++ {
		c += int(x >> (i) & 1)
	}
	return c
}

func PopCountBitClear(x uint64) int {
	var c int
	for x > 0 {
		x &= (x - uint64(1))
		c++
	}
	return c
}
