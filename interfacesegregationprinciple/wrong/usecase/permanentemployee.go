package usecase

import "fmt"

// ❌ PermanentEmployee harus mengimplementasikan semua metode
type PermanentEmployee struct {
	Name string
}

func (p PermanentEmployee) Work() error {
	fmt.Println(p.Name, "is working")
	return nil
}

func (p PermanentEmployee) Eat() error {
	fmt.Println(p.Name, "is eating in the cafeteria")
	return nil
}

func (p PermanentEmployee) TakeLeave() error {
	fmt.Println(p.Name, "is taking a leave")
	return nil
}
