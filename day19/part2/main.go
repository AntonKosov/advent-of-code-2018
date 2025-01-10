package main

import "fmt"

func main() {
	// initialRegister0Value := 0
	initialRegister0Value := 1
	sum := 0
	count := 875
	if initialRegister0Value == 1 {
		count += (27*28 + 29) * 30 * 14 * 32
	}

	// The original code:
	// for i := 1; i <= count; i++ {
	// 	for j := 1; j <= count; j++ {
	// 		if i*j == count {
	// 			sum += i
	// 		}
	// 	}
	// }

	// Optimized code:
	for i := 1; i <= count; i++ {
		if (count % i) == 0 {
			sum += i
		}
	}

	fmt.Println("Answer: ", sum)
}

// func main() {
// 	r0, r1, r3, r4, r5 := 0, 0, 0, 0, 0
//
// 	// #ip 2
// 	// 00 addi 2 16 2 - jump +16
// 	goto line17
// line01:
// 	// 01 seti 1 1 1
// 	r1 = 1
// line02:
// 	// 02 seti 1 4 3
// 	r3 = 1
// line03:
// 	// 03 mulr 1 3 5
// 	r5 = r1 * r3
// 	// 04 eqrr 5 4 5
// 	r5 = cmp(r5 == r4)
// 	// 05 addr 5 2 2
// 	if r5 == 1 {
// 		goto line07
// 	}
// 	// 06 addi 2 1 2
// 	goto line08
// line07:
// 	// 07 addr 1 0 0
// 	r0 += r1
// line08:
// 	// 08 addi 3 1 3
// 	r3++
// 	// 09 gtrr 3 4 5
// 	r5 = cmp(r3 > r4)
// 	// 10 addr 2 5 2
// 	if r5 == 1 {
// 		goto line12
// 	}
// 	// 11 seti 2 4 2
// 	goto line03
// line12:
// 	// 12 addi 1 1 1
// 	r1++
// 	// 13 gtrr 1 4 5
// 	r5 = cmp(r1 > r4)
// 	// 14 addr 5 2 2
// 	if r5 == 1 {
// 		goto line16
// 	}
// 	// 15 seti 1 0 2
// 	goto line02
// 	// 16 mulr 2 2 2 // Exit
// line16:
// 	goto done
// line17:
// 	// 17 addi 4 2 4 - first line after the initial jump
// 	r4 += 2
// 	// 18 mulr 4 4 4
// 	r4 *= r4
// 	// 19 mulr 2 4 4
// 	r4 *= 19
// 	// 20 muli 4 11 4
// 	r4 *= 11
// 	// 21 addi 5 1 5
// 	r5++
// 	// 22 mulr 5 2 5
// 	r5 *= 22
// 	// 23 addi 5 17 5
// 	r5 += 17
// 	// 24 addr 4 5 4
// 	r4 += r5
// 	// 25 addr 2 0 2 // first or second part
// 	if r0 == 1 {
// 		goto line27
// 	}
// 	// 26 seti 0 9 2
// 	goto line01
// line27:
// 	// 27 setr 2 3 5
// 	r5 = 27
// 	// 28 mulr 5 2 5
// 	r5 *= 28
// 	// 29 addr 2 5 5
// 	r5 += 29
// 	// 30 mulr 2 5 5
// 	r5 *= 30
// 	// 31 muli 5 14 5
// 	r5 *= 14
// 	// 32 mulr 5 2 5
// 	r5 *= 32
// 	// 33 addr 4 5 4
// 	r4 += r5
// 	// 34 seti 0 9 0
// 	r0 = 0
// 	// 36 seti 0 6 2
// 	goto line01
//
// done:
// 	fmt.Println("Answer: ", r0)
// }

// func cmp(cond bool) int {
// 	if cond {
// 		return 1
// 	}
//
// 	return 0
// }
