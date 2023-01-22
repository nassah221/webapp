package server

import (
	"net/http"
	"webapp/helper"
)

func (s *Server) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	s.logger.Println("logout GET handler")

	if err := s.session.DeleteCurrentSession(w, r); err != nil {
		helper.InternalServerError(w)
		return
	}

	s.logger.Println("redirecting to /login")
	http.Redirect(w, r, "/login", http.StatusFound)
}
