package handlers

import (
	"ascii-art-web/asciiart"
	"html/template"
	"net/http"
)

var Output string
var TextErr string
var errNum int = http.StatusOK

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/index.html"))
	if r.URL.Path != "/" {
		TextErr = "Page Not Found"
		errNum = http.StatusNotFound
		w.WriteHeader(errNum)
		errorPage(w, r, map[string]interface{}{"Output": Output, "TextErr": TextErr, "errNum": http.StatusNotFound})
		return
	}
	w.WriteHeader(errNum)
	tpl.Execute(w, nil)
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/home.html"))
	if r.Method == "POST" {

		input := r.FormValue("inputValue")
		banner := r.FormValue("banner")
		btn := r.FormValue("btn")
		if btn == "about" {
			tmpl := template.Must(template.ParseFiles("templates/index.html"))
			title := "About"
			title2 := "How to use it ?"
			text := "Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of the project ascii-art"
			tmpl.Execute(w, map[string]string{"title": title, "text": text, "title2": title2})
			return

		} else if btn == "start" {
			tpl.Execute(w, nil)
			return
		}
		Output, errNum = asciiart.AsciiArt(input, banner)
		TextErr = "OK"
		if errNum == http.StatusBadRequest {
			TextErr = "Bad Request"
			w.WriteHeader(errNum)
			errorPage(w, r, map[string]interface{}{"Output": Output, "TextErr": TextErr, "errNum": errNum})
			return
		} else if errNum == http.StatusInternalServerError {
			TextErr = "Internal Server Error"
			w.WriteHeader(errNum)
			errorPage(w, r, map[string]interface{}{"Output": Output, "TextErr": TextErr, "errNum": errNum})
			return

		}
		tpl.Execute(w, map[string]string{"Output": Output})
		return
	}

}
func errorPage(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	tpl := template.Must(template.ParseFiles("templates/error.html"))
	tpl.Execute(w, data)
}
