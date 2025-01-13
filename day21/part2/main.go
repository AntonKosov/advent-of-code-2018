package main

import "fmt"

func main() {
	nums := map[int]bool{}
	lastNum := -1
	r3 := 0
	for {
		r4 := r3 | 0x10000
		r3 = 7041048
		for ; r4 > 0; r4 /= 256 {
			r3 += r4 & 0xff
			r3 &= 0xffffff
			r3 *= 65899
			r3 &= 0xffffff
		}
		if nums[r3] {
			break
		}
		nums[lastNum] = true
		lastNum = r3
	}

	fmt.Println("Answer:", lastNum)
}
