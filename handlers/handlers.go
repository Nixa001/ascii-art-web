package handlers

import (
	"ascii-art-web/asciiart"
	"html/template"
	"net/http"
)

var Output string
var TextErr string
var NumErr int
var errors = true

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/index.html"))
	if r.URL.Path != "/" {
		TextErr = "Page Not Found"
		NumErr=http.StatusNotFound
		w.WriteHeader(NumErr)
		errorPage(w, r, map[string]interface{}{"Output": Output, "TextErr": TextErr, "NumErr": http.StatusNotFound})
		return
	}
	tpl.Execute(w, nil)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/home.html"))
	// if r.Method == "GET" {
	// 	http.Redirect(w, r, "/", http.StatusSeeOther)
	// 	return
	// } else if 
	if r.Method == "POST" {

		input := r.FormValue("inputValue")
		banner := r.FormValue("banner")
		Output, errors = asciiart.AsciiArt(input, banner)
		TextErr = "OK"
		NumErr = http.StatusOK
		if errors {
			TextErr = "Bad Request"
			NumErr = http.StatusBadRequest
			errOutput := Output
			w.WriteHeader(NumErr)
			tpl.Execute(w, map[string]string{"errOutput": errOutput})
			return
			// errorPage(w, r, map[string]interface{}{"Output": Output, "TextErr": TextErr, "NumErr": NumErr})
			// return
		}
		tpl.Execute(w, map[string]string{"Output": Output})
		return
	} else {
		TextErr = "Internal Server Error"
		NumErr = http.StatusInternalServerError
		errorPage(w, r, map[string]interface{}{"Output": Output, "TextErr": TextErr, "NumErr": NumErr})
		return
	}

}
func errorPage(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	tpl := template.Must(template.ParseFiles("templates/error.html"))
	tpl.Execute(w, data)
}
