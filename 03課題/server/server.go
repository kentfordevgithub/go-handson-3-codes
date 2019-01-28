package main

import (
	"html/template"
	"log"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパース
	t := template.Must(template.ParseFiles("./hello.html.tpl"))
	// テンプレートを描画
	s := "Hello, world."
	if err := t.ExecuteTemplate(w, "hello.html.tpl", s); err != nil {
		log.Fatal(err)
	}
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパース
	t := template.Must(template.ParseFiles("./form.html.tpl"))
	// テンプレートを描画
	s := "Hello, world."
	if err := t.ExecuteTemplate(w, "form.html.tpl", s); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/form", handleForm)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
