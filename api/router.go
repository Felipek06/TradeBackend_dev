package api

import (
	_ "github.com/Felipek06/TradeBackend_dev.git/docs"
	"github.com/Felipek06/TradeBackend_dev.git/services"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func router(app *fiber.App, userService *services.NewUserService, authService *services.NewAuthService) {
	handler := &NewHandler{userService: userService, authService: authService}

	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.Redirect("/docs/index.html", fiber.StatusFound)
	})
	app.Get("/docs/", func(c *fiber.Ctx) error {
		return c.Redirect("/docs/index.html", fiber.StatusFound)
	})
	app.Get("/docs/*", fiberSwagger.WrapHandler)

	usersRoutes(app, handler)
}

func usersRoutes(app *fiber.App, handler *NewHandler) {
	app.Post("/api/users", handler.CreateUser)
	app.Post("/api/login", handler.Login)
}
