package controller

import (
	"chirp-api/app/module/chirp/request"
	"chirp-api/app/module/chirp/service"
	"chirp-api/utils/response"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ChirpController struct {
	chirpService *service.ChirpService
}

type IChirpController interface {
	Index(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Destroy(c *fiber.Ctx) error
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

func (con *ChirpController) Show(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return err
	}

	chirp, err := con.chirpService.GetChiprsByID(id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"The Chirp retrieved successfully!"},
		Data:     chirp,
	})
}

func (con *ChirpController) Store(c *fiber.Ctx) error {
	req := new(request.PostRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	chirp, err := con.chirpService.CreateChirp(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"The chirp was created successfully!"},
		Data:     chirp,
	})
}

func (con *ChirpController) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return err
	}

	req := new(request.PostRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}

	chirp, err := con.chirpService.UpdateChirp(id, *req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"The article was updated successfully!"},
		Data:     chirp,
	})
}

func (con *ChirpController) Destroy(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return err
	}

	if err = con.chirpService.DeleteChirp(id); err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.Messages{"The chirp was deleted successfully!"},
	})
}
