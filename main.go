package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

func add_user(isimsoyisim string, telefon string, eposta string) bool {

	db, err := sql.Open("mysql", "root:1Mhszxisq4r@tcp(127.0.0.1:3306)/deneme1")
	if err != nil {
		panic(err)
	}
	add, err := db.Query("INSERT INTO telefondefter (isimsoyisim,telefon,eposta) VALUES (?,?,?)", (isimsoyisim), (telefon), (eposta))
	if err != nil {
		panic(err)
	}

	fmt.Println(add)
	defer db.Close()
	return true
}

func kayitekle1(w http.ResponseWriter, r *http.Request) {
	var tmplt = template.Must(template.ParseFiles("html/index.html"))
	tmplt.Execute(w, nil)
}

func kayitekle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var isimsoyisim = r.Form["isimsoyisim"]
	var telefon = r.Form["telefon"]
	var eposta = r.Form["eposta"]
	fmt.Println(isimsoyisim, telefon, eposta)
	if add_user(isimsoyisim[0], telefon[0], eposta[0]) {
		var tmplt = template.Must(template.ParseFiles("html/index3.html"))
		tmplt.Execute(w, nil)
	} else {
		var tmplt = template.Must(template.ParseFiles("html/error.html"))
		tmplt.Execute(w, nil)
	}
}

var id string
var isimsoyisim string
var telefon string
var eposta string

var db *sql.DB

func listele(HTMLcevap http.ResponseWriter,
	HTMListek *http.Request) {
	var tmplt = template.Must(template.ParseFiles("html/index2.html"))
	// sorgu cevabı kesite aktarılsın:
	Satirlar, _ := db.Query("SELECT * FROM telefondefter")
	defer Satirlar.Close()
	Kitaplar := make([]*Kitap, 0)
	for Satirlar.Next() {
		satir := new(Kitap)
		Satirlar.Scan(&satir.id, &satir.isimsoyisim, &satir.telefon, &satir.eposta)
		Kitaplar = append(Kitaplar, satir)
		tmplt.Execute(HTMLcevap, nil)

	}
	// HTML cevabını oluşturalım:
	var _ = template.Must(template.ParseFiles("html/index2.html"))
	for i, kitap := range Kitaplar {
		fmt.Fprintf(HTMLcevap, "%d, %s- %s - %s- %s\n",
			i, kitap.id, kitap.isimsoyisim, kitap.telefon, kitap.eposta)
		log.Println(kitap.id, ":", kitap.isimsoyisim, ":", kitap.telefon, ":", kitap.eposta)

	}
}
func main() {
	http.Handle("/scss/", http.StripPrefix("/scss/", http.FileServer(http.Dir("scss"))))
	http.Handle("/src/", http.StripPrefix("/src/", http.FileServer(http.Dir("src"))))
	http.Handle("/rtl/", http.StripPrefix("/rtl/", http.FileServer(http.Dir("rtl"))))
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("html"))))
	http.Handle("/layouts/", http.StripPrefix("/layouts/", http.FileServer(http.Dir("layouts"))))
	http.Handle("/tmp/", http.StripPrefix("/tmp/", http.FileServer(http.Dir("tmp"))))

	fmt.Println("Server Başlatılıyor")

	http.HandleFunc("/", kayitekle1)
	http.HandleFunc("/kayitekle", kayitekle)
	http.HandleFunc("/kayitlist", listele)
	http.ListenAndServe(":8999", nil)

}
