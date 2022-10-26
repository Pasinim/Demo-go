package core

type Item struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Sku  int    `db:"sku"`
}

type Collection struct {
	Id       int
	Name     string
	Articles []Item
}
