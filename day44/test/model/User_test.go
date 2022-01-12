/**
 * Created by visual studio code
 * @author: Iceberry
 * @date: 2022-01-11
 * @version: 1.0
 * User方法测试
 */
package model

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("测试开始")
	m.Run()
}

func TestUser(t *testing.T) {
	fmt.Println("开始测试User中的相关方法")
	t.Run("测试添加用户：", TestAddUser)
}

func TestAddUser(t *testing.T) {
	fmt.Println("添加测试用户：子函数")
	// user := &model.User{}
	// user.AddUser()
	// user.AddUserV2()
}
