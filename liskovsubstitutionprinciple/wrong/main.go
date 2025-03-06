package main

import "fmt"

// Bird adalah interface yang mengharuskan burung bisa terbang
type Bird interface {
	Fly()
}

// Sparrow bisa terbang (benar)
type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Burung gereja terbang rendah.")
}

// Penguin tidak bisa terbang, tetapi tetap mengimplementasikan Bird (salah)
type Penguin struct{}

func (p Penguin) Fly() {
	panic("Penguin tidak bisa terbang!") // ❌ Ini melanggar LSP!
}

// Fungsi yang mengharapkan burung bisa terbang
func makeBirdFly(b Bird) {
	b.Fly()
}

func main() {
	var bird Bird = Penguin{}
	makeBirdFly(bird) // ❌ Akan panic: "Penguin tidak bisa terbang!"
}
