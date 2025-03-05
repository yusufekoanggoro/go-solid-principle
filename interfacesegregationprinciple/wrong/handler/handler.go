package handler

import "gosolidprinciple/interfacesegregationprinciple/wrong/entity"

// ❌ Memaksa semua karyawan menggunakan interface yang sama
func HandleEmployee(employee entity.Employee) {
	employee.Work()
	employee.Eat()
	employee.TakeLeave()
}
