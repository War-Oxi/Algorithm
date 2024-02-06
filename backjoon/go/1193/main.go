// https://www.acmicpc.net/problem/1193 [분수찾기]
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)

	var numerator, denominator int // 분자, 분모
	
	var tmpNum int = 1
	
	numerator = 1
	denominator = tmpNum
	
	for i := 1; i <= N; i++ {
		if numerator == tmpNum {
			tmpNum++
			numerator = 1
			denominator = tmpNum
			continue
		}
		numerator++
		denominator--
	}
}