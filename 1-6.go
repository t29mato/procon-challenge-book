package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := nextsInt(n)

	// 降順でソートする
	// https://xn--go-hh0g6u.com/pkg/sort/#Reverse
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	result := 0

	for i := 2; i < n; i++ {
		if a[i-2] < a[i-1]+a[i] {
			result = a[i-2] + a[i-1] + a[i]
			break
		}
	}
	fmt.Println(result)
}

// I/O
var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	result, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return result
}

func nextsInt(n int) []int {
	sc.Scan()
	var result []int
	for i := 0; i < n; i++ {
		tmp, err := strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
		result = append(result, tmp)
		sc.Scan()
	}
	return result
}
