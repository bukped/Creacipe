package koleksicontroller

import (
	"Creacipe/entities"
	"Creacipe/models/koleksimodel"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

//fungsi untuk menampilkan halaman koleksi
func Koleksi(w http.ResponseWriter, r *http.Request) {
	koleksiresep := koleksimodel.GetAll()
	data := map[string]any{
		"koleksiresep": koleksiresep,
	}
	temp, err := template.ParseFiles("views/koleksiresep/koleksi.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}
//fungsi untuk menambah resep
func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/form/form.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}
	if r.Method == "POST" {
		var koleksi entities.Koleksi
		koleksi.Judul = r.FormValue("judul")
		koleksi.Bahan = r.FormValue("bahan")
		koleksi.Langkah = r.FormValue("langkah")
		koleksi.JenisMasakanID, _ = strconv.Atoi(r.FormValue("jenis_masakan_id"))
		koleksi.TingkatKesulitanID, _ = strconv.Atoi(r.FormValue("tingkat_kesulitan_id"))
		koleksi.Gambar = uploadFile(r)
		koleksi.CreatedAt = time.Now()
		koleksi.UpdatedAt = time.Now()

		ok := koleksimodel.Create(koleksi)
		if !ok {
			temp, _ := template.ParseFiles("views/form/form.html")
			temp.Execute(w, nil)
		}
		http.Redirect(w, r, "/koleksi", http.StatusSeeOther)
	}
}
//fungsi untuk edit resep
func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/form/edit.html")
		if err != nil {
			panic(err)
		}
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		koleksi, ok := koleksimodel.Detail(id)
		if !ok {
			http.Error(w, "Koleksi not found", http.StatusNotFound)
			return
		}
		data := map[string]any{
			"koleksi": koleksi,
		}
		temp.Execute(w, data)
	}
	if r.Method == "POST" {
		var koleksi entities.Koleksi
		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		koleksi.Id = id
		koleksi.Judul = r.FormValue("judul")
		koleksi.Bahan = r.FormValue("bahan")
		koleksi.Langkah = r.FormValue("langkah")
		koleksi.JenisMasakanID, _ = strconv.Atoi(r.FormValue("jenis_masakan_id"))
		koleksi.TingkatKesulitanID, _ = strconv.Atoi(r.FormValue("tingkat_kesulitan_id"))
		koleksi.UpdatedAt = time.Now()

		// Upload gambar
		gambar := uploadFile(r)
		if gambar != "" {
			koleksi.Gambar = gambar
		} else {
			// Jika tidak ada gambar baru yang diunggah, gunakan gambar yang lama
			oldKoleksi, _ := koleksimodel.Detail(id)
			koleksi.Gambar = oldKoleksi.Gambar
		}

		if ok := koleksimodel.Update(koleksi); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}
		http.Redirect(w, r, "/koleksi", http.StatusSeeOther)
	}
}
//fungsi untuk delete resep
func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}
	if err := koleksimodel.Delete(id); err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/koleksi", http.StatusSeeOther)
}

//fungsi search
func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
			query := r.URL.Query().Get("q")
			koleksiresep := koleksimodel.Search(query)

			data := map[string]interface{}{
					"koleksiresep": koleksiresep,
					"Query": query,
			}

			temp, err := template.ParseFiles("views/koleksiresep/koleksi.html")
			if err != nil {
					panic(err)
			}
			temp.Execute(w, data)
	}
}

//fungsi untuk menangani upload file
func uploadFile(r *http.Request) string {
	file, handler, err := r.FormFile("gambar")
	if err != nil {
		return ""
	}
	defer file.Close()

	// Get the file extension and validate it
	ext := filepath.Ext(handler.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" && ext != ".bmp" && ext != ".tiff" && ext != ".avif" && ext != ".webp" {
		return ""
	}

	// Generate a unique file name to prevent overwriting
	filename := filepath.Join("upload", handler.Filename)
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return ""
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return ""
	}

	return handler.Filename
}

