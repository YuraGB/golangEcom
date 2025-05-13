package service

import (
	"golang-server/ent"
	"golang-server/utils"
	"log"
)

func GetAllCategories() ([]string, error) {
	var categoryList []string
	
	err := utils.GetDataFromAPI("/products/category-list", &categoryList)
	
	if err != nil {
		return nil, err
	}

	return categoryList, nil

}


type ApiResponse struct {
    Products []*ent.Product `json:"products"`
}

func GetCategoryProducts(name string) (ApiResponse, error) {
  var productList ApiResponse

  err := utils.GetDataFromAPI("/products/category/" + name, &productList);

  if err != nil {
	log.Println("28: category_service", err)
	return ApiResponse{}, err 
}

  return productList, nil
}
