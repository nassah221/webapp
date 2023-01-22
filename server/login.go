package server

import (
	"errors"
	"html/template"
	"net/http"
	"webapp/db"
	"webapp/helper"
	"webapp/model"

	"golang.org/x/crypto/bcrypt"
)

func (s *Server) LoginGETHandler(w http.ResponseWriter, r *http.Request) {
	s.logger.Println("login GET handler")

	sess, err := s.session.Get(r, "session")
	if err != nil {
		helper.InternalServerError(w)
		return
	}

	if _, ok := sess.Values["user_id"]; ok {
		s.logger.Println("session exists, redirecting to /")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	s.logger.Println("redirecting to /login")
	helper.ExecuteTemplate(w, "login.html", nil)
}

func (s *Server) LoginPOSTHandler(w http.ResponseWriter, r *http.Request) {
	username := template.HTMLEscapeString(r.PostFormValue("username"))
	password := template.HTMLEscapeString(r.PostFormValue("password"))

	// basic form validation
	if username == "" || password == "" {
		helper.ExecuteTemplate(w, "login.html", "username and password cannot be empty")
		return
	}

	user, err := s.authenticateUser(username, password)
	if err != nil {
		switch err {
		case ErrInvalidLogin:
			helper.ExecuteTemplate(w, "login.html", err.Error())
		case db.ErrNotFound:
			helper.ExecuteTemplate(w, "login.html", err.Error())
		default:
			helper.InternalServerError(w)
		}
		s.logger.Printf("user %s failed login: %v", username, err)
		return
	}

	sess, err := s.session.Get(r, "session")
	if err != nil {
		helper.InternalServerError(w)
		return
	}
	sess.Values["user_id"] = user.Username
	if err := sess.Save(r, w); err != nil {
		helper.InternalServerError(w)
		return
	}

	s.logger.Printf("user %s successful login", user.Username)

	s.logger.Println("redirecting to /")
	http.Redirect(w, r, "/", http.StatusFound)
}

var ErrInvalidLogin = errors.New("invalid login")

func (s *Server) authenticateUser(username, password string) (*model.User, error) {
	// check if username exists
	user, err := s.db.FindUser(username)
	if err != nil {
		return nil, err
	}

	// check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, ErrInvalidLogin
	}

	return user, nil
}
