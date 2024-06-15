package middlewares

import (
	"context"
	"log"
	"net/http"

	"github.com/PranitRout07/Practice-Golang/chi_and_htmx/initializers"
	"github.com/PranitRout07/Practice-Golang/chi_and_htmx/models"
	"github.com/go-chi/chi/v5"
)

// Define a new type for the context key
type contextKey string

const PostsKey contextKey = "posts"
var ID string
func PostCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var posts models.Posts
		id := chi.URLParam(r, "id")
		log.Println("URL parameter 'id':", id)
		ID = id
		if id := chi.URLParam(r, "id"); id != "" {
			stmt := "select * from posts where id=$1"
			log.Println(id, "hhhhhhh")
			row := initializers.DBConnection.QueryRow(stmt, id)

			err := row.Scan(&posts.Id, &posts.Title)

			if err != nil {
				log.Println("Error", err)
			}
		}
		
		ctx := context.WithValue(r.Context(), PostsKey, posts)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}