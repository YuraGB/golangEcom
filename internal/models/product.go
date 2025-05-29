package models

import "golang-server/ent"

type ExtendedProduct struct {
	*ent.Product
	Quantity int
}

type BasketProduct struct {
	ProductID int
	Quantity  int
}
