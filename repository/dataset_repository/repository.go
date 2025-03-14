package dataset_repository

import (
	"15min-city/entity"
	"15min-city/pkg/errs"
	"context"
)

type DatasetRepository interface {
	CreateDataset(ctx context.Context, dataset entity.Dataset) (*entity.Dataset, errs.ErrMessage)
	GetDatasetByID(ctx context.Context, id int) (*entity.Dataset, errs.ErrMessage)
	GetDatasetByName(ctx context.Context, name string) ([]entity.Dataset, errs.ErrMessage)
	GetDatasetByKecamatan(ctx context.Context, kecamatan string) ([]entity.Dataset, errs.ErrMessage)
	GetDatasetByKelurahan(ctx context.Context, kelurahan string) ([]entity.Dataset, errs.ErrMessage)
	GetDatasetByCategory(ctx context.Context, category string) ([]entity.Dataset, errs.ErrMessage)
	UpdateDataset(ctx context.Context, dataset entity.Dataset) errs.ErrMessage
	DeleteDataset(ctx context.Context, id int) errs.ErrMessage
	GetAllDatasets(ctx context.Context) ([]entity.Dataset, errs.ErrMessage)
	GetDatasetsByDistance(ctx context.Context, latitude, longitude, distance float64) ([]entity.Dataset, errs.ErrMessage)
}
