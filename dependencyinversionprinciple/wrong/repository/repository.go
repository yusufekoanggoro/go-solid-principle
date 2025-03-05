package repository

import "fmt"

// Repository langsung digunakan oleh UserService tanpa abstraksi
type Repository struct{}

func NeRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetUser(id int) string {
	return fmt.Sprintf("User %d from SQL DB", id)
}
