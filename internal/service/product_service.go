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

func GetSearchProducts (param string)([]*ent.Product, error) {
	var data ProductList

	err := utils.GetDataFromAPI("/products/search?q=" + param, &data)
	if err != nil {		
		return nil, err
	}

	return data.Products, nil
}

func GetProduct (id string) (*ent.Product, error) {
	var product *ent.Product

	err := utils.GetDataFromAPI("/product/" + id, &product)
	if err != nil {		
		return nil, err
	}

	return product, nil
}
