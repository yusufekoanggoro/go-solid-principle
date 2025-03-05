package strategy

import (
	"fmt"
	"gosolidprinciple/opencloseprinciple/right/entity"
)

// Implementasi OCP dengan Polymorphism
// Polimorfisme dalam pemrograman berorientasi objek (OOP) adalah konsep di mana satu antarmuka (interface) dapat memiliki banyak bentuk implementasi.
// Dengan kata lain, objek yang berbeda dapat diperlakukan secara sama melalui satu antarmuka umum.
type PaymentProcessor interface {
	Process(payment *entity.Payment) error
}

type BankTransfer struct{}

func (b *BankTransfer) Process(payment *entity.Payment) error {
	fmt.Println("Processing bank transfer...")
	return nil
}

type CreditCard struct{}

func (c *CreditCard) Process(payment *entity.Payment) error {
	fmt.Println("Processing credit card payment...")
	return nil
}
