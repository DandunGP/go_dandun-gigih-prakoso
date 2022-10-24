/*

BEFORE

package main

type kendaraan struct {
	totalroda       int
	kecepatanperjam int
}

type mobil struct {
	kendaraan
}

func (m *mobil) berjalan() {
	m.tambahkecepatan(10)
}

func (m *mobil) tambahkecepatan(kecepatanbaru int) {
	m.kecepatanperjam = m.kecepatanperjam + kecepatanbaru
}

func main() {
	mobilcepat := mobil{}
	mobilcepat.berjalan()
	mobilcepat.berjalan()
	mobilcepat.berjalan()

	mobillamban := mobil{}
	mobillamban.berjalan()
}

*/

/* AFTER */

package main

import "fmt"

type kendaraan struct {
	totalRoda int
	kecepatan int
}

func (m *kendaraan) berjalan() {
	m.tambahKecepatan(10)
}

func (m *kendaraan) tambahKecepatan(kecepatanbaru int) {
	m.kecepatan = m.kecepatan + kecepatanbaru
}

func main() {
	mobilcepat := kendaraan{}
	mobilcepat.berjalan()
	mobilcepat.berjalan()
	mobilcepat.berjalan()
	fmt.Println(mobilcepat.kecepatan)

	mobillamban := kendaraan{}
	mobillamban.berjalan()
	fmt.Println(mobillamban.kecepatan)
}
