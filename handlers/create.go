package handlers

import (
	"net/http"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products

func (a *Articles) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	//ar := r.Context().Value(KeyProduct{}).(data.Article)

	//a.l.Debug("Inserting product: %#v\n", ar)
	//a.articlesDB.AddArticle(ar)
	a.l.Info("Create article")
}
