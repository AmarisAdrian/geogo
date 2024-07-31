package repository

import (
	"api/models"
	"errors"

	"gorm.io/gorm"
)

type PredioRepository interface {
	GetAll() ([]models.PrediosModel, error)
	GetById(t_id int) (models.PrediosModel, error)
	Create(Predios models.PrediosModel) (models.PrediosModel, error)
	Update(id int, Predios models.PrediosModel) (models.PrediosModel, error)
	Delete(id int) error
}

type GormPredioRepository struct {
	db *gorm.DB
}

func NewPredioRepository(db *gorm.DB) *GormPredioRepository {
	return &GormPredioRepository{
		db: db,
	}
}
func (r *GormPredioRepository) GetAll() ([]models.PrediosModel, error) {
	var predio []models.PrediosModel
	if err := r.db.Find(&predio).Error; err != nil {
		return nil, err
	}
	return predio, nil
}

func (r *GormPredioRepository) GetById(t_id int) (models.PrediosModel, error) {
	var predio models.PrediosModel
	if err := r.db.First(&predio, t_id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.PrediosModel{}, err
		}
		return models.PrediosModel{}, err
	}
	return predio, nil
}

func (r *GormPredioRepository) Create(predio models.PrediosModel) (models.PrediosModel, error) {
	if err := r.db.Create(&predio).Error; err != nil {
		return models.PrediosModel{}, err
	}
	return predio, nil
}

func (r *GormPredioRepository) Update(id int, predio models.PrediosModel) (models.PrediosModel, error) {
	if err := r.db.Model(&models.PrediosModel{}).Where("t_id = ?", id).Updates(predio).Error; err != nil {
		return models.PrediosModel{}, err
	}
	return predio, nil
}

func (r *GormPredioRepository) Delete(id int) error {
	if err := r.db.Delete(&models.PrediosModel{}, id).Error; err != nil {
		return err
	}
	return nil
}
