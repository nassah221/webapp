package server

import (
	"html/template"
	"net/http"
	"webapp/helper"
	"webapp/model"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) RegisterGETHandler(w http.ResponseWriter, _ *http.Request) {
	s.logger.Println("register GET handler")
	helper.ExecuteTemplate(w, "register.html", nil)
}

func (s *Server) RegisterPOSTHandler(w http.ResponseWriter, r *http.Request) {
	s.logger.Println("register POST handler")

	// a production app probably has a validation on these fields
	username := template.HTMLEscapeString(r.PostFormValue("username"))
	password := template.HTMLEscapeString(r.PostFormValue("password"))

	// basic form validation
	if username == "" || password == "" {
		helper.ExecuteTemplate(w, "login.html", "username and password cannot be empty")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		helper.InternalServerError(w)
		return
	}

	if err := s.db.CreateUser(model.User{ //nolint
		ID:       uuid.NewString(),
		Username: username,
		Password: hash,
	}); err != nil {
		s.logger.Printf("user %s failed to register %v", username, err)
		helper.ExecuteTemplate(w, "register.html", "username is already taken")
		return
	}

	// delete current session if any when new user registered
	if err := s.session.DeleteCurrentSession(w, r); err != nil {
		helper.InternalServerError(w)
		return
	}

	s.logger.Printf("user %s registered", username)
	s.logger.Println("redirecting to /")

	http.Redirect(w, r, "/login", http.StatusFound)
}
