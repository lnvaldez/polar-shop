package handlers

import (
	"polar-shop/service"
	"polar-shop/models"

	"net/http"
	"html/template"
)

func ProductListHandler(productService *service.ProductService, tmpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        prods, err := productService.GetAvailableProducts()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        data := struct {
            Products []model.Product
        }{
            Products: prods,
        }

        err = tmpl.ExecuteTemplate(w, "product_list.html", data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}