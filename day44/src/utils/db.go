/**
 * Created by visual studio code
 * @author: Iceberry
 * @date: 2022-01-11
 * @version: 1.0
 * 连接MySql数据库
 */
package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/go_test")
	if err != nil {
		panic(err.Error())
	}
}
