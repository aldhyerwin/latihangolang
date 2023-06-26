package main

import "fmt"

type Mobil struct {
	Roda       [4]string
	PintuKiri  Pintu
	PintuKanan Pintu
}

type Pintu struct {
	BunyiKetuk string
	BunyiBuka  string
}

func (m *Mobil) GantiRoda(jenisBan string) {
	for i := 0; i < len(m.Roda); i++ {
		m.Roda[i] = jenisBan
	}
}

func main() {
	mobil := Mobil{
		Roda: [4]string{"ban karet", "ban karet", "ban karet", "ban karet"},
		PintuKiri: Pintu{
			BunyiKetuk: "Beep",
			BunyiBuka:  "Knock",
		},
		PintuKanan: Pintu{
			BunyiKetuk: "Knock",
			BunyiBuka:  "Beep",
		},
	}

	mobil.GantiRoda("ban kayu")

	fmt.Println("Jenis Ban Rodanya adalah ", mobil.Roda)
	fmt.Println("Suara Pintu Kiri : ", mobil.PintuKiri)
	fmt.Println("Suara Pintu Kanan : ", mobil.PintuKanan)
}
