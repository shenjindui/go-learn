package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/**
  插入操作 使用第三方开源的mysql库: github.com/go-sql-driver/mysql （mysql驱动） github.com/jmoiron/sqlx （基于mysql驱动的封装）
*/
type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

/**
  初始化连接
*/
func init() {
	//Open()创建连接打开数据库
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3308)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	//defer database.Close() // 注意这行代码要写在上面err判断的下面
}

func main() {
	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Printf("%T", r)
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert success:", id)
}
