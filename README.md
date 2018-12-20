# Golang-Demo

## 这是一个用于初步学习Golang的Demo仓库

### 注意点:

##### 1.Go中函数与方法要区分开(Go中 结构体 可理解为PHP中 类 )

##### 2.想要获取指针变量ptr指向的值,用*ptr获得;而想要利用结构体指针structPtr获取结构体成员,直接用structPtr.number即可

##### 3.(重点)数组Array和切片Slice

**(1)数组的定义方式**
```
    var arr [10]int
    arr := [10]int
    arr := [10]int{1,2,3}
```

**(2)切片的定义方式**
```
    var sli []int
    sli := []int
    sli := []int{1,2,3}
```
以及
```
    var sli []int
    sli = make([]int, 5) //  默认元素值为0,cap参数默认等于len
    //  等同于
    sli := []int{0,0,0,0,0}
```

**(3)Array和Slice唯一的区别就在于有无长度**

**(4)Array的数据类型由元素类型+数组长度决定，即[2]int和[3]int是不同的；而Slice的数据类型仅由元素类型决定**

**(5)Array赋值时是值拷贝；Slice赋值时是指针拷贝**
    