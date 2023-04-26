package handlers

import (
	"ascii-art-web/asciiart"
	"html/template"
	"net/http"
)

type Data struct {
	Output  string
	NumErr  int
	TextErr string
}

var errors bool
var data = Data{}

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/index.html"))
	if r.URL.Path != "/" {
		data.NumErr = 404
		data.TextErr = "Page Not Found"
		errorPage(w, r, map[string]interface{}{"Output": data.Output, "TextErr": data.TextErr, "NumErr": data.NumErr})
		return
	}
	tpl.Execute(w, nil)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/home.html"))
	if r.Method == "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else if r.Method == "POST" {

		input := r.FormValue("inputValue")
		banner := r.FormValue("banner")
		data.Output, errors = asciiart.AsciiArt(input, banner)
		data.TextErr = "OK"
		data.NumErr = 200
		if !errors {
			data.TextErr = "Internal Server Error"
			data.NumErr = 500
			errorPage(w, r, map[string]interface{}{"Output": data.Output, "TextErr": data.TextErr, "NumErr": data.NumErr})
			return
		}
		tpl.Execute(w, map[string]string{"Output": data.Output})
		return
	} else {
		data.TextErr = "Bad Request"
		data.NumErr = 400
		errorPage(w, r, map[string]interface{}{"Output": data.Output, "TextErr": data.TextErr, "NumErr": data.NumErr})
		return
	}

}
func errorPage(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	tpl := template.Must(template.ParseFiles("templates/error.html"))
	tpl.Execute(w, data)
}
