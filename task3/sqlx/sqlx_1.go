package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type employee struct {
	Id         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

var db *sqlx.DB

func main() {
	fmt.Println("begin...")
	var err error
	db, err = sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/golang_test")
	defer db.Close()
	if err != nil {
		fmt.Println("connecting mysql failed.", err)
		return
	}

	ems := make([]employee, 0)
	err = db.Select(&ems, "select id,name,department,salary from employees where department=?", "技术部")
	if err != nil {
		fmt.Println("查询技术部sql异常", err)
		return
	}
	fmt.Println("部门是技术部的有:", ems)

	topSalary := employee{}
	err = db.Get(&topSalary, "select id, name,department,salary from employees order by salary desc limit 1")

	if err != nil {
		fmt.Println("查询最高工资sql异常", err)
		return
	}
	fmt.Println("最高工资的员工是:", topSalary)
}
