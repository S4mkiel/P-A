package fiber

import (
	"github.com/S4mkiel/p-a/infra/http/controller"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("fiber",
	FiberModule,
	fx.Invoke(RegisterControllers),
	fx.Provide(controller.NewPaController),
)

func RegisterControllers(app *fiber.App, paController *controller.PaController) {
	api := app.Group("/api")
	api = api.Group("/p-a")

	paController.RegisterRoutes(api)
}
