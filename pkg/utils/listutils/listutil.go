package listutils

import "fmt"

func T1() {
	fmt.Println("listutils.T1()")
}

func init() {
	fmt.Println("listutils.init() 1")
}

func init() {
	fmt.Println("listutils.init() 2")
}
