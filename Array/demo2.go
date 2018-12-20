

//  杨辉三角实现
//[1]
//[1 1]
//[1 2 1]
//[1 3 3 1]
//[1 4 6 4 1]
//[1 5 10 10 5 1]
//[1 6 15 20 15 6 1]
//[1 7 21 35 35 21 7 1]
//[1 8 28 56 70 56 28 8 1]
//[1 9 36 84 126 126 84 36 9 1]

package main
import "fmt"
func GetYangHuiTriangleNextLine(arr []int) []int {
	var tempArr []int
	var i int
	arrLen := len(arr)
	tempArr = append(tempArr, 1)
	if 0 == arrLen {
		return tempArr
	}
	for i = 0; i < arrLen-1; i++ {
		tempArr = append(tempArr, arr[i]+arr[i+1])
	}
	tempArr = append(tempArr, 1)
	return tempArr
}

func main() {

	//  下面两句相同
	//arr := []int {}
	var arr []int

	var i int
	for i = 0; i < 10; i++ {
		arr = GetYangHuiTriangleNextLine(arr)
		fmt.Println(arr)
	}
}
