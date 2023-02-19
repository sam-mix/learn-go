package utils

import (
	"fmt"
	"learn-go/pkg/utils/listutils"
)

func T1() {
	listutils.T1()
	fmt.Println("utils.T1()")
}

func init() {
	fmt.Println("utils.init() 1")
}

func init() {
	fmt.Println("utils.init() 2")
}

func init() {
	fmt.Println("utils.init() 3")
}
