package main

import (
	"fmt"
	"darius/dpkg"
)

func main() {
	s := dpkg.Test("niubi")
	fmt.Printf("你好啊,%s", s)
}

