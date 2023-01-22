package server

import (
	"context"
	"net/http"
	"webapp/helper"
)

func (s *Server) Authorize(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Println("auth middleware")

		sess, err := s.session.Get(r, "session")
		if err != nil {
			helper.InternalServerError(w)
			return
		}
		username, ok := sess.Values["user_id"]
		if !ok {
			s.logger.Printf("user %s auth failed, redirecting to /login", username)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		ctx := r.Context()
		valueCtx := context.WithValue(ctx, "user", username)
		r = r.Clone(valueCtx)

		s.logger.Printf("user %s auth successful", username)
		handler.ServeHTTP(w, r)
	}
}
