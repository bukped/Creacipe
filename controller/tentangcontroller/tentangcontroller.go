package tentangcontroller

import (
	"html/template"
	"net/http"
)
//fungsi untuk melihat halaman tentang
func Tentang(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/tentang/tentang.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}