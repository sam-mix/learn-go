package timeutils

import (
	"fmt"
	"learn-go/pkg/utils/stringutils"
)

func T1() {
	stringutils.T1()
	fmt.Println("timeutils.T1()")
}

func init() {
	fmt.Println("timeutils.init() 1")
}

func init() {
	fmt.Println("timeutils.init() 2")
}

func init() {
	fmt.Println("timeutils.init() 3")
}

func init() {
	fmt.Println("timeutils.init() 4")
}
