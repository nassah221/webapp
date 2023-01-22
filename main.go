package main

import (
	"log"
	"os"

	"webapp/db"
	"webapp/helper"
	"webapp/server"
	"webapp/session"

	"github.com/joho/godotenv"
)

func main() {
	helper.LoadTemplates("template/*.html")
	l := log.New(os.Stdout, "[APP] ", log.LstdFlags)

	if err := godotenv.Load(".env"); err != nil {
		l.Fatal("load .env file", err)
	}

	d := db.New()
	s := session.NewStore([]byte(os.Getenv("SESSION_SECRET")))
	h := server.New(d, s, l)
	h.SetupRoutes()

	port := os.Getenv("PORT")
	l.Println("Listening on port ", port)
	l.Println(h.Start(port))
}
