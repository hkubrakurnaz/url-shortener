package url_controller

import (
	"github.com/gofiber/fiber/v2"
	url_service "url-shortener/service/url"
)

type Controller interface {
	CreateShortUrl(ctx *fiber.Ctx) error
	GetUrl(ctx *fiber.Ctx) error
}

type controller struct {
	urlService url_service.Service
}

func (c controller) CreateShortUrl(ctx *fiber.Ctx) error {
	var payload *CreateShortUrl
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	response, err := c.urlService.CreateShortUrl(payload.Url)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"url": response})
}

func (c controller) CreateShortUrls(ctx *fiber.Ctx) error {
	var payload *CreateShortUrls
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	response, err := c.urlService.CreateShortUrls(payload.Urls)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"urls": ToCreateBulkResponse(response)})
}

func (c controller) GetUrl(ctx *fiber.Ctx) error {
	shortUrlParam := ctx.Params("shortUrl")
	response, err := c.urlService.GetLongUrl(shortUrlParam)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.Status(fiber.StatusPermanentRedirect).JSON(fiber.Map{"url": response})
}

func New(app *fiber.App, urlService url_service.Service) Controller {
	c := &controller{urlService: urlService}

	router := app.Group("/v1/urls")
	router.Post("", c.CreateShortUrl)
	router.Post("bulk", c.CreateShortUrls)
	router.Get("/:shortUrl", c.GetUrl)

	return c
}
