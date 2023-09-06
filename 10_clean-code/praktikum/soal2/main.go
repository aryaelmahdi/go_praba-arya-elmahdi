package main

type Kendaraan struct {
	Roda      int
	Kecepatan int
}

type Mobil struct {
	Kendaraan
}

func (m *Mobil) berjalan(kecepatanbaru int) {
	m.tambahKecepatan(kecepatanbaru)
}

func (m *Mobil) tambahKecepatan(kecepatanbaru int) {
	m.Kecepatan = m.Kecepatan + kecepatanbaru
}

func main() {
	mobilCepat := Mobil{}
	mobilCepat.berjalan(10)
	mobilCepat.berjalan(10)
	mobilCepat.berjalan(10)

	mobilLamban := Mobil{}
	mobilLamban.berjalan(5)
}
