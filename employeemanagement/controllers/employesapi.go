package controllers

import (
	"example/web-service-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/rahmanfadhil/gin-bookstore/models"
)

// GET /employees
// Get all employees
func FindEmployees(c *gin.Context) {
	var employees []models.Employee
	models.DB.Find(&employees)

	c.JSON(http.StatusOK, gin.H{"data": employees})
}

// POST /employees
// Create new book
func CreateEmployee(c *gin.Context) {
	// Validate input
	var input models.CreateEmployeeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create employee
	employee := models.Employee{EmpName: input.EmpName, EmpPosition: input.EmpPosition}
	models.DB.Create(&employee)

	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// GET /employees/:id
// Find a Employee
func FindEmployee(c *gin.Context) { // Get model if exist
	var employee models.Employee

	if err := models.DB.Where("id = ?", c.Param("id")).First(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// PATCH /employee/:id
// Update a employee
func UpdateEmployee(c *gin.Context) {
	// Get model if exist
	var employee models.Employee
	if err := models.DB.Where("id = ?", c.Param("id")).First(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateEmployeeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&employee).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// DELETE /employees/:id
// Delete a employee
func DeleteEmployee(c *gin.Context) {
	// Get model if exist
	var employee models.Employee
	if err := models.DB.Where("id = ?", c.Param("id")).First(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&employee)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
