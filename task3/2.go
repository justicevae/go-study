package main

import (
	"fmt"
)

// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
//要求 ：
//定义一个 Book 结构体，包含与 books 表对应的字段。
//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func getBooksByMinPrice(db *sqlx.DB, minPrice float64) ([]Book, error) {
	var books []Book
	query := `SELECT id,title,author,price FROM books WHERE price > :min_price ORDER BY price DESC`

	args := map[string]interface{}{
		"min_price": minPrice,
	}
	stmt, _ := db.PrepareNamed(query)
	_ = stmt.Select(&books, args)
	return books, nil
}

func main() {
	expensiveBooks, _ := getBooksByMinPrice(db, 50.0)
	fmt.Println(expensiveBooks)
}
