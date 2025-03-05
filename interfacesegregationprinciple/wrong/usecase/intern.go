package usecase

import "fmt"

// ❌ Intern juga harus mengimplementasikan `TakeLeave()` meskipun tidak digunakan
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

func (i Intern) TakeLeave() error {
	err := fmt.Errorf("intern %s cannot take leave", i.Name) // ❌ Tidak seharusnya ada\
	fmt.Println("Error: ", err)
	return err
}
