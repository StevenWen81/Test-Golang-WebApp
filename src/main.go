// main.go
package main

// Mengimport Library
import(
	"net/http"
	"html/template"
)

// Deklarasi variabel
var templates *template.Template

// Fungsi utama
func main() {
	templates = template.Must(template.ParseGlob("file/*.html"))
	// file/*.html => nama_folder/*.nama_ext

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil) 
	// ":XXXX" => Alamat localhost
}

// Fungsi pembantu untuk mengakses index.html
func indexHandler(w  http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil) // nama sesuai dengan file yang ada
}