package main

import (
	"Creacipe/config"
	"Creacipe/controller/homecontroller"
	"Creacipe/controller/koleksicontroller"
	"Creacipe/controller/tentangcontroller"
	"log"
	"net/http"
)

func main() {
	//connect Database
	log.Println("Connecting to the database...")
	config.ConnectDB()
	if config.DB == nil {
		log.Fatal("Failed to connect to the database")
	}
	// 1. HomePage
	http.HandleFunc("/", homecontroller.HomePage)
	// 2. form tambah resep
	http.HandleFunc("/koleksi/add", koleksicontroller.Add)
	// 3. list Koleksi resep
	http.HandleFunc("/koleksi", koleksicontroller.Koleksi)
	// 4. edit resep
	http.HandleFunc("/koleksi/edit", koleksicontroller.Edit)
	// 5. delete resep
	http.HandleFunc("/koleksi/delete", koleksicontroller.Delete)
	// 6. tentang pages
	http.HandleFunc("/tentang", tentangcontroller.Tentang)
	// 7. menghandler untuk url pacth img
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("assets/img/"))))
	// 8. menghandler upload
	http.Handle("/upload/", http.StripPrefix("/upload/", http.FileServer(http.Dir("upload"))))
	// 9. fitur search
	http.HandleFunc("/koleksi/search", koleksicontroller.Search)

	//run server
	log.Println("Server running on port 8080")
	http.ListenAndServe("localhost:8080", nil)
}