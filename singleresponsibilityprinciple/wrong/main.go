package main

type User struct {
	id   string
	name string
}

type UserUsecase struct {
}

// ❌ Melanggar SRP: `UserUsecase` tidak hanya bertanggung jawab atas logika bisnis,
// tetapi juga menangani operasi repository secara langsung.
func (uc *UserUsecase) GetUser(id int) *User {
	return uc.RepositoryGetUser(id) // ❌ Hardcoded ID, seharusnya menggunakan id dari parameter.
}

// ❌ Melanggar SRP: `UserUsecase` tidak seharusnya menangani akses ke data (repository).
// Seharusnya ada lapisan Repository yang bertanggung jawab atas pengambilan data.
func (uc *UserUsecase) RepositoryGetUser(id int) *User {
	return &User{
		id:   "123",
		name: "yusuf",
	}
}
