package web

import (
	"log"
	"memorycloud/cloud/internal/service"

	"github.com/gofiber/fiber/v2"
)

type WebController struct {
	serv *service.GRPCService
	app  *fiber.App
	log  *log.Logger
}

func CreateWebController(service *service.GRPCService, app *fiber.App, logger *log.Logger) *WebController {
	return &WebController{
		serv: service,
		app:  app,
		log:  logger,
	}
}

func (wc *WebController) RegisterRoutes() {}
