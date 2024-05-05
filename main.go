package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer func(out *bufio.Writer) {
		_ = out.Flush()
	}(out)
}

//////////////////////////////////////////////////
// CODE TEMPLATE FOR SOLVING PROBLEMS
// SOME FUNCTIONS MIGHT NOT BE USED
//////////////////////////////////////////////////

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func readInt(in *bufio.Reader) int {
	l, _ := strconv.Atoi(readLine(in))
	return l
}

func readLine(in *bufio.Reader) string {
	l, _ := in.ReadString('\n')
	return strings.TrimSpace(l)
}

func readArrString(in *bufio.Reader) []string {
	return strings.Split(readLine(in), " ")
}

func readArrInt(in *bufio.Reader) []int {
	r := readArrString(in)
	arr := make([]int, len(r))
	for i := 0; i < len(arr); i++ {
		arr[i], _ = strconv.Atoi(r[i])
	}
	return arr
}

func readArrInt64(in *bufio.Reader) []int64 {
	r := readArrString(in)
	arr := make([]int64, len(r))
	for i := 0; i < len(arr); i++ {
		arr[i], _ = strconv.ParseInt(r[i], 10, 64)
	}
	return arr
}
