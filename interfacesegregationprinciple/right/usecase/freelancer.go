package usecase

import "fmt"

// Freelancer → Hanya bisa bekerja
type Freelancer struct {
	Name string
}

func (f Freelancer) Work() error {
	fmt.Println(f.Name, "is working as a freelancer")
	return nil
}
