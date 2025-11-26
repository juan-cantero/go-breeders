package main

import (
	"flag"
	"fmt"
	"go-breeders/internal/breeder"
	"go-breeders/internal/cat"
	"go-breeders/internal/dog"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type application struct {
	templateMap    map[string]*template.Template
	config         appConfig
	DogHandler     *dog.Handler
	CatHandler     *cat.Handler
	BreederHandler *breeder.Handler
}

type appConfig struct {
	useCache bool
	dsn      string //data source name
}

func main() {
	app := application{
		templateMap: make(map[string]*template.Template),
	}

	flag.BoolVar(&app.config.useCache, "cache", false, "Use template cache")
	flag.StringVar(&app.config.dsn, "dsn",
		"mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "DSN")
	flag.Parse()

	db, err := initMySQLDB(app.config.dsn)
	if err != nil {
		log.Panic(err)
	}

	// Wire up Dog domain (Repository -> Service -> Handler)
	dogRepo := dog.NewMySQLRepository(db)
	dogService := dog.NewService(dogRepo)
	app.DogHandler = dog.NewHandler(dogService)

	// Wire up Cat domain
	catRepo := cat.NewMySQLRepository(db)
	catService := cat.NewService(catRepo)
	app.CatHandler = cat.NewHandler(catService)

	// Wire up Breeder domain
	breederRepo := breeder.NewMySQLRepository(db)
	breederService := breeder.NewService(breederRepo)
	app.BreederHandler = breeder.NewHandler(breederService)

	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("Starting on port", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
