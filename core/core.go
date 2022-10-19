package core

type Item struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Sku  string `db:"sku"`
}

type Collection struct {
	Id       string
	Name     string
	Articles []Item
}
