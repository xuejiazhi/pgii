package pg

import (
	"fmt"
	"testing"
	"time"
)

//func Test_RemoveNullStr(t *testing.T) {
//	str := "show     table     like     pub"
//	st := strings.Split(str, " ")
//	RemoveNullStr(&st)
//	fmt.Println(st)
//}

func Test_c(t *testing.T) {
	fmt.Print(time.Now().Unix() > 0)
}

func Test_b(t *testing.T) {
	a := 1 << 1
	fmt.Println(a)
	b := 1 << 2
	fmt.Println(b)
	fmt.Println(1 << 3)
	fmt.Println(1 << 4)
	fmt.Println(1 << 5)
	fmt.Println(1 << 6)
}
