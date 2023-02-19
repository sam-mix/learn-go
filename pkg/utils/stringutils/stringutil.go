package stringutils

import "fmt"

func T1() {
	fmt.Println("stringutils.T1()")
}

func init() {
	fmt.Println("stringutils.init() 1")
}

func init() {
	fmt.Println("stringutils.init() 2")
}

func init() {
	fmt.Println("stringutils.init() 3")
}
