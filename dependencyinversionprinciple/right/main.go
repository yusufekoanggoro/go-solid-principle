package main

import (
	"fmt"
	"gosolidprinciple/dependencyinversionprinciple/right/repository"
	"gosolidprinciple/dependencyinversionprinciple/right/usecase"
)

func main() {
	repo := repository.NewRepository()
	usecase := usecase.NewUserUsecase(repo)

	fmt.Println(usecase.GetUserName(1)) // Output: User 1 from SQL DB
}
