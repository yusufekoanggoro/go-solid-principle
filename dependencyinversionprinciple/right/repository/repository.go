package repository

import "fmt"

// Abstraksi Repository (Mengikuti DIP)
type Repository interface {
	GetUser(id int) string
}

// Implementasi Repository (SQL)
type RepositoryImpl struct{}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) GetUser(id int) string {
	return fmt.Sprintf("User %d from SQL DB", id)
}
