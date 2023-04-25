package handlers

import (
	"ascii-art-web/asciiart"
	"html/template"
	"net/http"
)

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/index.html"))
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/home.html"))
	if r.Method != "POST" {

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	input := r.FormValue("inputValue")
	banner := r.FormValue("banner")
	Data := asciiart.AsciiArt(input, banner)

	tpl.Execute(w, map[string]string{"Data": Data})

}
