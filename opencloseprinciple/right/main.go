package main

import (
	"fmt"
	"gosolidprinciple/opencloseprinciple/right/entity"
	"gosolidprinciple/opencloseprinciple/right/repository"
	"gosolidprinciple/opencloseprinciple/right/usecase"
)

func main() {
	paymentRepo := repository.PaymentRepository{}
	usacase := usecase.NewPaymentUsecase(&paymentRepo)

	payment := &entity.Payment{
		ID:     "234",
		Amount: 12000,
		Method: "bank_transfer",
	}

	if err := usacase.ProcessPayment(payment); err != nil {
		fmt.Println("Error processing payment:", err)
	} else {
		fmt.Println("Payment processed successfully")
	}
}
