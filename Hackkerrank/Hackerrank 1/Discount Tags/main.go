package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'getMaximumEvenSum' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts INTEGER_ARRAY val as parameter.
 */

func getMaximumEvenSum(val []int32) int64 {
	var total int64 = 0
	for i := 0; i < len(val); i++ {
		total += int64(val[i])
	}

	if total%2 == 0 {
		return total
	}

	lastOdd := 0

	for i := 0; i < len(val); i++ {
		if val[i]%2 == 1 && (lastOdd == 0 || int(val[i]) < lastOdd) {
			total += int64(lastOdd)
			total -= int64(val[i])
			lastOdd = int(val[i])
		}
	}

	return total
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	valCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var val []int32

	for i := 0; i < int(valCount); i++ {
		valItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		valItem := int32(valItemTemp)
		val = append(val, valItem)
	}

	result := getMaximumEvenSum(val)

	fmt.Fprintf(writer, "%d\n", result)

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
