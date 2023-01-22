package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type Store struct {
	store *sessions.CookieStore
}

func NewStore(secret []byte) *Store {
	store := sessions.NewCookieStore(secret)
	store.Options = &sessions.Options{ //nolint
		MaxAge:   60 * 15,
		HttpOnly: true,
	}
	return &Store{store: store}
}

func (s *Store) GetUsername(r *http.Request) (string, error) {
	session, err := s.store.Get(r, "session")
	if err != nil {
		return "", err
	}
	username, ok := session.Values["user_id"]
	if !ok {
		return "", err
	}

	return username.(string), err
}

func (s *Store) Get(r *http.Request, key string) (*sessions.Session, error) {
	return s.store.Get(r, key)
}

func (s *Store) DeleteCurrentSession(w http.ResponseWriter, r *http.Request) error {
	sess, err := s.store.Get(r, "session")
	if err != nil {
		return err
	}
	delete(sess.Values, "user_id")

	return sess.Save(r, w)
}
