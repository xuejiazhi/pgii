package main

type PgDatabase struct {
	Oid     int    `json:"oid"`
	DatName string `json:"datname"`
	//DatDba        int    `json:"datdba"`
	//Encoding      int    `json:"encoding"`
	//DatCollate    string `json:"datcollate"`
	//DatCType      string `json:"datctype"`
	//DatAllowConn  bool   `json:"datallowconn"`
	//DatConnLimit  int    `json:"datconnlimit"`
	//DatLastSysOid int    `json:"datlastsysoid"`
	//DatTableSpace int    `json:"dattablespace"`
	//DatAcl        string `json:"datacl"`
}
