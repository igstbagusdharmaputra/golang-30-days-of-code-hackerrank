package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var size int
	data := make(map[string]int)

	fmt.Scan(&size)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 1; i <= size; i++ {
		for scanner.Scan() {
			fullData := scanner.Text()
			extractData := strings.Split(fullData, " ")
			name := extractData[0]
			phoneNumber, _ := strconv.Atoi(extractData[1])
			data[name] = phoneNumber
			break
		}

	}
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		input := scanner.Text()
		_, check := data[input]
		if check {
			fmt.Printf("%v=%v\n", input, data[input])
		} else {
			fmt.Println("Not found")
		}
	}

}
