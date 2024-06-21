package koleksimodel

import (
	"Creacipe/config"
	"Creacipe/entities"
	"log"
)

//fungsi untuk mengambil semua data
func GetAll() []entities.Koleksi {
	rows, err := config.DB.Query("SELECT * FROM bukuresep")
	if err != nil {
		log.Println("Error executing query:",err)
		panic(err)
	}
	defer rows.Close()
	//buat variable untuk menampung seluruh data koleksi yang type nya struct
	var koleksiresep []entities.Koleksi
	//looping rows
	for rows.Next() {
		var koleksi entities.Koleksi
		//memindahkan nilai dari rows ke var koleksi dengan fungsi scan
		err := rows.Scan(&koleksi.Id, &koleksi.Judul, &koleksi.Bahan, &koleksi.Langkah, &koleksi.CreatedAt, &koleksi.UpdatedAt)
		if err != nil {
			log.Println("Error scanning row:", err)
			panic(err)
		}
		//memindahkan nilai dari var koleksi ke var koleksi resep
		koleksiresep = append(koleksiresep, koleksi)
	}
	return koleksiresep
}

//fungsi untuk menambah data
func Create(koleksi entities.Koleksi) bool {
	result, err := config.DB.Exec(
		"INSERT INTO bukuresep (judul, bahan, langkah, created_at, updated_at) VALUES (?, ?, ?, ?,?)",
		koleksi.Judul,
		koleksi.Bahan,
		koleksi.Langkah,
		koleksi.CreatedAt,
		koleksi.UpdatedAt,
	)
	if err != nil {
		panic(err)
	}
	//mengambil id terakhir
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return lastInsertId > 0
}

//mengambil data berdasarkan id 
func Detail(id int) entities.Koleksi {
	row := config.DB.QueryRow(`SELECT id, judul, bahan, langkah FROM bukuresep WHERE id = ? `, id)
	var koleksi entities.Koleksi
	if err := row.Scan(&koleksi.Id, &koleksi.Judul, &koleksi.Bahan, &koleksi.Langkah); err != nil {
		panic(err.Error())
	}
	return koleksi
}

//fungsi untuk update
func Update(id int, koleksi entities.Koleksi) bool {
	query, err := config.DB.Exec(`UPDATE bukuresep SET judul = ?, bahan = ?, langkah = ? , updated_at = ? where id = ?`, koleksi.Judul, koleksi.Bahan, koleksi.Langkah, koleksi.UpdatedAt, id)
	if err != nil {
		panic(err)
	}
	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}
	return result > 0
}

//fungsi untuk delete
func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM bukuresep WHERE id = ?", id)
	return err
}