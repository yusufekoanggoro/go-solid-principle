package main

type User struct {
	id   string
	name string
}

// UserRepository bertanggung jawab atas pengambilan data.
type UserRepository struct {
}

func (repo *UserRepository) GetUser(id int) *User {
	return &User{
		id:   "123",
		name: "yusuf",
	}
}

// UserUsecase bertanggung jawab atas logika bisnis.
type UserUsecase struct {
	repo *UserRepository
}

func (uc *UserUsecase) GetUser(id int) *User {
	return uc.repo.GetUser(id)
}
