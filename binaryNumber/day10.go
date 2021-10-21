package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int64
	var max, count int
	fmt.Scan(&num)
	binaryNumbers := strconv.FormatInt(num, 2)
	for _, value := range binaryNumbers {
		binaryNumber := string(value)
		if binaryNumber == "1" {
			count++
			if count > max {
				max = count
			}
		} else {
			count = 0
		}
	}
	fmt.Println(max)
}
