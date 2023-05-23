package util

import (
	"archive/zip"
	"fmt"
	"github.com/pierrec/lz4/v4"
	"io"
	"os"
)

// Compress 压缩数据
func Compress(str *[]byte) {
	//获取数据的byte和buf
	data := *str
	buf := make([]byte, lz4.CompressBlockBound(len(data)))

	//定义压缩实例
	var C lz4.Compressor
	n, _ := C.CompressBlock(data, buf)
	*str = buf[:n]
}

// UnCompress 解压缩数据
func UnCompress(str []byte) ([]byte, error) {
	out := make([]byte, 20*len(str))
	n, err := lz4.UncompressBlock(str, out)
	out = out[:n]
	return out, err
}

func Tar(source, target string) error {
	zipReader, err := zip.OpenReader("example.zip")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer zipReader.Close()

	for _, file := range zipReader.File {
		dest, err := os.Create(file.Name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer dest.Close()

		src, err := file.Open()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer src.Close()

		_, err = io.Copy(dest, src)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return nil
}
