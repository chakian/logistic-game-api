package routes

import (
	"github.com/chakian/logistic-game-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func TruckRoute(route fiber.Router) {
	route.Get("", controllers.GetTrucks)
}
