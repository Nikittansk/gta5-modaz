package gta5modaz

type Product struct {
	Id    int    `json:"-"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}