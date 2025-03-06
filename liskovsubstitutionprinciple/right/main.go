package main

import "fmt"

// Bird adalah semua burung
type Bird interface {
	Walk()
}

// Flyable adalah burung yang bisa terbang
type Flyable interface {
	Fly()
}

// Sparrow bisa terbang dan berjalan
type Sparrow struct{}

func (s Sparrow) Walk() {
	fmt.Println("Burung gereja berjalan.")
}

func (s Sparrow) Fly() {
	fmt.Println("Burung gereja terbang rendah.")
}

// Penguin hanya bisa berjalan dan berenang
type Penguin struct{}

func (p Penguin) Walk() {
	fmt.Println("Penguin berjalan di es.")
}

func (p Penguin) Swim() {
	fmt.Println("Penguin berenang di air dingin.")
}

// Fungsi yang hanya menerima burung yang bisa terbang
func makeItFly(f Flyable) {
	f.Fly()
}

func main() {
	sparrow := Sparrow{}
	makeItFly(sparrow) // ✅ Output: "Burung gereja terbang rendah."

	penguin := Penguin{}
	penguin.Walk() // ✅ Output: "Penguin berjalan di es."
	penguin.Swim() // ✅ Output: "Penguin berenang di air dingin."

	// makeItFly(penguin) ❌ Tidak bisa, karena Penguin bukan Flyable
}
