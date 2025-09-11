package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type book struct {
	Id     sql.NullInt32   `db:"id"`
	Title  sql.NullString  `db:"title"`
	Author sql.NullString  `db:"author"`
	Price  sql.NullFloat64 `db:"price"`
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

	books := make([]book, 0)
	err = db.Select(&books, "select id, title, author, price from books where price>?", 50)
	if err != nil {
		fmt.Println("查询超过50元价格书籍sql异常", err)
		return
	}

	for _, b := range books {
		fmt.Println(b.Id.Int32, b.Title.String, b.Author.String, b.Price.Float64)
	}
}
