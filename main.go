package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	PORT string
)

func main() {
	flag.StringVar(&PORT, "p", "80", "set service port number.")
	flag.Parse()

	hostname, _ := os.Hostname()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.ParseGlob("*.html"))
		tpl.ExecuteTemplate(w, "index.html", hostname)
	})
	fmt.Println("web service staring at PORT " + PORT + " ...")
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
