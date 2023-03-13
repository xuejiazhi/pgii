package util

import (
	"fmt"
	"github.com/pierrec/lz4/v4"
	"testing"
)

func Test_CompressBlock(t *testing.T) {
	s := `
-- DROP Table;
DROP Table abc;
CREATE TABLE "db_mcs.com".abc (
    a varchar NOT NULL,
    b varchar NOT NULL,
    c varchar NOT NULL DEFAULT 'abv'::character varying,
    d serial4 NOT NULL,
    column1 int2 NOT NULL GENERATED ALWAYS AS IDENTITY,
    e bigserial NOT NULL,
    f smallserial NOT NULL,
    g serial4 NOT NULL,
    CONSTRAINT abc_pk PRIMARY KEY (a),
    CONSTRAINT abc_un UNIQUE (b)
);
`
	data := []byte(s)
	fmt.Println("data=>", data)
	buf := make([]byte, lz4.CompressBlockBound(len(data)))

	var c lz4.Compressor
	n, err := c.CompressBlock(data, buf)
	if err != nil {
		fmt.Println(err)
	}
	if n >= len(data) {
		fmt.Printf("`%s` is not compressible", s)
	}
	buf = buf[:n] // compressed data
	fmt.Println(string(buf))

	// Allocate a very large buffer for decompression.
	out := make([]byte, 10*len(data))
	n, err = lz4.UncompressBlock(buf, out)
	if err != nil {
		fmt.Println(err)
	}
	out = out[:n] // uncompressed data

	fmt.Println(string(out[:len(s)]))

	//Output:
	//hello world
}
