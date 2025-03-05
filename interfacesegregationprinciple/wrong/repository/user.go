package repository

import "gosolidprinciple/interfacesegregationprinciple/wrong/entity"

// Pelanggaran ISP: UserRepository seharusnya hanya menangani user,
// tetapi memiliki method CreateOrder yang seharusnya menjadi tanggung jawab OrderRepository.
type User struct {
}

func NewUserRepository() UserRepository {
	return &User{}
}

// Method ini sesuai dengan tanggung jawab UserRepository.
func (or *User) CreateUser(entity.User) error {
	return nil
}

// Pelanggaran ISP: Method ini tidak seharusnya ada di UserRepository,
// karena tidak berhubungan dengan User. Seharusnya ada dalam OrderRepository.
func (or *User) CreateOrder(entity.Order) error {
	return nil
}
