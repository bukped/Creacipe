package koleksicontroller

import (
	"Creacipe/entities"
	"Creacipe/models/koleksimodel"
	"html/template"
	"net/http"
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
		koleksi := koleksimodel.Detail(id)
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
		koleksi.Judul = r.FormValue("judul")
		koleksi.Bahan = r.FormValue("bahan")
		koleksi.Langkah = r.FormValue("langkah")
		koleksi.UpdatedAt = time.Now()
		if ok := koleksimodel.Update(id, koleksi); !ok {
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

