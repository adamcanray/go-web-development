package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	// fmt.Println("hello dunia") // cetak ke command line
	mux := http.NewServeMux()

	// closure
	// aboutHandler := func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("About page"))
	// }

	// routes
	mux.HandleFunc("/", handler.HomeHandler) // root route, pattern yang lain yang tidak terdaftar akan lari kesini
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/mario", handler.MarioHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	// mux.HandleFunc("/about", aboutHandler)
	// mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Profile Page"))
	// })
	mux.HandleFunc("/post-get", handler.PostGet)
	mux.HandleFunc("/form", handler.Form)
	mux.HandleFunc("/process", handler.Process)

	// use static files
	fileServer := http.FileServer(http.Dir("assets"))
	// mux.Handle("/static/", fileServer) // will load -> /assets/static/style.css
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) // will load -> /assets/style.css

	log.Println("Starting web on port 8090")

	err := http.ListenAndServe(":8090", mux)
	log.Fatal(err)
}
