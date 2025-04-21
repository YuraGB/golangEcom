package service

type Basket struct {
	ID uint8
}

type Product struct {
	ID       uint8
	Name     string
	Price    float32
	Quantity uint8
}

func GetBasket() []Basket {
	return []Basket{
		{ID: 1},
	}
}

func AddProductToBasket(product Product) Product {
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

func UpdateProductQuantity(productID uint8, quantity uint8) Product {
	// Logic to update product quantity in the basket

	updatedProduct := Product{ID: productID, Name: "Updated Product", Price: 10.0, Quantity: quantity + 1}
	return updatedProduct
}
