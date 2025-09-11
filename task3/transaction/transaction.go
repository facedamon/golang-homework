package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type accounts struct {
	id      int
	balance float64
}

type transactions struct {
	id            int
	fromAccountId int
	toAccountId   int
	amount        float64
}

func main() {
	fmt.Println("begin...")
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang_test")
	if err != nil {
		fmt.Println("connecting mysql error", err)
		return
	}
	defer func(db *sql.DB) {
		if err := db.Close(); err != nil {
			fmt.Println("close mysql error", err)
			return
		}
	}(db)

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("db begin transaction failed.", err)
		return
	}
	A := accounts{1, 0}
	B := accounts{2, 0}

	err = tx.QueryRow("select balance from accounts where id=?", A.id).Scan(&A.balance)
	if err != nil {
		tx.Rollback()
		fmt.Println("查询A账户余额异常", err)
		return
	}
	if A.balance < 100 {
		tx.Rollback()
		fmt.Println("A账户余额=", A.balance, "不足100元，回滚事务")
		return
	}
	_, err = tx.Exec("update accounts set balance = balance - ? where id = ?", 100, A.id)
	if err != nil {
		tx.Rollback()
		fmt.Println("从A账户扣款失败", err)
		return
	}
	_, err = tx.Exec("update accounts set balance = balance + ? where id = ?", 100, B.id)
	if err != nil {
		tx.Rollback()
		fmt.Println("向B账户存款失败", err)
		return
	}

	//写入交易表
	_, err = tx.Exec("insert into transactions(from_account_id, to_account_id, amount) values (?,?,?)", A.id, B.id, 100)
	if err != nil {
		tx.Rollback()
		fmt.Println("写入交易表失败", err)
		return
	}

	//提交事务
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		fmt.Println("提交事务失败", err)
		return
	}

	fmt.Println("交易成功!!!")

}
