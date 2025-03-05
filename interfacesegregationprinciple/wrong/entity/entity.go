package entity

// ❌ Satu interface besar yang mencakup semua metode
type Employee interface {
	Work() error
	Eat() error
	TakeLeave() error
}
