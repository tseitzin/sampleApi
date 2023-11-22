package controllers

import (
	"net/http"
	"sampleApi/models"

	"github.com/gin-gonic/gin"
)

type CreateTireInput struct {
	TID             int    `json:"tid" gorm:"primary_key"`
	TireLocation    string `json:"tirelocation"`
	Brand           string `json:"make"`
	TireModel       string `json:"tiremodel"`
	TireSize        string `json:"tiresize"`
	MaxPressure     int    `json:"maxpressure"`
	CurrentPressure int    `json:"currentpressure"`
	TreadWear       string `json:"treadwear"`
	TireID          int    `json:"tireid"`
}

type UpdateTireInput struct {
	TID             int    `json:"tid" gorm:"primary_key"`
	TireLocation    string `json:"tirelocation"`
	Brand           string `json:"make"`
	TireModel       string `json:"tiremodel"`
	TireSize        string `json:"tiresize"`
	MaxPressure     int    `json:"maxpressure"`
	CurrentPressure int    `json:"currentpressure"`
	TreadWear       string `json:"treadwear"`
	TireID          int    `json:"tireid"`
}

// GetTires responds with the list of all tires as JSON.
// GetTires             godoc
// @Summary      Get Tires array
// @Description  Responds with the list of all Tires as JSON.
// @Tags         tires
// @Produce      json
// @Success      200  {array}  models.Tire
// @Router       /tires [get]
// @Security ApiKeyAuth
func FindTires(c *gin.Context) {
	var tire []models.Tire

	models.DB.Find(&tire).Find(&tire)

	c.JSON(http.StatusOK, gin.H{"data": tire})
}

// Create new tire
// PostBook             godoc
// @Summary      Store a new tire
// @Description  Takes a tire JSON and store in DB. Return saved JSON.
// @Tags         tires
// @Produce      json
// @Param        tire  body      models.Tire  true  "Tire JSON"
// @Success      200   {object}  models.Tire
// @Router       /tires [post]
// @Security ApiKeyAuth
func CreateTire(c *gin.Context) {
	// Validate input
	var input CreateTireInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create tire
	tire := models.Tire{
		TireLocation:    input.TireLocation,
		Brand:           input.Brand,
		TireModel:       input.TireModel,
		TireSize:        input.TireSize,
		MaxPressure:     input.MaxPressure,
		CurrentPressure: input.CurrentPressure,
		TreadWear:       input.TreadWear,
		TireID:          input.TireID}

	models.DB.Create(&tire)

	c.JSON(http.StatusOK, gin.H{"data": tire})
}

// Find a tire
// FindTire               godoc
// @Summary      Get single tire by v_id
// @Description  Returns the tire whose ID value matches the TID.
// @Tags         tires
// @Produce      json
// @Param        t_id  path      string  true  "search tire by t_id"
// @Success      200  {object}  models.Tire
// @Router       /tires/{t_id} [get]
// @Security ApiKeyAuth
func FindTire(c *gin.Context) { // Get model if exist
	var tire models.Tire

	if err := models.DB.Where("t_id = ?", c.Param("t_id")).First(&tire).Find(&tire).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tire})
}

// Delete a tire
// DeleteTire                godoc
// @Summary      Delete single tire by id
// @Description  Delete the tire whose ID value matches the ID.
// @Tags         tires
// @Produce      json
// @Param        t_id  path      string  true  "delete tire by t_id"
// @Success      200  {object}  models.Tire
// @Router       /tires/{t_id} [delete]
// @Security ApiKeyAuth
func DeleteTire(c *gin.Context) {
	// Get model if exist
	var tire models.Tire
	if err := models.DB.Where("t_id = ?", c.Param("t_id")).First(&tire).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&tire)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// Update a tire
// UpdateTire                godoc
// @Summary      Update single tire by id
// @Description  Updates the tire whose ID value matches the ID.
// @Tags         tires
// @Produce      json
// @Param        t_id  path      int  true  "update tire by id"
// @Param        tire  body      models.Tire  true  "Tire JSON"
// @Success      200  {object}  models.Tire
// @Router       /tires/{t_id} [patch]
// @Security ApiKeyAuth
func UpdateTire(c *gin.Context) {
	// Get model if exist
	var tire models.Tire
	if err := models.DB.Where("t_id = ?", c.Param("t_id")).First(&tire).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateTireInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UpdateTireInput := models.Tire{
		TireLocation:    input.TireLocation,
		Brand:           input.Brand,
		TireModel:       input.TireModel,
		TireSize:        input.TireSize,
		MaxPressure:     input.MaxPressure,
		CurrentPressure: input.CurrentPressure,
		TreadWear:       input.TreadWear,
		TireID:          input.TireID}

	models.DB.Model(&tire).Updates(&UpdateTireInput)

	c.JSON(http.StatusOK, gin.H{"data": UpdateTireInput})
}
