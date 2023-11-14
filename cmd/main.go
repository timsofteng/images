package main

import (
	"images/internal/config"
	"images/internal/httpClient"
	"images/internal/secrets"
	"images/pkg/delivery/grpcServer"
	"images/pkg/repo/unsplash"
	"images/pkg/usecases"
	"log"
)

func main() {
	s, err := secrets.New()

	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.New()

	if err != nil {
		log.Fatal(err)
	}

	newHTTPClient := httpClient.New()

	newUnsplashRepo := unsplash.New(newHTTPClient, s.UnsplashClientID)

	uc := usecases.New(newUnsplashRepo)

	err = grpcServer.New(uc, cfg.GRPCPort)

	if err != nil {
		log.Printf("%v", err)
	}
}
