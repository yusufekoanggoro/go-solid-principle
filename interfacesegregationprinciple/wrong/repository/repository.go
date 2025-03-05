package repository

import "gosolidprinciple/interfacesegregationprinciple/wrong/entity"

// Pelanggaran ISP: Interface ini memaksa implementasi UserRepository
// untuk juga menangani operasi CreateOrder, yang seharusnya berada dalam OrderRepository.
type UserRepository interface {
	CreateUser(user entity.User) error
	CreateOrder(order entity.Order) error
}

// Pelanggaran ISP: Struct RepositoryImpl seharusnya tidak mengelompokkan UserRepository
// yang mencampur dua tanggung jawab berbeda.
type RepositoryImpl struct {
	UserRepo UserRepository
}

func NewRepository() *RepositoryImpl {
	return &RepositoryImpl{
		UserRepo: NewUserRepository(), // UserRepository seharusnya hanya menangani user, bukan order.
	}
}
