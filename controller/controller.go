package controller

import "github.com/synchthia/remonpi/models"

type Controller interface {
	Set(*models.RemoteData) error
	Send(*models.RemoteData) error
	Generate(*models.RemoteData) ([][]int, error)
}

type Database interface {
	//Load()
	GetState() *models.State
	Save() error
	Load() error
	UpdateState(*models.RemoteData) error
}
