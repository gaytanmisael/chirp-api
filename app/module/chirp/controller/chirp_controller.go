package controller

import (
	"chirp-api/app/module/chirp/service"
	"chirp-api/utils/response"

	"github.com/gofiber/fiber/v2"
)

type ChirpController struct {
	chirpService *service.ChirpService
}

type IChirpController interface {
	Index(c *fiber.Ctx) error
}

func NewChirpController(chirpService *service.ChirpService) *ChirpController {
	return &ChirpController{
		chirpService: chirpService,
	}
}

func (con *ChirpController) Index(c *fiber.Ctx) error {
	posts, err := con.chirpService.GetChirps()
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"Chirp list retrieved successfully!"},
		Data:     posts,
	})
}
