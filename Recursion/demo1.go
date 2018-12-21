package main

import "fmt"

func Factorial(n uint64)(result uint64) {
	if (n > 0) {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func main() {
	var i int = 15

	//  利用返回15的阶乘,即1*2*3......*14*15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}