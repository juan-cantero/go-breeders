package main

import (
	"go-breeders/internal/breeder"
	"go-breeders/internal/cat"
	"go-breeders/internal/dog"
	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
	// Setup - wire up each domain with mock repositories
	// Repository -> Service -> Handler chain for each domain

	// Dog domain with mock
	dogRepo := dog.NewMockRepository()
	dogService := dog.NewService(dogRepo)
	dogHandler := dog.NewHandler(dogService)

	// Cat domain with mock
	catRepo := cat.NewMockRepository()
	catService := cat.NewService(catRepo)
	catHandler := cat.NewHandler(catService)

	// Breeder domain with mock
	breederRepo := breeder.NewMockRepository()
	breederService := breeder.NewService(breederRepo)
	breederHandler := breeder.NewHandler(breederService)

	testApp = application{
		DogHandler:     dogHandler,
		CatHandler:     catHandler,
		BreederHandler: breederHandler,
	}

	// Run all tests
	os.Exit(m.Run())
}
