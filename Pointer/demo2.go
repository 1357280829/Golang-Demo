package main

import "fmt"

func main() {
	var  ptr *int

	/* 空指针 */
	fmt.Printf("ptr 的值为 : %x\n", ptr  ) //  0
	fmt.Println(ptr) //  <nil>

	/* 空指针判断 */
	//if(ptr != nil)     /* ptr 不是空指针 */
	//if(ptr == nil)    /* ptr 是空指针 */
}