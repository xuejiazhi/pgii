package pg

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func Test_List(t *testing.T) {
	var SystemSchemaList = []string{"'pg_toast'", "'pg_temp_1'", "'pg_toast_temp_1'", "'pg_catalog'", "'information_schema'"}
	fmt.Println(strings.Join(SystemSchemaList, ","))
}

func Test_a(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(7)
	fmt.Println(num)
}
