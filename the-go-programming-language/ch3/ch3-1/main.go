package main

import (
	"fmt"
	"reflect"
)

func main() {
	p59()
	p60()
	p61_1()
	p61_2()
	p61_3()
	p62_1()
	p62_2()
}

func p59() {
	fmt.Println("--P59--")
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	var i int8 = 127
	fmt.Println(1, 1+1, i*i)
}

func p60() {
	fmt.Println("--P60--")
	var x uint8 = 1 << 1
	fmt.Printf("%08b\n", x)
	x |= 1 << 5
	fmt.Printf("%08b\n", x)

	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)
	fmt.Printf("%08b\n", ^(x | y))

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)
}

func p61_1() {
	fmt.Println("--P61_1--")
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}
}

func p61_2() {
	fmt.Println("--P61_2--")
	var apples int32 = 1
	var oranges int16 = 2
	var compote = int(apples) + int(oranges)

	fmt.Printf("%d (%s)\n", apples, reflect.TypeOf(apples))
	fmt.Printf("%d (%s)\n", oranges, reflect.TypeOf(oranges))
	fmt.Printf("%d (%s)\n", compote, reflect.TypeOf(compote))

	var compote2 = int16(apples) + oranges
	fmt.Printf("%d (%s)\n", compote2, reflect.TypeOf(compote2))
}

func p61_3() {
	fmt.Println("--P61_3--")
	f := 3.141
	i := int(f)

	fmt.Println(f, i)
	f = 1.99
	fmt.Println(i)
}

func p62_1() {
	fmt.Println("--P62_1--")
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)

	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
}

func p62_2() {
	fmt.Println("--P62_2--")
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)
}
