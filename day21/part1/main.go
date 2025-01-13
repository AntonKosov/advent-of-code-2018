package main

import "fmt"

func main() {
	r3 := 7041048
	for r4 := 256 * 256; r4 > 0; r4 /= 256 {
		r3 += r4 & 0xff
		r3 &= 0xffffff
		r3 *= 65899
		r3 &= 0xffffff
	}
	fmt.Println("Answer:", r3)
}

// func main() {
// 	for initialR0 := 0; ; initialR0++ {
// 		// #ip 2
// 		r0, r1, r3, r4, r5 := initialR0, 0, 0, 0, 0
// 		// 00 seti 123 0 3
// 		r3 = 123
// 	line01:
// 		// 01 bani 3 456 3
// 		r3 &= 456
// 		// 02 eqri 3 72 3
// 		r3 = cmp(r3 == 72)
// 		// 03 addr 3 2 2
// 		if r3 == 1 {
// 			goto line05
// 		}
// 		// 04 seti 0 0 2
// 		goto line01
// 	line05:
// 		// 05 seti 0 6 3
// 		r3 = 0
// 	line06:
// 		// 06 bori 3 65536 4
// 		r4 = r3 | 65536
// 		// 07 seti 7041048 8 3
// 		r3 = 7041048
// 	line08:
// 		// 08 bani 4 255 5
// 		r5 = r4 & 255
// 		// 09 addr 3 5 3
// 		r3 += r5
// 		// 10 bani 3 16777215 3
// 		r3 &= 16777215
// 		// 11 muli 3 65899 3
// 		r3 *= 65899
// 		// 12 bani 3 16777215 3
// 		r3 &= 16777215
// 		// 13 gtir 256 4 5
// 		r5 = cmp(256 > r4)
// 		// 14 addr 5 2 2
// 		if r5 == 1 {
// 			goto line16
// 		}
// 		// 15 addi 2 1 2
// 		goto line17
// 	line16:
// 		// 16 seti 27 6 2
// 		goto line28
// 	line17:
// 		// 17 seti 0 1 5
// 		r5 = 0
// 	line18:
// 		// 18 addi 5 1 1
// 		r1 = r5 + 1
// 		// 19 muli 1 256 1
// 		r1 *= 256
// 		// 20 gtrr 1 4 1
// 		r1 = cmp(r1 > r4)
// 		// 21 addr 1 2 2
// 		if r1 == 1 {
// 			goto line23
// 		}
// 		// 22 addi 2 1 2
// 		goto line24
// 	line23:
// 		// 23 seti 25 1 2
// 		goto line26
// 	line24:
// 		// 24 addi 5 1 5
// 		r5++
// 		// 25 seti 17 8 2
// 		goto line18
// 	line26:
// 		// 26 setr 5 2 4
// 		r4 = r5
// 		// 27 seti 7 9 2
// 		goto line08
// 	line28:
// 		// 28 eqrr 3 0 5
// 		r5 = cmp(r3 == r0)
// 		// 29 addr 5 2 2
// 		if r5 == 1 {
// 			fmt.Println("Answer:", initialR0)
// 			return
// 		}
// 		// 30 seti 5 3 2
// 		goto line06
// 	}
// }
