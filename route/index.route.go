package route

import (
	"go-fiber-gorm/handler"
	"go-fiber-gorm/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

	r.Post("/login", handler.LoginHandler)

	r.Get("/user", handler.UserHandlerGet)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Post("/user", handler.UserHandlerCreat)
	r.Put("/user/:id", handler.UserHandlerUpdate)
	r.Delete("/user/:id", middleware.Auth, handler.UserHandlerDelete)

}
