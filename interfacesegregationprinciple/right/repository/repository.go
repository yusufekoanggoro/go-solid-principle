package repository

import "gosolidprinciple/interfacesegregationprinciple/right/entity"

// UserRepository hanya untuk operasi terkait user
type UserRepository interface {
	Create(user entity.User) error
}

// OrderRepository hanya untuk operasi terkait order
type OrderRepository interface {
	Create(order entity.Order) error
}

type Repository interface {
	UserRepo() UserRepository
	OrderRepo() OrderRepository
}

type RepositoryImpl struct {
	UserRepo  UserRepository
	OrderRepo OrderRepository
}

func NewRepository() *RepositoryImpl {
	return &RepositoryImpl{
		UserRepo:  NewUserRepository(),
		OrderRepo: NewOrderRepository(),
	}
}
