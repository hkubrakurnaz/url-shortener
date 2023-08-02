package url

import (
	"github.com/gofiber/fiber/v2"
	"url-shortener/service/url"
)

type Controller interface {
	Register(app *fiber.App)
	CreateShortUrl(ctx *fiber.Ctx) error
	GetUrl(ctx *fiber.Ctx) error
}

type controller struct {
	urlService url.Service
}

func (c controller) Register(app *fiber.App) {
	router := app.Group("/v1/urls")
	router.Post("", c.CreateShortUrl)
	router.Get("/:shortUrl", c.GetUrl)
}

func (c controller) CreateShortUrl(ctx *fiber.Ctx) error {
	var payload *CreateShortUrl
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	response, _ := c.urlService.CreateShortUrl(payload.Url)
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"url": response})
}

func (c controller) GetUrl(ctx *fiber.Ctx) error {
	shortUrlParam := ctx.Params("shortUrl")
	response, _ := c.urlService.GetLongUrl(shortUrlParam)
	return ctx.Status(fiber.StatusPermanentRedirect).JSON(fiber.Map{"url": response})
}

func New(urlService url.Service) Controller {
	return &controller{urlService: urlService}
}
