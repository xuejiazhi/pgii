package pg

import (
	"fmt"
	"github.com/pierrec/lz4/v4"
	"io/ioutil"
	"os"
	"pgii/src/util"
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

func Test_Gzip(t *testing.T) {
	//初始化创建一个压缩文件
	//entry, _ := ioutil.ReadDir("dump_schema_db_mcs.com_1684141547")

	//path := "E:/work/go/src/pgii/dump_schema_db_mcs.com_1684141547"
	//var files []*os.File
	//fileList, _ := ioutil.ReadDir(path)
	//for _, file := range fileList {
	//	if file.IsDir() {
	//		continue
	//	} else {
	//		o, _ := os.OpenFile(path+"/"+file.Name(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	//		files = append(files, o)
	//		//o.Close()
	//	}
	//}
	sourceDir := "E:/work/go/src/pgii/dump_schema_db_mcs.com_1684141547/"
	targetDir := "E:/work/go/src/pgii/"

	if err := util.Tar(sourceDir, targetDir); err != nil {
		fmt.Printf("tar failed %s\n", err.Error())
	}
}

func Test_A(t *testing.T) {
	a := "abcd"
	c := len([]rune(a))
	fmt.Println(c)
}
