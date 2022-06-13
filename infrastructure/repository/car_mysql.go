package repository

import (
	"carProject/entity"
	"database/sql"
)

type CarMySQL struct {
	db *sql.DB
}

func NewCarMySQL(db *sql.DB) *CarMySQL {
	return &CarMySQL{
		db: db,
	}
}

func (repo *CarMySQL) Create(c *entity.Car) (entity.ID, error) {
	stmt, err := repo.db.Prepare("insert into car (id, brand, model, doorquantity, created_at) values(?,?,?,?,?)")
	if err != nil {
		return c.ID, err
	}
	_, err = stmt.Exec(
		c.ID, c.Brand, c.Model, c.DoorQuantity, c.CreatedAt.Format("2006-01-02"))
	if err != nil {
		return c.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return c.ID, err
	}
	return c.ID, nil
}

func (repo *CarMySQL) Get(id entity.ID) (*entity.Car, error) {
	stmt, err := repo.db.Prepare("select id, brand, model, doorquantity, created_at from car where id = ?")
	if err != nil {
		return nil, err
	}
	var c entity.Car
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&c.ID, &c.Brand, &c.Model, &c.DoorQuantity, &c.CreatedAt)
	}
	return &c, nil
}
