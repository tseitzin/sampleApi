package controllers

import (
	"net/http"
	"sampleApi/models"

	"github.com/gin-gonic/gin"
)

type CreateVehicleInput struct {
	VID      int           `json:"vid" gorm:"primary_key"`
	Make     string        `json:"make"`
	CarModel string        `json:"carmodel"`
	VIN      string        `json:"vin"`
	Year     int           `json:"year"`
	VehId    int           `json:"vehid"`
	Tires    []models.Tire `gorm:"ForeignKey:TireID"`
}

type UpdateVehicleInput struct {
	VID      int           `json:"vid" gorm:"primary_key"`
	Make     string        `json:"make"`
	CarModel string        `json:"carmodel"`
	VIN      string        `json:"vin"`
	Year     int           `json:"year"`
	VehId    int           `json:"vehid"`
	Tires    []models.Tire `gorm:"ForeignKey:TireID"`
}

// GetVehicles responds with the list of all vehicles as JSON.
// GetVehicles             godoc
// @Summary      Get Vehicles array
// @Description  Responds with the list of all vehicles as JSON.
// @Tags         vehicles
// @Produce      json
// @Success      200  {array}  models.Vehicle
// @Router       /vehicles [get]
// @Security ApiKeyAuth
func FindVehicles(c *gin.Context) {
	var vehicle []models.Vehicle

	models.DB.Find(&vehicle).Preload("Tires").Find(&vehicle)

	c.JSON(http.StatusOK, gin.H{"data": vehicle})
}

// Create new vehicle
// PostBook             godoc
// @Summary      Store a new vehicle
// @Description  Takes a vehicle JSON and store in DB. Return saved JSON.
// @Tags         vehicles
// @Produce      json
// @Param        vehicle  body      models.Vehicle  true  "Vehicle JSON"
// @Success      200   {object}  models.Vehicle
// @Router       /vehicles [post]
// @Security ApiKeyAuth
func CreateVehicle(c *gin.Context) {
	// Validate input
	var input CreateVehicleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create vehicle
	vehicle := models.Vehicle{
		Make:     input.Make,
		CarModel: input.CarModel,
		VIN:      input.VIN,
		Year:     input.Year,
		VehId:    input.VehId,
		Tires:    input.Tires}

	models.DB.Create(&vehicle)

	c.JSON(http.StatusOK, gin.H{"data": vehicle})
}

// Find a vehicle
// FindVehicle                godoc
// @Summary      Get single vehicle by v_id
// @Description  Returns the vehicle whose ID value matches the ID.
// @Tags         vehicles
// @Produce      json
// @Param        v_id  path      int  true  "search vehicle by v_id"
// @Success      200  {object}  models.Vehicle
// @Router       /vehicles/{v_id} [get]
// @Security ApiKeyAuth
func FindVehicle(c *gin.Context) { // Get model if exist
	var vehicle models.Vehicle

	if err := models.DB.Where("v_id = ?", c.Param("v_id")).First(&vehicle).Preload("Tires").Find(&vehicle).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": vehicle})
}

// Delete a vehicle
// DeleteVehicle                godoc
// @Summary      Delete single vehicle by id
// @Description  Delete the vehicle whose ID value matches the ID.
// @Tags         vehicles
// @Produce      json
// @Param        v_id  path      int  true  "delete vehicle by v_id"
// @Success      200  {object}  models.Vehicle
// @Router       /vehicles/{v_id} [delete]
// @Security ApiKeyAuth
func DeleteVehicle(c *gin.Context) {
	// Get model if exist
	var vehicle models.Vehicle
	if err := models.DB.Where("v_id = ?", c.Param("v_id")).First(&vehicle).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&vehicle)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// Update a vehicle
// UpdateVehicle                godoc
// @Summary      Update single vehicle by id
// @Description  Updates the vehicle whose VID value matches the VID.
// @Tags         vehicles
// @Produce      json
// @Param        v_id  path      int  true  "update vehicle by vid"
// @Param        vehicle  body      models.Vehicle  true  "Vehicle JSON"
// @Success      200  {object}  models.Vehicle
// @Router       /vehicles/{v_id} [patch]
// @Security ApiKeyAuth
func UpdateVehicle(c *gin.Context) {
	// Get model if exist
	var vehicle models.Vehicle
	if err := models.DB.Where("v_id = ?", c.Param("v_id")).First(&vehicle).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateVehicleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UpdateVehicleInput := models.Vehicle{
		Make:     input.Make,
		CarModel: input.CarModel,
		VIN:      input.VIN,
		Year:     input.Year,
		VehId:    input.VehId,
		Tires:    vehicle.Tires}

	models.DB.Model(&vehicle).Updates(&UpdateVehicleInput)

	c.JSON(http.StatusOK, gin.H{"data": UpdateVehicleInput})
}
