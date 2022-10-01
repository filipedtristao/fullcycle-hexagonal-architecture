package db_test

import (
	"github.com/filipedtristao/hexagonal-architecture/adapters/db"
	"github.com/filipedtristao/hexagonal-architecture/application"
	"github.com/stretchr/testify/require"
	"database/sql"
	"testing"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")

	createTable()
	createProduct()
}

func createTable() {
	sql := `
		CREATE TABLE products (
			id string,
			name string,
			price float,
			status string
		);
	`

	stmt, err := Db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec()
}

func createProduct() {
	sql := `
		INSERT INTO products (id, name, price, status) values ("1", "Product 1", 10.0, "enabled");
	`

	stmt, err := Db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("1")

	require.Nil(t, err)
	require.NotNil(t, product)
	require.Equal(t, "1", product.GetId())
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 10.0, product.GetPrice())
	require.Equal(t, "enabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product 1"
	product.Price = 10.0

	result, err := productDb.Save(product)

	require.Nil(t, err)
	require.NotNil(t, product)
	require.Equal(t, product.Id, result.GetId())
	require.Equal(t, product.Name, result.GetName())
	require.Equal(t, product.Price, result.GetPrice())
	require.Equal(t, product.Status, result.GetStatus())

	product.Name = "Updated name"
	product.Price = 20.0
	product.Status = "disabled"

	result, err = productDb.Save(product)

	require.Nil(t, err)
	require.NotNil(t, product)
	require.Equal(t, product.Id, result.GetId())
	require.Equal(t, product.Name, result.GetName())
	require.Equal(t, product.Price, result.GetPrice())
	require.Equal(t, product.Status, result.GetStatus())
}