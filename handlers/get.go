package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products

func (a *Articles) Get(w http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	//ar := r.Context().Value(KeyProduct{}).(data.Article)

	//a.l.Debug("Inserting product: %#v\n", ar)
	//a.articlesDB.AddArticle(ar)
	a.l.Info("Get article")

	vars := mux.Vars(r)
	a.l.Info("Get article", zap.String("id", vars["id"]))
	/*
		article, ok := data.ArticlesDb.Articles[vars["id"]]
		if !ok {
			http.Error(w, "Article not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(article)
	*/

}

func (a *Articles) GetTagSummary(w http.ResponseWriter, r *http.Request) {
	a.l.Info("Get tag summary")
}
