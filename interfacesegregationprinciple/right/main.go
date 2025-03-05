package main

import (
	"gosolidprinciple/interfacesegregationprinciple/right/handler"
	"gosolidprinciple/interfacesegregationprinciple/right/usecase"
)

func main() {
	// Data karyawan
	permanentEmployee := usecase.PermanentEmployee{Name: "Alice"}
	freelancer := usecase.Freelancer{Name: "Bob"}
	intern := usecase.Intern{Name: "Charlie"}

	// Memanggil fungsi yang sesuai dengan interface yang diimplementasikan
	handler.HandleWork(permanentEmployee)  // ✅ Bisa bekerja
	handler.HandleEat(permanentEmployee)   // ✅ Bisa makan
	handler.HandleLeave(permanentEmployee) // ✅ Bisa cuti

	handler.HandleWork(freelancer) // ✅ Bisa bekerja
	// handler.HandleEat(freelancer)  // ❌ ERROR: Freelancer tidak bisa makan di kantin
	// handler.HandleLeave(freelancer) // ❌ ERROR: Freelancer tidak bisa cuti

	handler.HandleWork(intern) // ✅ Bisa bekerja
	handler.HandleEat(intern)  // ✅ Bisa makan
	// handler.HandleLeave(intern) // ❌ ERROR: Intern tidak bisa cuti
}
