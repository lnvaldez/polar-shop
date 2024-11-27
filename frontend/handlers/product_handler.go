package handlers

import (
	"polar-shop/models"
	"polar-shop/service"

	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func RenderAvailableProductsPage(productService *service.ProductService, tmpl *template.Template) http.HandlerFunc {
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

        err = tmpl.ExecuteTemplate(w, "layout", data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}

func RenderOrderProductPage(productService *service.ProductService, tmpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        prodID := mux.Vars(r)["id"]

        prod, err := productService.GetProductById(prodID)
        if err != nil {
            http.Error(w, "Product not found", http.StatusNotFound)
            return
        }

        data := map[string]interface{}{
            "Product": prod,
        }

        err = tmpl.ExecuteTemplate(w, "layout", data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}