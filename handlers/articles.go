package handlers

import (
	"go.uber.org/zap"
)

type Articles struct {
	l *zap.Logger
}

func NewArticles(l *zap.Logger) *Articles {
	return &Articles{l}
}

/*
// getArticleByID returns an article given its ID
func (a *Articles) getArticleByID(id int) (Article, error) {
}

}


func (a *Articles) getRelatedTagsAndCount(tag string, date string) (int, []string, error) {

}
func (a *Articles) handlePostArticles(w http.ResponseWriter, r *http.Request) {
}
*/
