package entity

// Workable → Hanya untuk karyawan yang bekerja
type Workable interface {
	Work() error
}

// Eatable → Hanya untuk karyawan yang dapat makan di kantin
type Eatable interface {
	Eat() error
}

// Leaveable → Hanya untuk karyawan yang bisa mengambil cuti
type Leaveable interface {
	TakeLeave() error
}
