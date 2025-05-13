package service

import (
	"golang-server/ent"
	"golang-server/utils"
)

type ProductList struct {
	Products []*ent.Product `json:"products"`
}

func GetAllProducts() ([]*ent.Product, error) {
	var data ProductList

	err := utils.GetDataFromAPI("/products", &data)
	if err != nil {
		
		return nil, err // Handle error if needed
	}

	return data.Products, nil
}
