package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != 6 {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}
	var counter int
	max := math.Inf(-1)
	row, col := len(arr)-2, len(arr)-2
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			counter = 0
			counter += int(arr[i][j])
			counter += int(arr[i][j+1])
			counter += int(arr[i][j+2])
			counter += int(arr[i+1][j+1])
			counter += int(arr[i+2][j])
			counter += int(arr[i+2][j+1])
			counter += int(arr[i+2][j+2])

			if counter > int(max) {
				max = float64(counter)
			}
		}
	}
	fmt.Println(max)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
