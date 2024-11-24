package handlers

import (
	"net/http"
	"html/template"
)

func RenderLoginPage(tmpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        err := tmpl.ExecuteTemplate(w, "login.html", nil)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
}