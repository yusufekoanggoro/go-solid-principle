package repository

import "gosolidprinciple/interfacesegregationprinciple/right/entity"

type Order struct {
}

func NewOrderRepository() OrderRepository {
	return &Order{}
}

func (or *Order) Create(entity.Order) error {
	return nil
}
