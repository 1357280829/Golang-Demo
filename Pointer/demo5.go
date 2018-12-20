
/* Go指针作为函数参数 */

package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int= 200

	fmt.Printf("交换前 a 的值 : %d\n", a )
	fmt.Printf("交换前 b 的值 : %d\n", b )

	/* 调用函数用于交换值
	* &a 指向 a 变量的地址
	* &b 指向 b 变量的地址
	*/
	swap(&a, &b);
	//swap2(&a, &b);

	fmt.Printf("交换后 a 的值 : %d\n", a )
	fmt.Printf("交换后 b 的值 : %d\n", b )
}

func swap(x *int, y *int) {
	var temp int
	temp = *x    /* 保存 x 指向地址的值 */
	*x = *y      /* 将 y 指向地址的值赋值给 x 指向地址的值 */
	*y = temp    /* 将 temp 赋值给 y 指向地址的值 */
}


/* 交换函数这样写更加简洁，也是 go 语言的特性，可以用下，c++ 和 c# 是不能这么干的 */
//func swap2(x *int, y *int){
//	*x, *y = *y, *x
//}