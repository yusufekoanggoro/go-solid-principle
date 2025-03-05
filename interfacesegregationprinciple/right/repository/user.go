package repository

import "gosolidprinciple/interfacesegregationprinciple/right/entity"

type User struct {
}

func NewUserRepository() UserRepository {
	return &User{}
}

func (or *User) Create(entity.User) error {
	return nil
}
