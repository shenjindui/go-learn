package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person4 struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place4 struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db4 *sqlx.DB

func init() {

	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3308)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db4 = database
	//defer db.Close()  // 注意这行代码要写在上面err判断的下面
}

func main() {

	/*
	   _, err := Db.Exec("delete from person where user_id=?", 1)
	   if err != nil {
	       fmt.Println("exec failed, ", err)
	       return
	   }
	*/

	res, err := Db4.Exec("delete from person where user_id=?", 2)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}

	fmt.Println("delete succ: ", row)
}
