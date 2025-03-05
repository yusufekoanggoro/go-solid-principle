package main

import (
	"fmt"
	"gosolidprinciple/dependencyinversionprinciple/right/repository"
	"gosolidprinciple/dependencyinversionprinciple/right/usecase"
)

func main() {
	repo := repository.NewRepository()
	usease := usecase.NewUserUsecase(repo)

	fmt.Println(usease.GetUserName(1)) // Output: User 1 from SQL DB
}
