package main

import "fmt"

// Interface ini melanggar Interface Segregation Principle (ISP) karena menggabungkan
// beberapa fungsi yang tidak semua implementasi butuhkan. Robot tidak perlu Eat() dan Sleep().
type Worker interface {
	Work()
	Eat()
	Sleep()
}

type HumanWorker struct{}

func NewHumanWorker() Worker {
	return HumanWorker{}
}

func (h HumanWorker) Work() {
	fmt.Println("Manusia sedang kerja...")
}

func (h HumanWorker) Eat() {
	fmt.Println("Manusia sedang makan...")
}

func (h HumanWorker) Sleep() {
	fmt.Println("Manusia sedang tidur...")
}

type RobotWorker struct{}

func NewRobotWorker() Worker {
	return RobotWorker{}
}

func (r RobotWorker) Work() {
	fmt.Println("Robot sedang bekerja...")
}

// Robot dipaksa mengimplementasikan metode yang tidak perlu
func (r RobotWorker) Eat() {
	// Robot tidak makan, tapi tetap harus mengimplementasikan metode ini
	fmt.Println("Robot tidak bisa makan...")
}

func (r RobotWorker) Sleep() {
	// Robot tidak tidur, tapi tetap harus mengimplementasikan metode ini
	fmt.Println("Robot tidak bisa tidur...")
}

func main() {
	h := NewHumanWorker()
	h.Work()
	h.Eat()
	h.Sleep()

	r := NewRobotWorker()
	r.Work()
	r.Eat()
	r.Sleep()
}
