package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	sts := "CREATE TABLE IF NOT EXISTS categories(id string, name string NOT NULL, description string);"
	_, err := db.Exec(sts)
	if err != nil {
		panic(err)
	}
	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories(id,name,description) VALUES ($1,$2,$3)", id, name, description)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := make([]Category, 0)
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		categories = append(categories, Category{ID: id, Name: name, Description: description})
	}
	return categories, nil
}

func (c *Category) FindByID(id string) (Category, error) {
	var name, description string
	err := c.db.QueryRow(
		"SELECT name, description FROM categories WHERE id = $1",
		id,
	).Scan(
		&name,
		&description,
	)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Description: description}, nil
}
