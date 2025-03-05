package usecase

import (
	"errors"
	"gosolidprinciple/opencloseprinciple/right/entity"
	"gosolidprinciple/opencloseprinciple/right/repository"
	"gosolidprinciple/opencloseprinciple/right/strategy"
)

type PaymentUsecase struct {
	paymentRepository *repository.PaymentRepository
	strategies        map[string]strategy.PaymentProcessor
}

func NewPaymentUsecase(repo *repository.PaymentRepository) *PaymentUsecase {
	return &PaymentUsecase{
		paymentRepository: repo,
		strategies: map[string]strategy.PaymentProcessor{
			"bank_transfer": strategy.NewBankTransfer(),
			"credit_card":   strategy.NewCreditCard(),
		},
	}
}

func (uc *PaymentUsecase) ProcessPayment(payment *entity.Payment) error {
	processor, exist := uc.strategies[payment.Method]
	if !exist {
		return errors.New("unsupported payment")
	}

	if err := processor.Process(payment); err != nil {
		return err
	}

	return uc.paymentRepository.Save(payment)
}
