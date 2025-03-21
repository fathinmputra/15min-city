package dataset_db

import (
	"15min-city/entity"
	"15min-city/pkg/errs"
	"15min-city/repository/dataset_repository"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type datasetRepository struct {
	db *gorm.DB
}

// NewDatasetRepository creates a new instance of DatasetRepository.
func NewDatasetRepository(db *gorm.DB) dataset_repository.DatasetRepository {
	return &datasetRepository{
		db: db,
	}
}

func (r *datasetRepository) CreateDataset(ctx context.Context, dataset entity.Dataset) (*entity.Dataset, errs.ErrMessage) {
	if err := r.db.WithContext(ctx).Create(&dataset).Error; err != nil {
		return nil, errs.NewInternalServerError("failed to create dataset")
	}
	return &dataset, nil
}

func (r *datasetRepository) GetDatasetByID(ctx context.Context, id int) (*entity.Dataset, errs.ErrMessage) {
	var dataset entity.Dataset
	if err := r.db.WithContext(ctx).First(&dataset, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("dataset not found")
		}
		return nil, errs.NewInternalServerError("failed to get dataset")
	}
	return &dataset, nil
}

func (r *datasetRepository) GetDatasetByName(ctx context.Context, name string) ([]entity.Dataset, errs.ErrMessage) {
	var datasets []entity.Dataset
	if err := r.db.WithContext(ctx).Where("name LIKE ?", "%"+name+"%").Find(&datasets).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("dataset not found")
		}
		return nil, errs.NewInternalServerError("failed to get datasets by name")
	}
	return datasets, nil
}

func (r *datasetRepository) GetDatasetByKecamatan(ctx context.Context, kecamatan string) ([]entity.Dataset, errs.ErrMessage) {
	var datasets []entity.Dataset
	if err := r.db.WithContext(ctx).Where("kecamatan LIKE ?", "%"+kecamatan+"%").Find(&datasets).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("dataset not found")
		}
		return nil, errs.NewInternalServerError("failed to get datasets by kecamatan")
	}
	return datasets, nil
}

func (r *datasetRepository) GetDatasetByKelurahan(ctx context.Context, kelurahan string) ([]entity.Dataset, errs.ErrMessage) {
	var datasets []entity.Dataset
	if err := r.db.WithContext(ctx).Where("kelurahan LIKE ?", "%"+kelurahan+"%").Find(&datasets).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("dataset not found")
		}
		return nil, errs.NewInternalServerError("failed to get datasets by kelurahan")
	}
	return datasets, nil
}

func (r *datasetRepository) GetDatasetByCategory(ctx context.Context, category string) ([]entity.Dataset, errs.ErrMessage) {
	var datasets []entity.Dataset
	if err := r.db.WithContext(ctx).Where("category LIKE ?", "%"+category+"%").Find(&datasets).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("dataset not found")
		}
		return nil, errs.NewInternalServerError("failed to get datasets by category")
	}
	return datasets, nil
}

func (r *datasetRepository) UpdateDataset(ctx context.Context, dataset entity.Dataset) errs.ErrMessage {
	if err := r.db.WithContext(ctx).Model(&entity.Dataset{}).Where("id = ?", dataset.ID).Updates(dataset).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError("dataset not found")
		}
		return errs.NewInternalServerError("failed to update dataset")
	}
	return nil
}

func (r *datasetRepository) DeleteDataset(ctx context.Context, id int) errs.ErrMessage {
	var dataset entity.Dataset
	fmt.Println("Attempting to hard delete dataset with ID:", id)

	// Cek apakah dataset dengan ID tersebut ada
	if err := r.db.WithContext(ctx).First(&dataset, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError("Dataset not found")
		}
		return errs.NewInternalServerError("Failed to fetch dataset for deletion")
	}

	// Hard delete dataset jika ditemukan
	if err := r.db.WithContext(ctx).Unscoped().Delete(&dataset).Error; err != nil {
		fmt.Println("Error deleting dataset:", err)
		return errs.NewInternalServerError("Failed to delete dataset")
	}

	fmt.Println("Dataset hard deleted successfully")
	return nil
}

func (r *datasetRepository) GetAllDatasets(ctx context.Context) ([]entity.Dataset, errs.ErrMessage) {
	var datasets []entity.Dataset
	if err := r.db.WithContext(ctx).Find(&datasets).Error; err != nil {
		return nil, errs.NewInternalServerError("failed to retrieve datasets")
	}
	return datasets, nil
}

func (r *datasetRepository) GetDatasetsByDistance(ctx context.Context, latitude, longitude, distance float64) ([]entity.Dataset, errs.ErrMessage) {
    var datasets []entity.Dataset
    query := `
    SELECT 
        id, name, category, kecamatan, kelurahan, created_at, updated_at,
        ST_Distance_Sphere(
            point(longitude, latitude), 
            point(?, ?)
        ) AS distance
    FROM datasets
    HAVING distance <= ?
    ORDER BY distance
    LIMIT 25;
    `
    
    err := r.db.Raw(query, longitude, latitude, distance).Scan(&datasets).Error
    if err != nil {
        fmt.Println("Error running query:", err)
        return nil, errs.NewInternalServerError("failed to retrieve datasets")
    }

    return datasets, nil
}