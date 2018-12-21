

package main

import "fmt"

func main() {
	/* 1.创建集合 */
	var countryCapitalMap map[string]string
	/* 2.初始化集合 */
	countryCapitalMap = make(map[string]string) //  如果不初始化,countryCapitalMap的值为nil map,此时不能存放键值对
	//  等价于上面两条
  	//countryCapitalMap := map[string]string{}

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/* (疑问?) 查看元素在集合中是否存在;如果确定是真实的,则存在,否则不存在 */
	captial, ok := countryCapitalMap [ "新德里" ]
	fmt.Println(captial)
	fmt.Println(ok)
	if (ok) {
		fmt.Println("新德里的首都是", captial)
	} else {
		fmt.Println("新德里的首都不存在")
	}
}