package models

import (
	"github.com/jinzhu/gorm"
	"github.com/gofiber/fiber/v2"
	"hamakeja_api/database"
	
)

type Truck struct {
	gorm.Model 
	TruckNumber string `json:"truck_number" gorm:"unique; not null"`
	ParkedLocation string `json:"parked_location"`
	ImageUrl string `json:"image_url"`
	ContactNumber string `json:"contact_number"`
	OverallSize string `json:"overall_size"`
	CarryingCapacity string `json:"carrying_capacity"`
}

func CreateTruck(c *fiber.Ctx) error {
	db := database.DBConn

	truck := new(Truck)

	if err := c.BodyParser(truck); err != nil {
		c.Status(503).Send([]byte(err.Error()))
		return nil

	}

	db.Create(&truck)
	return c.JSON(truck)

}

func GetTruck(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var truck Truck
	db.Find(&truck, id)
	return c.JSON(truck)
}

func GetTrucks(c *fiber.Ctx) error {
	db := database.DBConn
	
	var trucks []Truck

	db.Find(&trucks)
	return c.JSON(trucks)

}

func UpdateTruck(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn 
	var truck Truck

	db.Find(&truck, id)

	if truck.TruckNumber == "" {
		c.Status(500).SendString("Truck with provided Id not found")

		return nil

	}

	type Data struct {
		TruckNumber string `json:"truck_number"`
		ParkedLocation string `json:"parked_location"`
		ContactNumber string `json:"contact_number"`
		OverallSize string `json:"overall_size"`
		CarryingCapacity string `json:"carrying_capacity"`
	}

	var updateData Data
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	truck.TruckNumber = updateData.TruckNumber
	truck.ParkedLocation = updateData.ParkedLocation
	truck.ContactNumber = updateData.ContactNumber
	truck.OverallSize = updateData.OverallSize
	truck.CarryingCapacity = updateData.CarryingCapacity
	db.Save(&truck)

	return c.JSON(truck)
}

func DeleteTruck(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DBConn

	var truck Truck

	db.First(&truck, id)

	if truck.TruckNumber == "" {

		c.Status(500).SendString("No truck found with given ID")
		return nil
		
	}
	db.Delete(&truck)

	return c.SendString("Truck Deleted Successfully")
}