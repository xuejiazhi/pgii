package util

import "github.com/pierrec/lz4/v4"

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
func UnCompress(str *[]byte) ([]byte, error) {
	out := make([]byte, 10*len(*str))
	n, err := lz4.UncompressBlock(*str, out)
	out = out[:n]
	return out, err
}
