// taken from https://stackoverflow.com/questions/30904718/range-of-immediate-values-in-armv8-a64-assembly/33265035#33265035

package main

import (
	"fmt"
)

func main() {
	// print_immediate32()
	print_immediate64()
}

// Dumps all legal bitmask immediates for ARM64
// Total number of unique 64-bit patterns:
//   1*2 + 3*4 + 7*8 + 15*16 + 31*32 + 63*64 = 5334

func print_immediate64() {
	var result uint64
	var size, length, e, rotation uint8
	for size = 2; size <= 64; size *= 2 {
		for length = 1; length < size; length++ {
			result = 0xffffffffffffffff >> (64 - length)
			for e = size; e < 64; e *= 2 {
				result |= result << e
			}
			for rotation = 0; rotation < size; rotation++ {
				fmt.Printf("0x%016x %s (size=%v, length=%v, rotation=%v)\n",
					result, uint64_to_binary(result),
					size, length, rotation)
				break
				result = (result >> 1) | (result << 63)
			}
		}
	}
}

func uint64_to_binary(x uint64) [64]uint8 {
	var b [64]uint8

	for i := 63; i >= 0; i-- {
		if x&1 != 0 {
			b[i] = '1'
		} else {
			b[i] = '0'
		}
		x >>= 1
	}
	return b
}

func print_immediate32() {
	var result uint32
	var size, length, e, rotation uint8
	for size = 2; size <= 32; size *= 2 {
		for length = 1; length < size; length++ {
			result = 0xffffffff >> (32 - length)
			for e = size; e < 32; e *= 2 {
				result |= result << e
			}
			for rotation = 0; rotation < size; rotation++ {
				fmt.Printf("0x%08x %s (size=%v, length=%v, rotation=%v)\n",
					result, uint32_to_binary(result),
					size, length, rotation)
				result = (result >> 1) | (result << 31)
			}
		}
	}
}

func uint32_to_binary(x uint32) [32]uint8 {
	var b [32]uint8

	for i := 31; i >= 0; i-- {
		if x&1 != 0 {
			b[i] = '1'
		} else {
			b[i] = '0'
		}
		x >>= 1
	}
	return b
}
