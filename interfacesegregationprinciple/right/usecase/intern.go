package usecase

import "fmt"

// Intern → Bisa bekerja dan makan, tapi tidak bisa cuti
type Intern struct {
	Name string
}

func (i Intern) Work() error {
	fmt.Println(i.Name, "is working as an intern")
	return nil
}

func (i Intern) Eat() error {
	fmt.Println(i.Name, "is eating in the cafeteria")
	return nil
}
