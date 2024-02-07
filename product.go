package gta5modaz

type Product struct {
	Id    int    `json:"-" db:"id"`
	Name  string `json:"name" db:"name" binding:"required"`
	Price int    `json:"price" db:"price"`
}