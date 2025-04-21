package service

func GetAllCategories() []string {
	return []string{
		"Category 1",
		"Category 2",
		"Category 3",
	}
}

func GetCategoryByID(id int) string {
	if id < 0 || id >= len(GetAllCategories()) {
		return "Category not found"
	}
	return GetAllCategories()[id]
}
