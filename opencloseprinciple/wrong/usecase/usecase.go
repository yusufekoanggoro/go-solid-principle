package usecase

import (
	"errors"
	"gosolidprinciple/opencloseprinciple/wrong/entity"
	"gosolidprinciple/opencloseprinciple/wrong/repository"
)

type PaymentUsecase struct {
	paymentRepository *repository.PaymentRepository
}

func NewPaymentUsecase(repo *repository.PaymentRepository) *PaymentUsecase {
	return &PaymentUsecase{
		paymentRepository: repo,
	}
}

// Fungsi ini melanggar Open/Closed Principle karena setiap kali ada metode pembayaran baru,
// kita harus mengubah kode dalam fungsi ini. Seharusnya, kita bisa menggunakan strategi
// atau polimorfisme untuk menambahkan metode pembayaran baru tanpa mengubah kode yang sudah ada.
func (uc *PaymentUsecase) ProcessPayment(payment *entity.Payment) error {
	if payment.Method == "bank_transfer" {
		println("Processing bank transfer...")
	} else if payment.Method == "credit_card" {
		println("Processing credit card payment...")
	} else {
		// Setiap kali ada metode pembayaran baru, kita harus menambah kondisi baru di sini,
		// yang melanggar prinsip OCP karena kode harus dimodifikasi terus-menerus.
		return errors.New("unsupported payment method")
	}

	return uc.paymentRepository.Save(payment)
}
