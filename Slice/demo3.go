
/* Go切片误区 */

package main

import "fmt"

//  错误示范(要求输出 {5,6,4,7,8,9} )
//func main() {
//
//	var a []int = []int{5,6,7,8,9}
//
//	first := a[:2]
//	fmt.Println("first :",first)
//
//	last := a[2:]
//	fmt.Println("last :", last)
//
//	new_first := append(first, 4) //  这时a变成了 [5,6,4,8,9] ;此时的append为修改
//	fmt.Println("new_first :",new_first)
//
//	new := append(new_first,last...) //  这时的new是 [5,6,4,4,8,9] ;此时的append为创建
//	fmt.Println("new :", new)
//}

//  正确示范(要求输出 {5,6,4,7,8,9} )
func main() {

	var a []int = []int{5,6,7,8,9}

	tempB := a[:2]
	b := make([]int, len(tempB), cap(a)*2)
	copy(b, tempB)
	printSlice(b)

	b = append(b, 4)
	printSlice(b)

	tempC := a[2:]
	c := make([]int, len(tempC), cap(a)*2)
	copy(c, tempC)
	printSlice(c)

	result := make([]int, 0, cap(a)*2)
	result = append(b, c...)
	printSlice(result)

}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}


