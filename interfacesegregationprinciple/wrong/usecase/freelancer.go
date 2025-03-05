package usecase

import "fmt"

// ❌ Freelancer tetap harus mengimplementasikan `Eat()` dan `TakeLeave()` meskipun tidak digunakan
type Freelancer struct {
	Name string
}

func (f Freelancer) Work() error {
	fmt.Println(f.Name, "is working as a freelancer")
	return nil
}

func (f Freelancer) Eat() error {
	err := fmt.Errorf("freelancer %s cannot eat in the cafeteria", f.Name) // ❌ Tidak seharusnya ada
	fmt.Println("Error: ", err)
	return err
}

func (f Freelancer) TakeLeave() error {
	err := fmt.Errorf("freelancer %s cannot take leave", f.Name) // ❌ Tidak seharusnya ada
	fmt.Println("Error: ", err)
	return err
}
