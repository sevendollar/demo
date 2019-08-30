package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	hostname, _ := os.Hostname()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.ParseGlob("*.html"))
		tpl.ExecuteTemplate(w, "index.html", hostname)
	})
	fmt.Println("service staring...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
