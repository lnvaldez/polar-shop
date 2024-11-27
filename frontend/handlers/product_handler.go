package handlers

import (
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

        data := map[string]interface{}{
            "Products": prods,
        }

        err = tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{
            "Content": func() error {
                return tmpl.ExecuteTemplate(w, "product-list-content", data)
            },
        })
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
}

func RenderOrderProductPage(productService *service.ProductService, tmpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        productName := mux.Vars(r)["name"]

        prod, err := productService.GetProductByName(productName)
        if err != nil {
            http.Error(w, "Product not found", http.StatusNotFound)
            return
        }

        data := map[string]interface{}{
            "Product": prod,
        }

        err = tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{
            "Content": func() error {
                return tmpl.ExecuteTemplate(w, "order-product-content", data)
            },
        })
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
}