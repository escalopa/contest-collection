package main

import (
	"bufio"
	"math"
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

	c := solve()
	out.WriteString(strconv.Itoa(c) + "\n")
}

func solve() int {
	a := readInt(in)
	b := readInt(in)

	// b / c = X
	// c / a = Y
	// b / a = X * Y
	XY := float64(b) / float64(a)
	if XY != math.RoundToEven(XY) || int(XY)%2 != 0 {
		return -1
	}

	return int(XY) / 2
}

//////////////////////////////////////////////////
// CODE TEMPLATE FOR SOLVING PROBLEMS
// SOME FUNCTIONS MIGHT NOT BE USED
//////////////////////////////////////////////////

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
