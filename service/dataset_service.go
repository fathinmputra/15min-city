package service

import (
	"15min-city/dto"
	"15min-city/entity"
	"15min-city/pkg/errs"
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
	GetDatasetByName(ctx context.Context, name string) (*dto.GetDatasetByNameResponse, errs.ErrMessage)
	GetDatasetByCategory(ctx context.Context, category string) ([]dto.GetDatasetByCategoryResponse, errs.ErrMessage)
	UpdateDataset(ctx context.Context, id int, datasetPayload dto.UpdateDatasetRequest) (*dto.UpdateDatasetResponse, errs.ErrMessage)
	DeleteDataset(ctx context.Context, id int) (*dto.DeleteDatasetResponse, errs.ErrMessage)
	GetAllDatasets(ctx context.Context) (*dto.GetAllDatasetsResponse, errs.ErrMessage)
}

func NewDatasetService(datasetRepo dataset_repository.DatasetRepository) DatasetService {
	return &datasetService{
		datasetRepo: datasetRepo,
	}
}

func (d *datasetService) CreateDataset(ctx context.Context, datasetPayload dto.CreateDatasetRequest) (*dto.CreateDatasetResponse, errs.ErrMessage) {
	dataset := entity.Dataset{
		LocationId:     datasetPayload.LocationId,
		Name:           datasetPayload.Name,
		Latitude:       datasetPayload.Latitude,
		Longitude:      datasetPayload.Longitude,
		BusinessStatus: datasetPayload.BusinessStatus,
		Kelurahan:      datasetPayload.Kelurahan,
		Kota:           datasetPayload.Kota,
		Category:       datasetPayload.Category,
	}

	createdDataset, err := d.datasetRepo.CreateDataset(ctx, dataset)
	if err != nil {
		return nil, err
	}

	response := dto.CreateDatasetResponse{
		Status:         http.StatusCreated,
		ID:             int(createdDataset.ID),
		LocationId:     createdDataset.LocationId,
		Name:           createdDataset.Name,
		Latitude:       createdDataset.Latitude,
		Longitude:      createdDataset.Longitude,
		BusinessStatus: createdDataset.BusinessStatus,
		Kelurahan:      createdDataset.Kelurahan,
		Kota:           createdDataset.Kota,
		Category:       createdDataset.Category,
		CreatedAt:      createdDataset.CreatedAt,
		UpdatedAt:      createdDataset.UpdatedAt,
	}

	return &response, nil
}

func (d *datasetService) GetDatasetByID(ctx context.Context, id int) (*dto.GetDatasetByIDResponse, errs.ErrMessage) {
	dataset, err := d.datasetRepo.GetDatasetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	response := dto.GetDatasetByIDResponse{
		Status:         http.StatusOK,
		ID:             int(dataset.ID),
		LocationId:     dataset.LocationId,
		Name:           dataset.Name,
		Latitude:       dataset.Latitude,
		Longitude:      dataset.Longitude,
		BusinessStatus: dataset.BusinessStatus,
		Kelurahan:      dataset.Kelurahan,
		Kota:           dataset.Kota,
		Category:       dataset.Category,
		CreatedAt:      dataset.CreatedAt,
		UpdatedAt:      dataset.UpdatedAt,
	}

	return &response, nil
}

func (d *datasetService) GetDatasetByName(ctx context.Context, name string) (*dto.GetDatasetByNameResponse, errs.ErrMessage) {
	dataset, err := d.datasetRepo.GetDatasetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	response := dto.GetDatasetByNameResponse{
		Status:         http.StatusOK,
		ID:             int(dataset.ID),
		LocationId:     dataset.LocationId,
		Name:           dataset.Name,
		Latitude:       dataset.Latitude,
		Longitude:      dataset.Longitude,
		BusinessStatus: dataset.BusinessStatus,
		Kelurahan:      dataset.Kelurahan,
		Kota:           dataset.Kota,
		Category:       dataset.Category,
		CreatedAt:      dataset.CreatedAt,
		UpdatedAt:      dataset.UpdatedAt,
	}

	return &response, nil
}

func (d *datasetService) GetDatasetByCategory(ctx context.Context, category string) ([]dto.GetDatasetByCategoryResponse, errs.ErrMessage) {
	datasets, err := d.datasetRepo.GetDatasetByCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	var response []dto.GetDatasetByCategoryResponse
	for _, dataset := range datasets {
		response = append(response, dto.GetDatasetByCategoryResponse{
			ID:             int(dataset.ID),
			LocationId:     dataset.LocationId,
			Name:           dataset.Name,
			Latitude:       dataset.Latitude,
			Longitude:      dataset.Longitude,
			BusinessStatus: dataset.BusinessStatus,
			Kelurahan:      dataset.Kelurahan,
			Kota:           dataset.Kota,
			Category:       dataset.Category,
			CreatedAt:      dataset.CreatedAt,
			UpdatedAt:      dataset.UpdatedAt,
		})
	}

	return response, nil
}

func (d *datasetService) UpdateDataset(ctx context.Context, id int, datasetPayload dto.UpdateDatasetRequest) (*dto.UpdateDatasetResponse, errs.ErrMessage) {
	// Ambil dataset yang ada
	existingDataset, err := d.datasetRepo.GetDatasetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update dataset dengan data baru
	existingDataset.LocationId = datasetPayload.LocationId
	existingDataset.Name = datasetPayload.Name
	existingDataset.Latitude = datasetPayload.Latitude
	existingDataset.Longitude = datasetPayload.Longitude
	existingDataset.BusinessStatus = datasetPayload.BusinessStatus
	existingDataset.Kelurahan = datasetPayload.Kelurahan
	existingDataset.Kota = datasetPayload.Kota
	existingDataset.Category = datasetPayload.Category

	// Simpan perubahan
	err = d.datasetRepo.UpdateDataset(ctx, *existingDataset)
	if err != nil {
		return nil, err
	}

	// Siapkan response
	response := dto.UpdateDatasetResponse{
		Status:         http.StatusOK,
		ID:             int(existingDataset.ID),
		LocationId:     existingDataset.LocationId,
		Name:           existingDataset.Name,
		Latitude:       existingDataset.Latitude,
		Longitude:      existingDataset.Longitude,
		BusinessStatus: existingDataset.BusinessStatus,
		Kelurahan:      existingDataset.Kelurahan,
		Kota:           existingDataset.Kota,
		Category:       existingDataset.Category,
		UpdatedAt:      existingDataset.UpdatedAt,
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

func (d *datasetService) GetAllDatasets(ctx context.Context) (*dto.GetAllDatasetsResponse, errs.ErrMessage) {
	datasets, err := d.datasetRepo.GetAllDatasets(ctx)
	if err != nil {
		return nil, err
	}

	// Prepare the response with the list of datasets
	responseDatasets := []dto.DatasetInfo{}
	for _, dataset := range datasets {
		responseDatasets = append(responseDatasets, dto.DatasetInfo{
			ID:             int(dataset.ID),
			LocationId:     dataset.LocationId,
			Name:           dataset.Name,
			Latitude:       dataset.Latitude,
			Longitude:      dataset.Longitude,
			BusinessStatus: dataset.BusinessStatus,
			Kelurahan:      dataset.Kelurahan,
			Kota:           dataset.Kota,
			Category:       dataset.Category,
			CreatedAt:      dataset.CreatedAt,
			UpdatedAt:      dataset.UpdatedAt,
		})
	}

	response := &dto.GetAllDatasetsResponse{
		Status:   http.StatusOK,
		Message:  "OK",
		Datasets: responseDatasets,
	}

	return response, nil
}
