package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// {1, 2, 2, 4}

func countTriplets(arr []int64, r int64) int64 {
	bMap := make(map[int64]int64)
	cMap := make(map[int64]int64)
	var count int64

	for _, value := range arr {
		// check for triplet
		if _, ok := bMap[value]; ok {
			count += bMap[value]
		}
		// check for double
		if _, ok := cMap[value]; ok {
			bMap[value*r] += cMap[value]
		}
		// always add to cMap
		cMap[value*r]++
	}

	return count
}

// test-case-2 output should be:  161700
// test-case-2 output should be:  166661666700000
func main() {
	file, err := os.Open("test-case-2")
	checkError(err)

	reader := bufio.NewReaderSize(file, 16*1024*1024)

	stdout, err := os.Create("OUTPUT-test-case-2")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(nr[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	r, err := strconv.ParseInt(nr[1], 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int64

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arr = append(arr, arrItem)
	}

	ans := countTriplets(arr, r)

	fmt.Fprintf(writer, "%d\n", ans)

	writer.Flush()
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
