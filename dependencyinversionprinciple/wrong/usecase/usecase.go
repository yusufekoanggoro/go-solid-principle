package usecase

import "gosolidprinciple/dependencyinversionprinciple/wrong/repository"

// ❌ UserUseCase bergantung langsung pada struct Repository, bukan interface
type UserUseCase struct {
	repo *repository.Repository
}

// ❌ Konstruktor menerima struct Repository langsung, bukan interface
func NewUserUseCase(repo *repository.Repository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

// ❌ UserUseCase memanggil metode GetUser dari struct Repository langsung
func (s *UserUseCase) GetUserName(id int) string {
	return s.repo.GetUser(id)
}
