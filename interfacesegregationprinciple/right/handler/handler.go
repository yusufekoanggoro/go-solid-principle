package handler

import "gosolidprinciple/interfacesegregationprinciple/right/entity"

// HandleWork → Menjalankan fungsi kerja karyawan
func HandleWork(worker entity.Workable) {
	worker.Work()
}

// HandleEat → Hanya bisa menerima karyawan yang bisa makan
func HandleEat(eater entity.Eatable) {
	eater.Eat()
}

// HandleLeave → Hanya bisa menerima karyawan yang bisa cuti
func HandleLeave(leaveTaker entity.Leaveable) {
	leaveTaker.TakeLeave()
}
