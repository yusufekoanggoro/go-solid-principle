package usecase

import "gosolidprinciple/dependencyinversionprinciple/right/repository"

// ✅ UserUseCase hanya bergantung pada interface, bukan implementasi konkret.
// Ini memastikan bahwa UserUseCase tidak terikat dengan implementasi spesifik repository.
type UserUseCase struct {
	repo repository.Repository
}

// ✅ Konstruktor menerima interface Repository, bukan struct konkret.
// Dengan cara ini, kita bisa mengganti implementasi repository dengan mudah tanpa mengubah kode utama.
func NewUserUsecase(repo repository.Repository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

// ✅ UserUseCase menggunakan metode dari interface Repository tanpa tahu implementasi spesifiknya.
// Ini mematuhi Dependency Inversion Principle (DIP), karena high-level module tidak bergantung pada low-level module, tetapi keduanya bergantung pada abstraksi.
func (s *UserUseCase) GetUserName(id int) string {
	return s.repo.GetUser(id)
}
