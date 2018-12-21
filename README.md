# Golang-Demo

## 这是一个用于初步学习Golang的Demo仓库

### 注意点:

#### 1.Go中函数与方法要区分开(Go中 结构体 可理解为PHP中 类 )

#### 2.想要获取指针变量ptr指向的值,用*ptr获得;而想要利用结构体指针structPtr获取结构体成员,直接用structPtr.number即可

#### 3.(重点)数组Array和切片Slice

(1)数组的定义方式
```
    var arr [10]int
    arr := [10]int
    arr := [10]int{1,2,3}
```

(2)切片的定义方式
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

(3)Array和Slice唯一的区别就在于有无长度

(4)Array的数据类型由元素类型+数组长度决定，即[2]int和[3]int是不同的；而Slice的数据类型仅由元素类型决定

(5)Array赋值时是值拷贝；Slice赋值时是指针拷贝

#### 4.error异常处理

Go中的异常处理error源码,非常简单

```
    // Copyright 2011 The Go Authors. All rights reserved.
    // Use of this source code is governed by a BSD-style
    // license that can be found in the LICENSE file.

    // Package errors implements functions to manipulate errors.
    package errors

    // New returns an error that formats as the given text.
    func New(text string) error {
            return &errorString{text}
    }

    // errorString is a trivial implementation of error.
    type errorString struct {
            s string
    }

    func (e *errorString) Error() string {
            return e.s
    }
```

可以使用`import "errors"`来实现Go自带的异常,也可以自己定义,自定义方式可以参照error源码模拟实现

#### 5.自定义包

在`$GOPATH/src`下面开发自己包,看以下案例

包源码文件目录`$GOPATH/src/darius/dpkg/dpkg.go`

包源码如下:

```
    package dpkg //  包名

    func Test(s string) string {
            newStr := s
            return newStr
    }
```

开发完包后,使用Go命令`go install darius/dpkg`(注意这里指向目录)

该命令会帮我们自动生成包对象`$GOPATH/pkg/darius/dpkg.a`

接着就可以在项目中使用自己开发的包,使用方式如下

```
package main

import (
	"fmt"
	"darius/dpkg" //  注意引入的是路径
)

func main() {
	s := dpkg.Test("niubi") //  使用时才会用到包名
	fmt.Printf("你好啊,%s", s)
}
```