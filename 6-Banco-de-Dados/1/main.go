package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/mydb")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	product := NewProduct("Macbook M4", 2500.00)
	err = inserProduct(db, product)
	if err != nil {
		panic(err)
	}
	product.Price = 2200
	err = UpdateProduct(db, product)
	if err != nil {
		panic(err)
	}

	// p, err := selectProduct(db, product.ID)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Product: ID=%s, Name=%s, Price=%.2f\n", p.ID, p.Name, p.Price)

	// products, err := selectAllProducts(db)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, p := range products {
	// 	fmt.Printf("Product: %v, possui o preço de: %.2f\n", p.Name, p.Price)
	// }

	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Product deleted successfully")
}

func inserProduct(db *sql.DB, product *Product) error {
	rows, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer rows.Close()

	_, err = rows.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var product Product
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func selectAllProducts(db *sql.DB) ([]*Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productList []*Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		productList = append(productList, &product)
	}
	return productList, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
