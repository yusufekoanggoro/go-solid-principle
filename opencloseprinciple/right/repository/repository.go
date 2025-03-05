package repository

import "gosolidprinciple/opencloseprinciple/right/entity"

type PaymentRepository struct {
}

func (r *PaymentRepository) Save(payment *entity.Payment) error {
	return nil
}
