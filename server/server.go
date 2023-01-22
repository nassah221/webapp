package server

import (
	"log"
	"net/http"
	"time"
	database "webapp/db"
	sess "webapp/session"

	"github.com/gorilla/mux"
)

type Server struct {
	db      database.DB
	session *sess.Store
	logger  *log.Logger
	router  *mux.Router
}

func New(
	db database.DB,
	session *sess.Store,
	logger *log.Logger,
) *Server {
	return &Server{
		db:      db,
		session: session,
		logger:  logger,
		router:  mux.NewRouter(),
	}
}

func (s *Server) SetupRoutes() {
	s.router.HandleFunc("/", s.Authorize(s.IndexGETHandler)).Methods(http.MethodGet)
	s.router.HandleFunc("/", s.Authorize(s.IndexPOSTHandler)).Methods(http.MethodPost)
	s.router.HandleFunc("/login", s.LoginGETHandler).Methods(http.MethodGet)
	s.router.HandleFunc("/login", s.LoginPOSTHandler).Methods(http.MethodPost)
	s.router.HandleFunc("/logout", s.LogoutHandler).Methods(http.MethodGet)
	s.router.HandleFunc("/register", s.RegisterGETHandler).Methods(http.MethodGet)
	s.router.HandleFunc("/register", s.RegisterPOSTHandler).Methods(http.MethodPost)

	fs := http.FileServer(http.Dir("./static/"))
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}

func (s *Server) Start(port string) error {
	srv := &http.Server{ //nolint
		Handler:      s.router,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		Addr:         ":" + port,
	}

	return srv.ListenAndServe()
}
