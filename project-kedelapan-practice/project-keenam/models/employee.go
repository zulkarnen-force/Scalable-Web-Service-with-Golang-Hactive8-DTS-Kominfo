package models

import "time"

type Employee struct {
	ID int 			 `gorm:"primaruKey"`
	Name string 	 `json:"name"`
	Age int 		 `json:"age"`
	Division string	 `json:"division"`
	CreatedAt time.Time
	UpdatedAt time.Time
}