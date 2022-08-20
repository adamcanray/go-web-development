package handler

import (
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// pattern yang tidak terdaftar, selain / akan lari ke 404
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("home nich"))

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
		return
	}

	// map dengan tipe valuenya interface kosong, akan bisa diisi dengan tipe apa saja.
	// data := map[string]interface{}{
	// 	"title":   "I'm learning golang web",
	// 	"content": "I'am learning golang web with Agung Setiawan",
	// 	"id":      1,
	// }
	// data := entity.Product{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 3}
	data := []entity.Product{
		{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 3},
		{ID: 1, Name: "Fortuner", Price: 420000000, Stock: 8},
		{ID: 1, Name: "Pajero", Price: 420000000, Stock: 10},
		{ID: 1, Name: "Yaris", Price: 320000000, Stock: 1},
	}

	err = tmpl.Execute(w, data) // execute template, menampilkan template html
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world, saya sedang belajar golang web."))
}

func MarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mario from Nintendo"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	// get query string
	id := r.URL.Query().Get("id")

	// convert string to int
	idNumb, err := strconv.Atoi(id)

	// validasi id harus angka atau id tidak kurang dari 1
	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	// fmt.Fprintf(w, "Product page: %d", idNumb)

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
		return
	}

	// map dengan tipe valuenya interface kosong, akan bisa diisi dengan tipe apa saja.
	data := map[string]interface{}{
		"title": "Product",
		"id":    idNumb,
	}

	err = tmpl.Execute(w, data) // execute template, menampilkan template html
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method // GET/POST/etc

	switch method {
	case "GET":
		w.Write([]byte("Ini adalah GET"))
	case "POST":
		w.Write([]byte("Ini adalah POST"))
	default:
		http.Error(w, "Error is happening, keep calm.", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil) // execute template, menampilkan template html
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
			return
		}
		return
	}

	http.Error(w, "Error is happening, keep calm.", http.StatusBadRequest)
}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm() // mengolah form
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		message := r.Form.Get("message")

		data := map[string]interface{}{"name": name, "message": message}

		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data) // execute template, menampilkan template html
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
			return
		}

		return
	}
	http.Error(w, "Error is happening, keep calm.", http.StatusBadRequest)
}
