package koleksimodel

import (
	"Creacipe/config"
	"Creacipe/entities"
	"log"
)

//fungsi untuk mengambil semua data
func GetAll() []entities.Koleksi {
	rows, err := config.DB.Query(`
		SELECT k.id, k.judul, k.bahan, k.langkah, k.gambar, k.jenis_masakan_id, k.tingkat_kesulitan_id, k.created_at, k.updated_at, jm.nama_jenis, tk.nama_level FROM bukuresep k
		LEFT JOIN jenis_masakan jm ON k.jenis_masakan_id = jm.id
		LEFT JOIN tingkat_kesulitan tk ON k.tingkat_kesulitan_id = tk.id
	`)
	if err != nil {
		log.Println("Error executing query:", err)
		panic(err)
	}
	defer rows.Close()

	var koleksiresep []entities.Koleksi
	for rows.Next() {
		var koleksi entities.Koleksi
		var jenisMasakan entities.JenisMasakan
		var tingkatKesulitan entities.TingkatKesulitan

		err := rows.Scan(&koleksi.Id, &koleksi.Judul, &koleksi.Bahan, &koleksi.Langkah, &koleksi.Gambar, &koleksi.JenisMasakanID, &koleksi.TingkatKesulitanID, &koleksi.CreatedAt, &koleksi.UpdatedAt, &jenisMasakan.NamaJenis, &tingkatKesulitan.NamaLevel)
		if err != nil {
			log.Println("Error scanning row:", err)
			panic(err)
		}

		koleksi.JenisMasakan = jenisMasakan
		koleksi.TingkatKesulitan = tingkatKesulitan
		koleksiresep = append(koleksiresep, koleksi)
	}
	return koleksiresep
}

//fungsi untuk menambah data
func Create(koleksi entities.Koleksi) bool {
	_, err := config.DB.Exec(`
		INSERT INTO bukuresep (judul, bahan, langkah, gambar, jenis_masakan_id, tingkat_kesulitan_id, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`, koleksi.Judul,
	koleksi.Bahan, 
	koleksi.Langkah, 
	koleksi.Gambar, 
	koleksi.JenisMasakanID, 
	koleksi.TingkatKesulitanID, 
	koleksi.CreatedAt, 
	koleksi.UpdatedAt)

	if err != nil {
		log.Println("Error executing query:", err)
		return false
	}
	return true
}

//mengambil data berdasarkan id 
func Detail(id int) (entities.Koleksi, bool) {
	var koleksi entities.Koleksi
	var jenisMasakan entities.JenisMasakan
	var tingkatKesulitan entities.TingkatKesulitan

	row := config.DB.QueryRow(`
		SELECT k.id, k.judul, k.bahan, k.langkah, k.gambar, k.jenis_masakan_id, k.tingkat_kesulitan_id, k.created_at, k.updated_at,
		jm.nama_jenis, tk.nama_level
		FROM bukuresep k
		LEFT JOIN jenis_masakan jm ON k.jenis_masakan_id = jm.id
		LEFT JOIN tingkat_kesulitan tk ON k.tingkat_kesulitan_id = tk.id
		WHERE k.id = ?
	`, id)

	err := row.Scan(&koleksi.Id, &koleksi.Judul, &koleksi.Bahan, &koleksi.Langkah, &koleksi.Gambar, &koleksi.JenisMasakanID, &koleksi.TingkatKesulitanID, &koleksi.CreatedAt, &koleksi.UpdatedAt, &jenisMasakan.NamaJenis, &tingkatKesulitan.NamaLevel)
	if err != nil {
		log.Println("Error scanning row:", err)
		return koleksi, false
	}

	koleksi.JenisMasakan = jenisMasakan
	koleksi.TingkatKesulitan = tingkatKesulitan
	return koleksi, true
}

//fungsi untuk update
func Update(koleksi entities.Koleksi) bool {
	_, err := config.DB.Exec(`
		UPDATE bukuresep
		SET judul = ?, bahan = ?, langkah = ?, gambar = ?, jenis_masakan_id = ?, tingkat_kesulitan_id = ?, updated_at = ?
		WHERE id = ?
	`, koleksi.Judul, koleksi.Bahan, koleksi.Langkah, koleksi.Gambar, koleksi.JenisMasakanID, koleksi.TingkatKesulitanID, koleksi.UpdatedAt, koleksi.Id)

	if err != nil {
		log.Println("Error executing query:", err)
		return false
	}
	return true
}

//fungsi untuk delete
func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM bukuresep WHERE id = ?", id)
	return err
}

//fungsi search
func Search(query string) []entities.Koleksi {
	rows, err := config.DB.Query(`
			SELECT k.id, k.judul, k.bahan, k.langkah, k.gambar, k.jenis_masakan_id, k.tingkat_kesulitan_id, k.created_at, k.updated_at, jm.nama_jenis, tk.nama_level
			FROM bukuresep k
			LEFT JOIN jenis_masakan jm ON k.jenis_masakan_id = jm.id
			LEFT JOIN tingkat_kesulitan tk ON k.tingkat_kesulitan_id = tk.id
			WHERE LOWER(k.judul) LIKE ? OR LOWER(k.bahan) LIKE ? OR LOWER(jm.nama_jenis) LIKE ? OR LOWER(tk.nama_level) LIKE ?
	`, "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%")
	if err != nil {
			log.Println("Error executing query:", err)
			panic(err)
	}
	defer rows.Close()

	var koleksiresep []entities.Koleksi

	for rows.Next() {
			var koleksi entities.Koleksi
			var jenisMasakan entities.JenisMasakan
			var tingkatKesulitan entities.TingkatKesulitan
			err := rows.Scan(&koleksi.Id, &koleksi.Judul, &koleksi.Bahan, &koleksi.Langkah, &koleksi.Gambar, &koleksi.JenisMasakanID, &koleksi.TingkatKesulitanID, &koleksi.CreatedAt, &koleksi.UpdatedAt, &jenisMasakan.NamaJenis, &tingkatKesulitan.NamaLevel)
			if err != nil {
					log.Println("Error scanning row:", err)
					panic(err)
			}
			koleksi.JenisMasakan = jenisMasakan
			koleksi.TingkatKesulitan = tingkatKesulitan
			koleksiresep = append(koleksiresep, koleksi)
	}

	return koleksiresep
}

