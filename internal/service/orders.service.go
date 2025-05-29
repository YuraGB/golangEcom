package service

import (
	"context"
	"golang-server/ent"
	"golang-server/ent/order"
	"golang-server/internal/models"
)

func GetOrders(userId int, db *ent.Client, ctx context.Context) ([]*ent.Order, error) {
	// Query the database to get the user's orders
	orders, err := db.Order.
		Query().
		Where(order.UserIDEQ(userId)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	// Return the list of orders
	return orders, nil
}

func GetOrderById(orderId int, db *ent.Client, ctx context.Context) (*ent.Order, error) {
	// Query the database to get the order by ID
	foundOrder, err := db.Order.
		Query().
		Where(order.IDEQ(orderId)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	// Return the found order
	return foundOrder, nil
}

func CreateOrder(userId int, db *ent.Client, ctx context.Context, orderData models.OrderInput) (*ent.Order, error) {
	// Create a new order in the database
	newOrder, err := db.Order.
		Create().
		SetStatus(order.Status(orderData.Status)).
		SetState(orderData.Order.State).
		SetUserID(userId).
		SetAddress(orderData.Order.Address).
		SetCity(orderData.Order.City).
		SetPaymentType(order.PaymentType(orderData.Order.PaymentType)).
		SetZip(orderData.Order.Zip).
		SetTotalPrice(orderData.TotalPrice).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	// Return the created order
	return newOrder, nil
}
