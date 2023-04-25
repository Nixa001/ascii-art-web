package handlers

import (
	"ascii-art-web/asciiart"
	"fmt"
	"html/template"
	"net/http"
)

// var tpl *template.Template

//	func renderTemplate() {
//		renderTemplate(w, "home")
//	}
func ServerHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/index.html"))
	tpl.ExecuteTemplate(w, "index.html", nil)
	// renderTemplate(w, "home")
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
	fmt.Println(Data)
	fmt.Println("Data")
	// d := struct {
	// 	inputValue string
	// }{
	// 	inputValue: input,
	// }
	tpl.ExecuteTemplate(w, "home.html", nil)

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
