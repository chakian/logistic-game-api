package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type Truck struct {
	Id        int    `json:"id"`
	TruckType string `json:"truckType"`
	Name      string `json:"name"`
}

func GetTrucks(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"trucks": getTrucksOfUser(),
		},
	})
}

func getTrucksOfUser() []*Truck {
	var trucks = []*Truck{
		{
			Id:        1,
			TruckType: "Buyuk Kamyon",
			Name:      "Aslan Kamyonum",
		},
		{
			Id:        2,
			TruckType: "Kucuk Kamyon",
			Name:      "Minnos",
		},
	}
	return trucks
}
