package service

import (
	"context"
	"golang-server/ent"
	"golang-server/ent/basket"
	"golang-server/internal/models"
	"golang-server/utils"
	"log"
	"time"
)

func GetBasket(userId int, db *ent.Client, ctx context.Context) ([]*ent.Basket, error) {
	userBasket, err := db.Basket.
		Query().
		Where(basket.UserIDEQ(userId)).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	return userBasket, nil
}

func AddProductToBasket(userId int, db *ent.Client, ctx context.Context, item models.BasketProduct) (*ent.Basket, error) {
	// Перевірка: чи вже є такий запис
	existing, err := db.Basket.
		Query().
		Where(
			basket.UserID(userId),
			basket.ProductID(item.ProductID),
		).
		Only(ctx)

	if err == nil {
		log.Println("Updating existing basket item")
		// Якщо запис вже існує — оновлюємо кількість
		updated, err := db.Basket.
			UpdateOneID(existing.ID).
			SetQuantity(existing.Quantity + item.Quantity). // або item.Quantity, залежно від логіки
			SetUpdatedAt(time.Now()).
			Save(ctx)

		if err != nil {
			log.Println("Error in AddProductToBasket: Can't update existing item", err.Error())
			return nil, err
		}

		return updated, nil
	}

	// Якщо запису немає — створюємо
	if ent.IsNotFound(err) {
		log.Println("Creating new basket item")
		product, err := db.Basket.
			Create().
			SetUserID(userId).
			SetPrice(item.Price).
			SetProductID(item.ProductID).
			SetQuantity(item.Quantity).
			Save(ctx)

		if err != nil {
			log.Println("Error in AddProductToBasket: Can't create basket item", err.Error())
			return nil, err
		}

		return product, nil
	}

	// Якщо інша помилка — повертаємо її
	log.Println("Error in AddProductToBasket: unexpected error", err.Error())
	return nil, err
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

func UpdateProductQuantity(userId int, db *ent.Client, ctx context.Context, items []models.BasketProduct) *ent.Basket {
	var updated *ent.Basket

	for _, item := range items {
		basketItem, err := db.Basket.
			Query().
			Where(
				basket.UserIDEQ(userId),
				basket.ProductIDEQ(item.ProductID),
			).
			Only(ctx)

		if err != nil {
			log.Println("Error in UpdateProductQuantity: Can't get basket item", err.Error())
			continue
		}

		updatedItem, err := db.Basket.
			UpdateOne(basketItem).
			SetQuantity(item.Quantity).
			Save(ctx)

		if err != nil {
			log.Println("Error in UpdateProductQuantity: Can't update basket item", err.Error())
			continue
		}

		updated = updatedItem
	}

	return updated
}

func CreateBasket(userId int, db *ent.Client, ctx context.Context, items []models.BasketProduct) (*models.BasketInsertResult, error) {
	var result models.BasketInsertResult

	for _, item := range items {
		basket, err := db.Basket.
			Create().
			SetUserID(userId).
			SetProductID(item.ProductID).
			SetQuantity(item.Quantity).
			Save(ctx)

		if err != nil {
			result.Failures = append(result.Failures, models.BasketInsertFailure{
				ProductID: item.ProductID,
				Error:     err.Error(),
			})
			continue
		}

		result.Inserted = append(result.Inserted, basket)
	}
	log.Println("Basket created successfully")

	return &result, nil
}

func MergeBaskets(userId int, items []models.BasketProduct, existingBasket []*ent.Basket, db *ent.Client, ctx context.Context) (*models.BasketInsertResult, error) {
	var result models.BasketInsertResult

	tx, err := db.Tx(ctx)
	if err != nil {
		log.Println("Error starting transaction:", err)
		return nil, err
	}
	defer func() {
		_ = tx.Rollback() // safe rollback
	}()

	existingMap := make(map[int]*ent.Basket)
	for _, b := range existingBasket {
		existingMap[b.ProductID] = b
	}

	for _, item := range items {
		if basketItem, found := existingMap[item.ProductID]; found {
			updatedItem, err := tx.Basket.
				UpdateOneID(basketItem.ID).
				SetQuantity(basketItem.Quantity + item.Quantity).
				Save(ctx)
			if err != nil {
				log.Println("Error updating basket item:", err)
				result.Failures = append(result.Failures, models.BasketInsertFailure{
					ProductID: item.ProductID,
					Error:     err.Error(),
				})
				continue
			}
			result.Inserted = append(result.Inserted, updatedItem)
		} else {
			newItem, err := tx.Basket.
				Create().
				SetUserID(userId).
				SetProductID(item.ProductID).
				SetQuantity(item.Quantity).
				Save(ctx)
			if err != nil {
				log.Println("Error creating basket item:", err)
				result.Failures = append(result.Failures, models.BasketInsertFailure{
					ProductID: item.ProductID,
					Error:     err.Error(),
				})
				continue
			}
			result.Inserted = append(result.Inserted, newItem)
		}
	}
	// Commit the transaction
	if err := tx.Commit(); err != nil {
		log.Println("Error committing transaction:", err)
		return nil, err
	}

	return &result, nil
}

func GetBasketProducts(basket []*ent.Basket) ([]*ent.Product, error) {
	var productIDs []int
	for _, b := range basket {
		productIDs = append(productIDs, b.ProductID)
	}

	productList, err := utils.GetProductsByIDs(productIDs)
	if err != nil {
		return nil, err
	}

	return productList, nil
}
