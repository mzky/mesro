package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"strconv"
)

var (
	db, _ = leveldb.OpenFile("DB", nil)
	dbret string
)

func main() {

	defer db.Close()
	dbPut("abc", "")
	dbret, _ = dbGet("abc")

	for i := 0; i < 1000; i++ {
		dbPut("abc", dbret+strconv.Itoa(i))
		dbret, _ = dbGet("abc")
	}
	fmt.Printf("--------%s\n", dbret)
}

func dbPut(key string, value string) {
	db.Put([]byte(key), []byte(value), nil)
}
func dbGet(key string) (string, error) {
	value, err := db.Get([]byte(key), nil)
	return string(value[:]), err
}
