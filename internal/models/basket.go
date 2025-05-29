package models

import "golang-server/ent"

type BasketInsertFailure struct {
	ProductID int    `json:"product_id"`
	Error     string `json:"error"`
}

type BasketInsertResult struct {
	Inserted []*ent.Basket
	Failures []BasketInsertFailure
}
