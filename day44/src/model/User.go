/**
 * Created by visual studio code
 * @author: Iceberry
 * @date: 2022-01-11
 * @version: 1.0
 * 用户模型
 */
package model

import (
	"day44/src/utils"
	"fmt"
)

/**
 * User 结构体
 */
type User struct {
	Id       int
	Username string
	Password string
	Email    string
}

/**
 * AddUser 添加用户
 * @version: 1.0
 */
func (user *User) AddUser() error {
	sqlString := "INSERT INTO users(username,password,email) VALUES(?,?,?)"

	//预处理
	inStmt, err := utils.Db.Prepare(sqlString)
	if err != nil {
		fmt.Println("预编译失败：", err)
		return err
	}

	//执行
	_, err2 := inStmt.Exec("admin", "123456", "admin@admin.com")
	if err2 != nil {
		fmt.Println("执行出现异常")
		return err2
	}
	return nil
}

/**
 * AddUserV2 添加用户
 * @version: 2.0
 */
func (user *User) AddUserV2() error {
	sqlString := "INSERT INTO users(username,password,email) VALUES(?,?,?)"

	//执行
	_, err2 := utils.Db.Exec(sqlString, "admin2", "888888", "admin2@admin.com")
	if err2 != nil {
		fmt.Println("执行出现异常")
		return err2
	}
	return nil
}
