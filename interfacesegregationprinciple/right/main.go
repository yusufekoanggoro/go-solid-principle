package main

import "fmt"

// ✅ Interface ini hanya berisi metode yang relevan untuk manusia.
// Ini mengikuti Interface Segregation Principle (ISP) karena tidak memaksa semua pekerja untuk mengimplementasikan metode yang tidak mereka butuhkan.
type HumanWorker interface {
	Work()
	Eat()
	Sleep()
}

type HumanWorkerImpl struct {
}

// ✅ Mengembalikan interface HumanWorker, bukan struct konkret.
// Dengan cara ini, kita bisa mengganti implementasi dengan mudah tanpa memodifikasi kode yang menggunakannya.
func NewHumanWorker() HumanWorker {
	return HumanWorkerImpl{}
}

func (rh HumanWorkerImpl) Work() {
	fmt.Println("Manusia sedang bekerja...")
}

func (rh HumanWorkerImpl) Eat() {
	fmt.Println("Manusia sedang makan...")
}

func (rh HumanWorkerImpl) Sleep() {
	fmt.Println("Manusia sedang tidur...")
}

// ✅ RobotWorker hanya memiliki metode yang diperlukan untuk robot.
// Ini sesuai dengan ISP karena tidak memaksa robot untuk mengimplementasikan metode Eat() dan Sleep(), yang tidak relevan bagi robot.
type RobotWorker interface {
	Work()
}

type RobotWorkerImpl struct {
}

// ✅ Konstruktor mengembalikan interface RobotWorker, bukan struct konkret.
// Hal ini memungkinkan fleksibilitas dalam mengganti implementasi robot tanpa mengubah kode yang menggunakannya.
func NewRobotWorker() RobotWorker {
	return RobotWorkerImpl{}
}

func (rh RobotWorkerImpl) Work() {
	fmt.Println("Robot sedang kerja...")
}

func main() {
	h := NewHumanWorker()
	h.Work()
	h.Eat()
	h.Sleep()

	r := NewRobotWorker()
	r.Work()
}
