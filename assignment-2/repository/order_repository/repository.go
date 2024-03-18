package order_repository

import (
	"assignment-2/entity"
	"assignment-2/pkg/errs"
)

type Repository interface {
	ReadOrderById(orderId int) (*entity.Order, errs.Error)
	CreateOrderWithItems(orderPayload entity.Order, itemPayload []entity.Item) errs.Error
	ReadOrders() ([]OrderItemMapped, errs.Error)
	UpdateOrder(orderPayload entity.Order, itemPayload []entity.Item) errs.Error
	DeleteOrderById(orderId int) errs.Error
}
