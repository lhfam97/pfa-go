package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lhfam97/pfa-go/internal/order/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

// sql create table on mysql
// CREATE TABLE `orders` (
//   `id` varchar(255) NOT NULL,
//   `price` float NOT NULL,
//   `tax` float NOT NULL,
//   `final_price` float NOT NULL,
//   PRIMARY KEY (`id`)
// )

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}
