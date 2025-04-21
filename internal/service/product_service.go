package service

func GetAllProducts() []Product {
	return []Product{
		{ID: 1, Name: "Product 1", Price: 10.0, Quantity: 1},
		{ID: 2, Name: "Product 2", Price: 20.0, Quantity: 1},
		{ID: 3, Name: "Product 3", Price: 30.0, Quantity: 1},
	}
}
