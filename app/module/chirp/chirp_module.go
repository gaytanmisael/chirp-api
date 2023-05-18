package chirp

import (
	"chirp-api/app/module/chirp/controller"
	"chirp-api/app/module/chirp/repository"
	"chirp-api/app/module/chirp/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type ChirpRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

var NewChirpModule = fx.Options(
	fx.Provide(repository.NewChirpRepository),
	fx.Provide(service.NewChirpService),

	fx.Provide(controller.NewController),

	fx.Provide(NewChirpRouter),
)

func NewChirpRouter(fiber *fiber.App, controller *controller.Controller) *ChirpRouter {
	return &ChirpRouter{
		App:        fiber,
		Controller: controller,
	}
}

func (r *ChirpRouter) RegisterChirpRoutes() {
	chirpController := r.Controller.Chirp

	r.App.Route("/chirps", func(router fiber.Router) {
		router.Get("/", chirpController.Index)
		router.Get("/:id", chirpController.Show)
		router.Post("/", chirpController.Store)
		router.Patch("/:id", chirpController.Update)
		router.Delete("/:id", chirpController.Destroy)
	})
}
