package handlers

import (
	"net/http"
	"html/template"
)

func RenderRegisterPage(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "register.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return 
		}
	}
}

func RenderLoginPage(tmpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        err := tmpl.ExecuteTemplate(w, "login.html", nil)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
}