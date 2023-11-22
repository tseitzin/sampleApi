package controllers

import (
	"net/http"
	"sampleApi/models"

	"github.com/gin-gonic/gin"
)

type CreateCustomerInput struct {
	LastName    string           `json:"lastname" binding:"required"`
	FirstName   string           `json:"firstname" binding:"required"`
	PhoneNumber string           `json:"phonenumber" binding:"required"`
	Address1    string           `json:"address1"`
	Address2    string           `json:"address2"`
	City        string           `json:"city"`
	State       string           `json:"state"`
	Zip         string           `json:"zip"`
	Vehicles    []models.Vehicle `gorm:"ForeignKey:VehId"`
}

type UpdateCustomerInput struct {
	LastName    string           `json:"lastname"`
	FirstName   string           `json:"firstname"`
	PhoneNumber string           `json:"phonenumber"`
	Address1    string           `json:"address1"`
	Address2    string           `json:"address2"`
	City        string           `json:"city"`
	State       string           `json:"state"`
	Zip         string           `json:"zip"`
	Vehicles    []models.Vehicle `gorm:"ForeignKey:VehId"`
}

// GetCustomers responds with the list of all customers as JSON.
// GetCustomers             godoc
// @Summary      Get Customers array
// @Description  Responds with the list of all customers as JSON.
// @Tags         customers
// @Produce      json
// @Success      200  {array}  models.Customer
// @Router       /customers [get]
// @Security ApiKeyAuth
func FindCustomers(c *gin.Context) {
	var customer []models.Customer

	models.DB.Preload("Vehicles").Preload("Vehicles.Tires").Find(&customer)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// Create new customer
// PostBook             godoc
// @Summary      Store a new customer
// @Description  Takes a customer JSON and store in DB. Return saved JSON.
// @Tags         customers
// @Produce      json
// @Param        customer  body      models.Customer  true  "Customer JSON"
// @Success      200   {object}  models.Customer
// @Router       /customers [post]
// @Security ApiKeyAuth
func CreateCustomer(c *gin.Context) {
	// Validate input
	var input CreateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create customer
	customer := models.Customer{
		LastName:    input.LastName,
		FirstName:   input.FirstName,
		PhoneNumber: input.PhoneNumber,
		Address1:    input.Address1,
		Address2:    input.Address2,
		City:        input.City,
		State:       input.State,
		Zip:         input.Zip,
		Vehicles:    input.Vehicles}

	models.DB.Create(&customer)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// Delete a customer
// DeleteCustomer                godoc
// @Summary      Delete single customer by id
// @Description  Delete the customer whose ID value matches the ID.
// @Tags         customers
// @Produce      json
// @Param        id  path      string  true  "delete customer by id"
// @Success      200  {object}  models.Customer
// @Router       /customers/{id} [delete]
// @Security ApiKeyAuth
func DeleteCustomer(c *gin.Context) {
	// Get model if exist
	var customer models.Customer
	if err := models.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&customer)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// Find a customer
// FindBook                godoc
// @Summary      Get single customer by id
// @Description  Returns the customer whose ID value matches the ID.
// @Tags         customers
// @Produce      json
// @Param        id  path      string  true  "search customer by id"
// @Success      200  {object}  models.Customer
// @Router       /customers/{id} [get]
// @Security ApiKeyAuth
func FindCustomer(c *gin.Context) { // Get model if exist
	var customer models.Customer

	if err := models.DB.Where("id = ?", c.Param("id")).First(&customer).Preload("Vehicles").Find(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// Update a customer
// UpdateCustomer                godoc
// @Summary      Update single customer by id
// @Description  Updates the book whose ID value matches the ID.
// @Tags         customers
// @Produce      json
// @Param        id  path      int  true  "update customer by id"
// @Param        customer  body      models.Customer  true  "Customer JSON"
// @Success      200  {object}  models.Customer
// @Router       /customers/{id} [patch]
// @Security ApiKeyAuth
func UpdateCustomer(c *gin.Context) {
	// Get model if exist
	var customer models.Customer
	if err := models.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UpdateCustomerInput := models.Customer{
		ID:          customer.ID,
		LastName:    input.LastName,
		FirstName:   input.FirstName,
		PhoneNumber: input.PhoneNumber,
		Address1:    input.Address1,
		Address2:    input.Address2,
		City:        input.City,
		State:       input.State,
		Zip:         input.Zip,
		Vehicles:    customer.Vehicles}

	models.DB.Model(&customer).Updates(&UpdateCustomerInput)

	c.JSON(http.StatusOK, gin.H{"data": UpdateCustomerInput})
}
