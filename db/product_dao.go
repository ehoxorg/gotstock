// Functions to act as Data Access Layer for Product Resource.
package db

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var (
	conn string
	db   *sql.DB
)

func openConn() *sql.DB {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Println("Ping failed!")
	}
	return db
}

func InsertProduct(p Product) (*Product, error) {
	_, err := db.Exec("INSERT INTO product (product_code, product_name, stock_quantity) VALUES ($1, $2, $3)",
		p.ProductCode,
		p.Name,
		p.StockQuantity,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	res, err, _ := GetProduct(&p.ProductCode, nil)
	return res, err
}

func UpdateProduct(p Product) (*Product, error) {
	_, err := db.Exec("UPDATE product SET product_name=$1, stock_quantity=$2 WHERE product_code=$3",
		p.Name,
		p.StockQuantity,
		p.ProductCode,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	res, err, _ := GetProduct(&p.ProductCode, nil)
	return res, err
}

func DeleteProduct(code string) error {
	_, err := db.Exec("DELETE FROM product WHERE product_code = $1", code)
	return err
}

func GetAllProducts() ([]Product, error) {
	rows, err := db.Query("SELECT product_code, product_name, stock_quantity from product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ps := []Product{}
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ProductCode, &p.Name, &p.StockQuantity)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		ps = append(ps, p)
	}
	return ps, nil
}

// Get product from DB based on code or id value. If returned bool is true, then no qualifying row was found.
func GetProduct(code *string, id *int64) (*Product, error, bool) {
	var r *sql.Row
	if id != nil {
		r = db.QueryRow("SELECT product_code, product_name, stock_quantity from product where id = $1", *id)
	} else {
		r = db.QueryRow("SELECT product_code, product_name, stock_quantity from product where product_code = $1", *code)
	}
	return rowToProduct(r)
}

func rowToProduct(r *sql.Row) (*Product, error, bool) {
	var p *Product = &Product{}
	err := r.Scan(&p.ProductCode, &p.Name, &p.StockQuantity)
	if err != nil {
		log.Print(err)
		if err == sql.ErrNoRows {
			return nil, err, true
		} else {
			return nil, err, false
		}
	}
	return p, nil, false
}

func init() {
	conn = os.Getenv("PSQL_CONN_STRING")
	db = openConn()
	runMig, err := strconv.ParseBool(os.Getenv("RUN_UP_MIGRATION"))
	if err != nil {
		log.Default().Println("Cannot parse RUN_UP_MIGRATION env variable.")
	}

	if runMig {
		driver, err := postgres.WithInstance(db, &postgres.Config{DatabaseName: "gotstockapi"})
		if err != nil {
			log.Fatal(err)
		}
		log.Default().Println("Running migration...")
		m, err := migrate.NewWithDatabaseInstance(
			"file://db/migrations",
			"postgres", driver)
		if err != nil {
			log.Fatal(err)
		}
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	}
}
