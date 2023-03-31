package pg

import (
	"fmt"
	"strings"
	"testing"
)

func Test_List(t *testing.T) {
	var SystemSchemaList = []string{"'pg_toast'", "'pg_temp_1'", "'pg_toast_temp_1'", "'pg_catalog'", "'information_schema'"}
	fmt.Println(strings.Join(SystemSchemaList, ","))
}
