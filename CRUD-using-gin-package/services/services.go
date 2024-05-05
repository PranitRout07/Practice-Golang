package services

import (
	internal "github.com/PranitRout07/Practice-Golang/CRUD-using-gin-package/internal/models"
	"gorm.io/gorm"
)

type DataService struct {
	db *gorm.DB
}

type Svc struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"required"`
}

func (s *DataService) InitService(db *gorm.DB){
	s.db = db
	s.db.AutoMigrate(&internal.Details{})
}

func (s *DataService) GetDataService(data Svc) Svc {
	return data
}

func (s *DataService) PostDataService(data Svc) Svc {
	return data
}

func (s *DataService) DeleteDataService(data Svc) Svc {

	return data

}
