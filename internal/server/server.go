package server

import (
	"github.com/gofiber/fiber/v2"

	"url-shortener/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "url-shortener",
			AppName:      "url-shortener",
		}),

		db: database.New(),
	}

	return server
}
