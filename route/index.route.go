package route

import (
	"go-fiber-gorm/handler"
	"go-fiber-gorm/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

	r.Post("/login", handler.LoginHandler)

	r.Get("/user", middleware.Auth, handler.UserHandlerGet)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Post("/user", handler.UserHandlerCreat)
	r.Put("/user/:kode", handler.UserHandlerUpdate)
	r.Delete("/user/:kode", handler.UserHandlerDelete)

}
