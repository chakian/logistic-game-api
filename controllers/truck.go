package controllers

import (
	"fmt"

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

func AddTruck(c *fiber.Ctx) error {
	type Request struct {
		Name string `json:"name"`
	}

	var body Request

	err := c.BodyParser(&body)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	truck := &Truck{
		Id:        5,
		TruckType: "",
		Name:      body.Name,
	}

	// TODO: add truck to db

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"truck": truck,
		},
	})
}

func getTrucksOfUser() []*Truck {
	// TODO: really, get trucks of user
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
