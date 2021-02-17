package main

import (
	"fmt"
  "bufio"
  "os"
	"strconv"
)

func main() {
  input := bufio.NewScanner(os.Stdin)
	input.Scan()
	i, _ := strconv.ParseInt(input.Text(), 10, 64)
	fmt.Printf("the facorial of %d is %d", i, factorial(i))
}

func factorial(n int64) int64 {
	if n == 0 || n == 1 {
		return 1
	} else {
		return factorial(n-1) * n
	}
}
