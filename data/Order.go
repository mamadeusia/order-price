package data

import (
	"context"
	"fmt"
	"preh/server"
)

var (
	CreateOrder     = createOrder
	GetOrderDetails = getOrderDetails
	UpdateOrder     = updateOrder
)

func createOrder(ctx context.Context, orderCreate *OrderCreate) (*Order, error) {
	orderHash := asSha256(orderCreate)
	server.GetCurrentPrice(ctx, orderCreate.Pair)
	order := &Order{
		ID:     orderHash,
		UserId: orderCreate.UserId,
		Price:  orderCreate.Price,
		Amount: orderCreate.Amount,
		Pair:   orderCreate.Pair,
	}
	user, err := getUserDetails(ctx, order.UserId)
	user.Ordersid = append(user.Ordersid, order.ID)

	//TODO : i have to revert transaction if it failed
	if err = updateToRedis(ctx, user); err != nil {
		return nil, err
	}
	return order, nil

}

func getOrderDetails(ctx context.Context, id string) (*Order, error) {
	order := &Order{
		ID: id,
	}
	if err := getFromRedis(ctx, order); err != nil {
		return nil, err
	}
	fmt.Println(order)
	return order, nil

}
func updateOrder(ctx context.Context, orderUpdate *OrderUpdate) error {
	order := &Order{
		ID: orderUpdate.ID,
	}
	if err := getFromRedis(ctx, order); err != nil {
		return err
	}
	order.Price = orderUpdate.Price
	return updateToRedis(ctx, order)
}
