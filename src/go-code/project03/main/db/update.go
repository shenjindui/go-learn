package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person2 struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place2 struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db2 *sqlx.DB

func init() {

	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3308)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db2 = database
	//defer db.Close()  // 注意这行代码要写在上面err判断的下面
}

func main() {

	res, err := Db2.Exec("update person set username=? where user_id=?", "stu0003", 3)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	row, err := res.RowsAffected() //row 影响的行数
	if err != nil {
		fmt.Println("rows failed, ", err)
	}
	fmt.Println("update succ:", row)

}
