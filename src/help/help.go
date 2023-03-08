package help

import (
	"fmt"
	"pgii/src/util"
)

func Help(cms ...string) {
	if len(cms) != 1 {
		fmt.Println("Failed:Help is fail")
		return
	}
	//接参数
	param := util.TrimLower(cms[0])
	switch param {
	case "show":
		PrintShow()
	default:
		fmt.Println("Help <command>")
	}
}

func PrintShow() {
	fmt.Println("Command: <show>")
	fmt.Println("  Usage:")
	fmt.Println("	-- show db : print database list")
	fmt.Println(`	pgmcs~10.161.30.207# show db
	+-------+------------+----------+----------+------------+----------+-----------+-----------+------------+------------+-------------------------------------+
	| #OID  | DBNAME     | AUTH     | ENCODING | LC_COLLATE | LC_CTYPE | ALLOWCONN | CONNLIMIT | LASTSYSOID | TABLESPACE | ACL                                 |
	+-------+------------+----------+----------+------------+----------+-----------+-----------+------------+------------+-------------------------------------+
	| 12292 | postgres   | postgres | UTF8     | C          | C        | true      |        -1 | 12291      | pg_default | <nil>                               |
	| 16384 | clouddb[✓] | postgres | UTF8     | C          | C        | true      |        -1 | 12291      | pg_default | <nil>                               |
	| 1     | template1  | postgres | UTF8     | C          | C        | true      |        -1 | 12291      | pg_default | {=c/postgres,postgres=CTc/postgres} |
	| 12291 | template0  | postgres | UTF8     | C          | C        | false     |        -1 | 12291      | pg_default | {=c/postgres,postgres=CTc/postgres} |
	+-------+------------+----------+----------+------------+----------+-----------+-----------+------------+------------+-------------------------------------+`)
	fmt.Println("	-- show tb : print table list")
}
