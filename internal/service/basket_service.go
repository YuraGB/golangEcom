package service

import "golang-server/ent"

type Basket struct {
	ID uint8
}



func GetBasket() []Basket {
	return []Basket{
		{ID: 1},
	}
}

func AddProductToBasket(product *ent.Product)  *ent.Product {
	// Logic to add product to basket
	return product
}

func RemoveProductFromBasket(productID uint8) uint8 {
	// Logic to remove product from basket

	return productID
}

func ClearBasket() bool {
	// Logic to clear the basket

	return true
}

func RemoveBasket() bool {
	// Logic to remove the basket
	return true
}

func UpdateProductQuantity(productID uint8, quantity uint8) *ent.Product {
	// Logic to update product quantity in the basket

	updatedProduct := ent.Product{}
	return &updatedProduct
}
