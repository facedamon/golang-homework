package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	id    int
	name  string
	age   int
	grade string
}

func insert(db *sql.DB) {
	stmt, err := db.Prepare("insert into students(name, age, grade) values (?, ?, ?)")
	if err != nil {
		fmt.Println("Failed to prepare statement:", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("张三", 20, "三年级")
	if err != nil {
		fmt.Println("Failed to execute statement", err)
		return
	}
}

func update(db *sql.DB) {
	_, err := db.Exec("update students set grade=? where name=?", "四年级", "张三")
	if err != nil {
		fmt.Println("Failed to exec update:", err)
		return
	}
}

func del(db *sql.DB) {
	_, err := db.Exec("delete from students where age<?", 15)
	if err != nil {
		fmt.Println("Failed to exec update:", err)
		return
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang_test")
	if err != nil {
		fmt.Println("connecting mysql error", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("close mysql error", err)
		}
	}(db)

	_ = db.Ping()
	fmt.Println("successfully connected to mysql")

	//insert(db)

	rows, err := db.Query("select id, name, age, grade from students where age > ?", 18)
	if err != nil {
		fmt.Println("query mysql error", err)
		return
	}
	defer rows.Close()
	students := make([]Student, 0)
	for rows.Next() {
		stu := new(Student)
		err = rows.Scan(&stu.id, &stu.name, &stu.age, &stu.grade)
		if err != nil {
			fmt.Println("scaning rows error", err)
			return
		}
		students = append(students, *stu)
	}
	fmt.Println("查询大于18岁的记录=", students)

	//update(db)
	del(db)
}
