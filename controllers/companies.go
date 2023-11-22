package controllers

import (
	"net/http"
	"sampleApi/models"

	"github.com/gin-gonic/gin"
)

type CreateCompanyInput struct {
	CompanyName string `json:"companyname"`
	PhoneNumber string `json:"phonenumber"`
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
}

type UpdateCompanyInput struct {
	CompanyName string `json:"companyname"`
	PhoneNumber string `json:"phonenumber"`
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
}

// GetCompanies responds with the list of all companies as JSON.
// GetCompanies             godoc
// @Summary      Get companies array
// @Description  Responds with the list of all customers as JSON.
// @Tags         companies
// @Produce      json
// @Success      200  {array}  CreateCompanyInput
// @Router       /companies [get]
// @Security ApiKeyAuth
func FindCompanies(c *gin.Context) {
	var company []models.Company

	models.DB.Find(&company).Find(&company)

	c.JSON(http.StatusOK, gin.H{"data": company})
}

// Create new company
// PostBook             godoc
// @Summary      Store a new company
// @Description  Takes a company JSON and store in DB. Return saved JSON.
// @Tags         companies
// @Produce      json
// @Param        company  body      CreateCompanyInput  true  "Company JSON"
// @Success      200   {object}  CreateCompanyInput
// @Router       /companies [post]
// @Security ApiKeyAuth
func CreateCompany(c *gin.Context) {
	// Validate input
	var input CreateCompanyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create company
	company := models.Company{
		CompanyName: input.CompanyName,
		PhoneNumber: input.PhoneNumber,
		Address1:    input.Address1,
		Address2:    input.Address2,
		City:        input.City,
		State:       input.State,
		Zip:         input.Zip}

	models.DB.Create(&company)

	c.JSON(http.StatusOK, gin.H{"data": company})
}
