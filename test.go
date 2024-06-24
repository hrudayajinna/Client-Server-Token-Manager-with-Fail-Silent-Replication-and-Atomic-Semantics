package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

func main() {

	index, hashVa := hasher(2, 7)
	fmt.Println(index)
	fmt.Println(hashVa)

}

func hasher(start int, end int) (int, uint64) {
	hasher := sha256.New()
	var min uint64
	var temp uint64
	var index int
	for i := start; i < end; i++ {
		hasher.Write([]byte(fmt.Sprintf("%s %d", "karthik", i)))
		temp = binary.BigEndian.Uint64(hasher.Sum(nil))
		if i == start {
			min = temp
			index = 0
		}
		if i != start && min > temp {
			min = temp
			index = i
		}
	}

	return index, min
}
