package util

import "fmt"

// CodeToAEHA - IR Code to AEHA Hex Code
func CodeToAEHA(entries [][]int) {
	offset := -1
	byt := 0

	for i := range entries {
		on := entries[i][0]
		off := entries[i][1]
		if off == 0 {
			break
		}

		if on > 1500 {
			if offset >= 0 {
				fmt.Printf("\n")
			}
			offset = 0
			byt = 0
		} else {
			r := 0
			if off > on*2 {
				r = 1
			}
			byt |= (r) << offset
			offset = offset + 1

			if (offset & 7) == 0 {
				fmt.Printf("%02X ", byt)
				offset = 0
				byt = 0
			}
		}
	}
}

// SignalToCode - Hex code to IR Code
func SignalToCode(T int, signal [][]int, interval int) []int {
	var code []int
	for i := range signal {
		c := signal[i]
		code = append(code, T*8, T*4)

		for j := range c {
			for k := 0; k < 8; k++ {
				code = append(code, T)
				if c[j]&(1<<k) != 0 {
					code = append(code, T*3)
				} else {
					code = append(code, T)
				}
			}
		}

		code = append(code, T)
		if i < len(signal)-1 {
			code = append(code, interval)
		}
	}

	return code
}
