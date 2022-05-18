package model

type Employee struct{
  EmpID int `json:"id" gorm:"primaryKey"`
  Name string `json:"name"`
  Salary int `json:"salary"`
  Gender string `json:"gender"`
  Country string `json:"country"`
}