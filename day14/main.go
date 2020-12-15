package main

import (
	"fmt"
	"bufio"
	"os"
)

const IntSize = 36 // bits
const MaskAll1 = (1 << IntSize) - 1
type Mask struct {
	mask0, mask1 int
}

func parseMask(str string) (mask0, mask1 int) {
	mask0 = MaskAll1
	for i := 0; i < IntSize; i++ {
		curBit := str[IntSize - i - 1]
		if curBit == '1' {
			mask1 |= 1 << i
		} else if curBit == '0' {
			mask0 = mask0 &^ (1 << i)
		}
	}
	fmt.Printf("(x & %b) | %b\n", mask0, mask1)
	return
}

func parseMaskFloating(str string) (masks []Mask) {
	var floatingPlaces []int
	mask0 := MaskAll1
	mask1 := 0
	for i := 0; i < IntSize; i++ {
		curBit := str[IntSize - i - 1]
		if curBit == '1' {
			mask1 |= 1 << i
		} else if curBit == 'X' {
			mask0 = mask0 &^ (1 << i)
			floatingPlaces = append(floatingPlaces, i)
		}
	}
	fmt.Printf("(x & %b) | %b\n", mask0, mask1)
	fmt.Println(floatingPlaces)

	masks = []Mask{Mask{mask0, mask1}}
	for _, place := range floatingPlaces {
		numMasks := len(masks)
		for i := 0; i < numMasks; i++ {
			mask := masks[i]
			newMask1 := mask.mask1 | 1 << place
			newMask := Mask{mask.mask0, newMask1}
			masks = append(masks, newMask)
		}
	}

	return
}

func sumMap(m map[int]int) (sum int) {
	for _, value := range m {
		sum += value
	}
	return
}

func main() {
	mem := make(map[int]int)
	var mask0, mask1 int

	mem2 := make(map[int]int)
	var masks []Mask

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line[1] == 'a' { // mask
			var maskStr string
			fmt.Sscanf(line, "mask = %s", &maskStr)
			mask0, mask1 = parseMask(maskStr)

			masks = parseMaskFloating(maskStr)
		} else { // mem
			var address, value int
			fmt.Sscanf(line, "mem[%d] = %d", &address, &value)
			maskedValue := (value & mask0) | mask1
			mem[address] = maskedValue

			for _, mask := range masks {
				maskedAddress := (address & mask.mask0) | mask.mask1
				mem2[maskedAddress] = value
				// fmt.Printf("mem2[%b] = %d\n", maskedAddress, value)
			}
		}
	}

	sum := sumMap(mem)
	fmt.Println("sum 1", sum)

	sum2 := sumMap(mem2)
	fmt.Println("sum 2", sum2)
}

