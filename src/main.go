package main

import (
	"log"

	"github.com/DavinY/go-restful-gin/config"
	"github.com/DavinY/go-restful-gin/conn/mongodb"
	// CustomerHandler "github.com/andromedash/colour-card-apps/module/v1/customer/presenter"
	// CustomerRepo "github.com/andromedash/colour-card-apps/module/v1/customer/repository"
	// CustomerUsecase "github.com/andromedash/colour-card-apps/module/v1/customer/usecase"
	// OfferHandler "github.com/andromedash/colour-card-apps/module/v1/offer/presenter"
	// OfferRepo "github.com/andromedash/colour-card-apps/module/v1/offer/repository"
	// OfferUsecase "github.com/andromedash/colour-card-apps/module/v1/offer/usecase"
	// "github.com/gorilla/mux"
)

const (
	logPrefix = "[APP INIT]"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(logPrefix, "failed to load config:", err)
	}
	config.Set(cfg)
	log.Println("starting application: initializing...")

	err = mongodb.Init()
	if err != nil {
		log.Fatal(logPrefix, "failed to init mongodb", err)
	}
	if err != nil {
		log.Fatal(logPrefix, "failed to init mongodb", err)
	}

	// mongoDBClient := mongodb.GetDB()

}
