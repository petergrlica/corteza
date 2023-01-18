package rest

import (
	"github.com/go-chi/chi/v5"

	"github.com/cortezaproject/corteza/server/acme/rest/handlers"
	"github.com/cortezaproject/corteza/server/pkg/auth"
)

func MountRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		// Protect all _private_ routes
		r.Group(func(r chi.Router) {
			r.Use(auth.HttpTokenValidator("api"))

			handlers.NewUsers(Users{}.New()).MountRoutes(r)
		})
	}
}
