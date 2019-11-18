package controller

import "github.com/synchthia/remonpi/models"

// Controller - Interface of Controller
type Controller interface {
	Set(*models.RemoteData) error
	Generate(*models.RemoteData, *models.GenerateOption) ([][]int, error)
	Send([][]int) error
}

// Database - Individual controller database
type Database interface {
	//Load()
	GetState() *models.State
	Save() error
	Load() error
	UpdateState(*models.RemoteData) error
}
