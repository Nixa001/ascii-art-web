package handlers

import (
	"html/template"
	"net/http"
)

// var tpl *template.Template

//	func renderTemplate() {
//		renderTemplate(w, "home")
//	}
func ServerHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.html"))
	tpl.ExecuteTemplate(w, "index.html", nil)
	// renderTemplate(w, "home")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.html"))

	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	input := r.FormValue("inputValue")
	d := struct {
		inputValue string
	}{
		inputValue: input,
	}
	tpl.ExecuteTemplate(w, "home.html", d)

}

// renderTemplate(w, "home")

// func renderTemplate(w http.ResponseWriter, tmpl string) {
// 	t, err := template.ParseFiles("./templates/" + tmpl + ".html")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	t.Execute(w, nil)
// }
