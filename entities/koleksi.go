package entities

import (
	"strings"
	"time"
)

type Koleksi struct {
	Id        int
	Judul     string
	Bahan     string
	Langkah   string
	Gambar    string
	JenisMasakanID      int
	TingkatKesulitanID  int
	CreatedAt time.Time
	UpdatedAt time.Time
	JenisMasakan JenisMasakan
	TingkatKesulitan TingkatKesulitan
}
type JenisMasakan struct {
	Id        int
	NamaJenis string
}

type TingkatKesulitan struct {
	Id        int
	NamaLevel string
}

// Method to split Bahan agar ke bawah 
func (k Koleksi) BahanList() []string {
	return strings.Split(k.Bahan, "\n")
}

// Method to split Langkah agar ke bawah
func (k Koleksi) LangkahList() []string {
	return strings.Split(k.Langkah, "\n")
}