package pg

import (
	"fmt"
	"github.com/pierrec/lz4/v4"
	"io/ioutil"
	"os"
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

func TestParams_LoadTable(t *testing.T) {
	fileName := "E:\\work\\go\\src\\pgii\\dump_table_user_1681291521.pgi"
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("cannot read the file", err)
		return
	}
	// UPDATE: close after checking error
	defer f.Close() //

	content, err := ioutil.ReadAll(f)
	fmt.Println(string(content))

	out := make([]byte, 10*len(content))
	n, err := lz4.UncompressBlock(content, out)
	out = out[:n] // uncompressed data
}
