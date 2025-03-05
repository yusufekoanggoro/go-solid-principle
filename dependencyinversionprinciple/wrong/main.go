package main

import (
	"fmt"
	"gosolidprinciple/dependencyinversionprinciple/wrong/repository"
	"gosolidprinciple/dependencyinversionprinciple/wrong/usecase"
)

func main() {
	repo := &repository.Repository{}
	usease := usecase.NewUserUseCase(repo)

	fmt.Println(usease.GetUserName(2)) // Output: User 2 from SQL DB
}
