package strategy

import (
	"fmt"
	"gosolidprinciple/opencloseprinciple/right/entity"
)

// Implementasi OCP dengan Polymorphism

type PaymentProcessor interface {
	Process(payment *entity.Payment) error
}

type BankTransfer struct{}

func NewBankTransfer() PaymentProcessor {
	return &BankTransfer{}
}

func (b *BankTransfer) Process(payment *entity.Payment) error {
	fmt.Println("Processing bank transfer...")
	return nil
}

type CreditCard struct{}

func NewCreditCard() PaymentProcessor {
	return &CreditCard{}
}

func (c *CreditCard) Process(payment *entity.Payment) error {
	fmt.Println("Processing credit card payment...")
	return nil
}
