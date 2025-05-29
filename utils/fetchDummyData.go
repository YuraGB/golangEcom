package utils

import (
	"encoding/json"
	"fmt"
	"golang-server/ent"
	"io"
	"log"
	"net/http"
	"time"
)

const dummyBase = "https://dummyjson.com"

func GetDataFromAPI(endpoint string, out interface{}) error {
	fullURL := fmt.Sprintf("%s%s", dummyBase, endpoint)

	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(fullURL)
	if err != nil {
		return fmt.Errorf("failed to fetch: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, out); err != nil {
		log.Println("Response body:", string(body)) // Додаткове логування
		return fmt.Errorf("json error: %w", err)
	}

	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("json error: %w", err)
	}

	return nil
}

func GetProductsByIDs(productIDs []int) ([]*ent.Product, error) {
	var products []*ent.Product

	// Перебираємо список ID та фетчимо кожен продукт
	for _, id := range productIDs {
		endpoint := fmt.Sprintf("/products/%d", id)
		var product ent.Product

		// Отримуємо продукт за його ID
		if err := GetDataFromAPI(endpoint, &product); err != nil {
			return nil, err
		}

		// Додаємо продукт у список
		products = append(products, &product)
	}

	return products, nil
}
