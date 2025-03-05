package repository

import "gosolidprinciple/opencloseprinciple/wrong/entity"

type PaymentRepository struct {
}

func (r *PaymentRepository) Save(payment *entity.Payment) error {
	return nil
}
