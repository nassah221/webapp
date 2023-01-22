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
		l.Println("load .env: ", err)

		l.Println("setting default env vars")
		// a bit clunky
		if err := os.Setenv("PORT", "8080"); err != nil {
			l.Fatal("failed to set port: ", err)
		}
		if err := os.Setenv("SESSION_SECRET", "s3cr3t"); err != nil {
			l.Fatal("failed to set session secret: ", err)
		}
	}

	d := db.New()
	s := session.NewStore([]byte(os.Getenv("SESSION_SECRET")))
	h := server.New(d, s, l)
	h.SetupRoutes()

	port := os.Getenv("PORT")
	l.Println("Listening on port ", port)
	l.Println(h.Start(port))
}
