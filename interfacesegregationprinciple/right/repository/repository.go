package repository

import "gosolidprinciple/interfacesegregationprinciple/right/entity"

// WorkRepository → Hanya menangani penyimpanan pekerja
type WorkRepository interface {
	SaveWorkRecord(worker entity.Workable) error
}
