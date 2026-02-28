package algorithm

import (
	"fmt"
	"testing"
)

func Test_front(t *testing.T) {
	//a := front("1+2+3-4")
	//fmt.Println(a)
	//fmt.Println(beforeCal(a))

	b := front("1+2*3*4/5*6")
	fmt.Println(b)
	fmt.Println(beforeCal(b))
}
