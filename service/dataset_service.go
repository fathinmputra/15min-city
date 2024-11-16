package service

import (
	"15min-city/dto"
	"15min-city/entity"
	"15min-city/pkg/errs"
	"15min-city/pkg/helpers"
	"15min-city/repository/dataset_repository"
	"context"
	"net/http"
)

type datasetService struct {
	datasetRepo dataset_repository.DatasetRepository
}

type DatasetService interface {
	CreateDataset(ctx context.Context, datasetPayload dto.CreateDatasetRequest) (*dto.CreateDatasetResponse, errs.ErrMessage)
	GetDatasetByID(ctx context.Context, id int) (*dto.GetDatasetByIDResponse, errs.ErrMessage)
	GetDatasetByName(ctx context.Context, name string) ([]dto.GetDatasetByNameResponse, errs.ErrMessage)
	GetDatasetByKecamatan(ctx context.Context, kecamatan string) ([]dto.GetDatasetByKecamatanResponse, errs.ErrMessage)
	GetDatasetByKelurahan(ctx context.Context, kelurahan string) ([]dto.GetDatasetByKelurahanResponse, errs.ErrMessage)
	GetDatasetByCategory(ctx context.Context, category string) ([]dto.GetDatasetByCategoryResponse, errs.ErrMessage)
	UpdateDataset(ctx context.Context, id int, datasetPayload dto.UpdateDatasetRequest) (*dto.UpdateDatasetResponse, errs.ErrMessage)
	DeleteDataset(ctx context.Context, id int) (*dto.DeleteDatasetResponse, errs.ErrMessage)
	GetAllDatasets(ctx context.Context) ([]dto.GetAllDatasetsResponse, errs.ErrMessage)
}

func NewDatasetService(datasetRepo dataset_repository.DatasetRepository) DatasetService {
	return &datasetService{
		datasetRepo: datasetRepo,
	}
}

func (d *datasetService) CreateDataset(ctx context.Context, datasetPayload dto.CreateDatasetRequest) (*dto.CreateDatasetResponse, errs.ErrMessage) {
	dataset := entity.Dataset{
		Name:      datasetPayload.Name,
		Latitude:  datasetPayload.Latitude,
		Longitude: datasetPayload.Longitude,
		Category:  datasetPayload.Category,
	}

	// Mendapatkan Kecamatan dan Kelurahan menggunakan reverse geocoding
	kecamatan, kelurahan, err := helpers.ReverseGeocode(dataset.Latitude, dataset.Longitude)
	if err != nil {
		return nil, errs.NewInternalServerError("Failed to fetch geocoding data")
	}

	dataset.Kecamatan = kecamatan
	dataset.Kelurahan = kelurahan

	createdDataset, err := d.datasetRepo.CreateDataset(ctx, dataset)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	response := dto.CreateDatasetResponse{
		// Status:    http.StatusCreated,
		ID:        int(createdDataset.ID),
		Name:      createdDataset.Name,
		Latitude:  createdDataset.Latitude,
		Longitude: createdDataset.Longitude,
		Category:  createdDataset.Category,
		Kecamatan: createdDataset.Kecamatan,
		Kelurahan: createdDataset.Kelurahan,
		CreatedAt: createdDataset.CreatedAt,
		UpdatedAt: createdDataset.UpdatedAt,
	}

	return &response, nil
}

func (d *datasetService) GetDatasetByID(ctx context.Context, id int) (*dto.GetDatasetByIDResponse, errs.ErrMessage) {
	dataset, err := d.datasetRepo.GetDatasetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	response := dto.GetDatasetByIDResponse{
		// Status:    http.StatusOK,
		ID:        int(dataset.ID),
		Name:      dataset.Name,
		Latitude:  dataset.Latitude,
		Longitude: dataset.Longitude,
		Category:  dataset.Category,
		Kecamatan: dataset.Kecamatan,
		Kelurahan: dataset.Kelurahan,
		CreatedAt: dataset.CreatedAt,
		UpdatedAt: dataset.UpdatedAt,
	}

	return &response, nil
}

func (d *datasetService) GetDatasetByName(ctx context.Context, name string) ([]dto.GetDatasetByNameResponse, errs.ErrMessage) {
	datasets, err := d.datasetRepo.GetDatasetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	var response []dto.GetDatasetByNameResponse
	for _, dataset := range datasets {
		response = append(response, dto.GetDatasetByNameResponse{
			// Status:    http.StatusOK,
			ID:        int(dataset.ID),
			Name:      dataset.Name,
			Latitude:  dataset.Latitude,
			Longitude: dataset.Longitude,
			Category:  dataset.Category,
			Kecamatan: dataset.Kecamatan,
			Kelurahan: dataset.Kelurahan,
			CreatedAt: dataset.CreatedAt,
			UpdatedAt: dataset.UpdatedAt,
		})
	}

	return response, nil
}

func (d *datasetService) GetDatasetByKecamatan(ctx context.Context, name string) ([]dto.GetDatasetByKecamatanResponse, errs.ErrMessage) {
	datasets, err := d.datasetRepo.GetDatasetByKecamatan(ctx, name)
	if err != nil {
		return nil, err
	}

	var response []dto.GetDatasetByKecamatanResponse
	for _, dataset := range datasets {
		response = append(response, dto.GetDatasetByKecamatanResponse{
			// Status:    http.StatusOK,
			ID:        int(dataset.ID),
			Name:      dataset.Name,
			Latitude:  dataset.Latitude,
			Longitude: dataset.Longitude,
			Category:  dataset.Category,
			Kecamatan: dataset.Kecamatan,
			Kelurahan: dataset.Kelurahan,
			CreatedAt: dataset.CreatedAt,
			UpdatedAt: dataset.UpdatedAt,
		})
	}

	return response, nil
}

func (d *datasetService) GetDatasetByKelurahan(ctx context.Context, name string) ([]dto.GetDatasetByKelurahanResponse, errs.ErrMessage) {
	datasets, err := d.datasetRepo.GetDatasetByKelurahan(ctx, name)
	if err != nil {
		return nil, err
	}

	var response []dto.GetDatasetByKelurahanResponse
	for _, dataset := range datasets {
		response = append(response, dto.GetDatasetByKelurahanResponse{
			// Status:    http.StatusOK,
			ID:        int(dataset.ID),
			Name:      dataset.Name,
			Latitude:  dataset.Latitude,
			Longitude: dataset.Longitude,
			Category:  dataset.Category,
			Kecamatan: dataset.Kecamatan,
			Kelurahan: dataset.Kelurahan,
			CreatedAt: dataset.CreatedAt,
			UpdatedAt: dataset.UpdatedAt,
		})
	}

	return response, nil
}

func (d *datasetService) GetDatasetByCategory(ctx context.Context, name string) ([]dto.GetDatasetByCategoryResponse, errs.ErrMessage) {
	datasets, err := d.datasetRepo.GetDatasetByCategory(ctx, name)
	if err != nil {
		return nil, err
	}

	var response []dto.GetDatasetByCategoryResponse
	for _, dataset := range datasets {
		response = append(response, dto.GetDatasetByCategoryResponse{
			// Status:    http.StatusOK,
			ID:        int(dataset.ID),
			Name:      dataset.Name,
			Latitude:  dataset.Latitude,
			Longitude: dataset.Longitude,
			Category:  dataset.Category,
			Kecamatan: dataset.Kecamatan,
			Kelurahan: dataset.Kelurahan,
			CreatedAt: dataset.CreatedAt,
			UpdatedAt: dataset.UpdatedAt,
		})
	}

	return response, nil
}

func (d *datasetService) UpdateDataset(ctx context.Context, id int, datasetPayload dto.UpdateDatasetRequest) (*dto.UpdateDatasetResponse, errs.ErrMessage) {
	// Retrieve the existing dataset
	existingDataset, err := d.datasetRepo.GetDatasetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update dataset with new data
	if datasetPayload.Name != "" {
		existingDataset.Name = datasetPayload.Name
	}
	if datasetPayload.Latitude != 0 {
		existingDataset.Latitude = datasetPayload.Latitude
	}
	if datasetPayload.Longitude != 0 {
		existingDataset.Longitude = datasetPayload.Longitude
	}
	if datasetPayload.Category != "" {
		existingDataset.Category = datasetPayload.Category
	}
	if datasetPayload.Kecamatan != "" {
		existingDataset.Kecamatan = datasetPayload.Kecamatan
	}
	if datasetPayload.Kelurahan != "" {
		existingDataset.Kelurahan = datasetPayload.Kelurahan
	}

	// Save changes
	err = d.datasetRepo.UpdateDataset(ctx, *existingDataset)
	if err != nil {
		return nil, err
	}

	// Prepare response
	response := dto.UpdateDatasetResponse{
		// Status:    http.StatusOK,
		ID:        int(existingDataset.ID),
		Name:      existingDataset.Name,
		Latitude:  existingDataset.Latitude,
		Longitude: existingDataset.Longitude,
		Category:  existingDataset.Category,
		Kecamatan: existingDataset.Kecamatan,
		Kelurahan: existingDataset.Kelurahan,
		UpdatedAt: existingDataset.UpdatedAt,
	}

	return &response, nil
}

func (d *datasetService) DeleteDataset(ctx context.Context, id int) (*dto.DeleteDatasetResponse, errs.ErrMessage) {
	err := d.datasetRepo.DeleteDataset(ctx, id)
	if err != nil {
		return nil, err
	}

	response := dto.DeleteDatasetResponse{
		Status:  http.StatusOK,
		Message: "Dataset successfully deleted",
	}

	return &response, nil
}

func (d *datasetService) GetAllDatasets(ctx context.Context) ([]dto.GetAllDatasetsResponse, errs.ErrMessage) {
	// Panggil repository untuk mendapatkan semua dataset
	datasets, err := d.datasetRepo.GetAllDatasets(ctx)
	if err != nil {
		return nil, err
	}

	// Siapkan slice response
	var response []dto.GetAllDatasetsResponse
	for _, dataset := range datasets {
		response = append(response, dto.GetAllDatasetsResponse{
			// Status:    http.StatusOK,
			ID:        int(dataset.ID),
			Name:      dataset.Name,
			Latitude:  dataset.Latitude,
			Longitude: dataset.Longitude,
			Category:  dataset.Category,
			Kecamatan: dataset.Kecamatan,
			Kelurahan: dataset.Kelurahan,
			CreatedAt: dataset.CreatedAt,
			UpdatedAt: dataset.UpdatedAt,
		})
	}

	return response, nil
}
