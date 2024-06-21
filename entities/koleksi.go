package entities

import (
	"strings"
	"time"
)

type Koleksi struct {
	Id        uint
	Judul     string
	Bahan     string
	Langkah   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Method to split Bahan agar ke bawah 
func (k Koleksi) BahanList() []string {
	return strings.Split(k.Bahan, "\n")
}

// Method to split Langkah agar ke bawah
func (k Koleksi) LangkahList() []string {
	return strings.Split(k.Langkah, "\n")
}