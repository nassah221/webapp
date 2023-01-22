package server

import (
	"html/template"
	"log"
	"net/http"
	"time"
	"webapp/helper"
	"webapp/model"
)

func (s *Server) IndexGETHandler(w http.ResponseWriter, r *http.Request) {
	s.logger.Println("index GET handler")

	username, ok := r.Context().Value("user").(string)
	if !ok {
		panic("should never happen")
	}

	user, err := s.db.FindUser(username)
	if err != nil {
		log.Printf("find user: %v", err)
		// in-mem db deletes users on server restart, so we need to clear existing session if any
		// if the server has restarted and the user doesn't exist anymore from a previous session
		if err := s.session.DeleteCurrentSession(w, r); err != nil {
			helper.InternalServerError(w)
			return
		}

		s.logger.Println("redirecting to /")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	helper.ExecuteTemplate(w, "index.html", struct {
		User     string
		Feedback []model.Feedback
	}{
		User:     user.Username,
		Feedback: user.Feedback,
	})
}

func (s *Server) IndexPOSTHandler(w http.ResponseWriter, r *http.Request) {
	s.logger.Println("index POST handler")

	username, err := s.session.GetUsername(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	title := template.HTMLEscapeString(r.PostFormValue("title"))
	comment := template.HTMLEscapeString(r.PostFormValue("comment"))

	if title != "" && comment != "" {
		s.logger.Printf("user %s created feedback", username)
		s.db.AppendFeedback(username, model.Feedback{ //nolint
			Title:   title,
			Comment: comment,
			Time:    time.Now().Format(time.RFC822),
		})
	}

	s.logger.Println("redirecting to /")
	http.Redirect(w, r, "/", http.StatusFound)
}
