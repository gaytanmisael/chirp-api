package controller

import "chirp-api/app/module/chirp/service"

type Controller struct {
	Chirp *ChirpController
}

func NewController(chirpService *service.ChirpService) *Controller {
	return &Controller{
		Chirp: &ChirpController{chirpService: chirpService},
	}
}
