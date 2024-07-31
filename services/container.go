package services

import (
	"api/config"
	"api/repository"

	"gorm.io/gorm"
)

type Container struct {
	DB         *gorm.DB
	PredioRepo repository.PredioRepository
	// Agrega otros repositorios aquí
}

func NewContainer(cfg *config.Config) (*Container, error) {
	dbService := NewDatabaseService(cfg)
	db := dbService.DB

	container := &Container{
		DB:         db,
		PredioRepo: repository.NewPredioRepository(db),
		// Inicializa otros repositorios aquí
	}

	return container, nil
}
