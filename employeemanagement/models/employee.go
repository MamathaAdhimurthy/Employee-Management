package models

type Employee struct {
	ID          int64  `json:"id" gorm:"primary_key"`
	EmpName     string `json:"empname"`
	EmpPosition string `json:"empposition"`
}

type CreateEmployeeInput struct {
	EmpName     string `json:"empname" binding:"required"`
	EmpPosition string `json:"empposition" binding:"required"`
}
type UpdateEmployeeInput struct {
	EmpName     string `json:"empname"`
	EmpPosition string `json:"empposition"`
}
