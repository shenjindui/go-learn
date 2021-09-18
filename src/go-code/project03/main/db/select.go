package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person1 struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place1 struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db1 *sqlx.DB

func init() {

	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3308)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db1 = database
	//defer db.Close()  // 注意这行代码要写在上面err判断的下面
}

func main() {

	var person []Person1 //返回一个切片
	err := Db1.Select(&person, "select user_id, username, sex, email from person where sex=?", "man")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("select succ:", cap(person))
}
