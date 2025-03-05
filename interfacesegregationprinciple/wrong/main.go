package main

import (
	"gosolidprinciple/interfacesegregationprinciple/wrong/handler"
	"gosolidprinciple/interfacesegregationprinciple/wrong/usecase"
)

func main() {
	// ❌ Semua tipe karyawan menggunakan interface Employee yang salah
	permanentEmployee := usecase.PermanentEmployee{Name: "Alice"}
	freelancer := usecase.Freelancer{Name: "Bob"}
	intern := usecase.Intern{Name: "Charlie"}

	// ❌ Memanggil handler yang salah
	handler.HandleEmployee(permanentEmployee) // ✅ Berjalan normal
	handler.HandleEmployee(freelancer)        // ❌ ERROR saat memanggil Eat() & TakeLeave()
	handler.HandleEmployee(intern)            // ❌ ERROR saat memanggil TakeLeave()
}
