package router

import (
	"chirp-api/app/module/chirp"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App         fiber.Router
	ChirpRouter *chirp.ChirpRouter
}

func NewRouter(fiber *fiber.App, chirpRouter *chirp.ChirpRouter) *Router {
	return &Router{
		App:         fiber,
		ChirpRouter: chirpRouter,
	}
}

func (r *Router) Register() {
	r.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👊")
	})

	r.ChirpRouter.RegisterChirpRoutes()
}
