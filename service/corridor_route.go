package service

import (
	"15min-city/dto"
	"15min-city/entity"
	"15min-city/pkg/errs"
	"15min-city/repository/corridor_route_repository"
	"context"
	"net/http"
)

type corridorRouteService struct {
	corridorRouteRepo corridor_route_repository.CorridorRouteRepository
}

type CorridorRouteService interface {
	CreateCorridorRoute(ctx context.Context, corridorRoutePayload dto.CreateCorridorRouteRequest) (*dto.CreateCorridorRouteResponse, errs.ErrMessage)
	GetCorridorRouteByID(ctx context.Context, id uint) (*dto.GetCorridorRouteByIDResponse, errs.ErrMessage)
	GetCorridorRouteByName(ctx context.Context, name string) (*dto.GetCorridorRouteByNameResponse, errs.ErrMessage)
	UpdateCorridorRoute(ctx context.Context, id uint, corridorRoutePayload dto.UpdateCorridorRouteRequest) (*dto.UpdateCorridorRouteResponse, errs.ErrMessage)
	DeleteCorridorRoute(ctx context.Context, id uint) (*dto.DeleteCorridorRouteResponse, errs.ErrMessage)
	GetAllCorridorRoutes(ctx context.Context) (*dto.GetAllCorridorRoutesResponse, errs.ErrMessage)
}

func NewCorridorRouteService(corridorRouteRepo corridor_route_repository.CorridorRouteRepository) CorridorRouteService {
	return &corridorRouteService{
		corridorRouteRepo: corridorRouteRepo,
	}
}

func (r *corridorRouteService) CreateCorridorRoute(ctx context.Context, corridorRoutePayload dto.CreateCorridorRouteRequest) (*dto.CreateCorridorRouteResponse, errs.ErrMessage) {
	corridorRoute := entity.Corridor_Route{
		Name:      corridorRoutePayload.Name,
		Latitude:  corridorRoutePayload.Latitude,
		Longitude: corridorRoutePayload.Longitude,
		Route:     corridorRoutePayload.Route,
		Direction: corridorRoutePayload.Direction,
	}

	createdCorridorRoute, err := r.corridorRouteRepo.CreateCorridorRoute(ctx, corridorRoute)
	if err != nil {
		return nil, err
	}

	response := dto.CreateCorridorRouteResponse{
		Status:    http.StatusCreated,
		ID:        createdCorridorRoute.ID,
		Name:      createdCorridorRoute.Name,
		Latitude:  createdCorridorRoute.Latitude,
		Longitude: createdCorridorRoute.Longitude,
		Route:     createdCorridorRoute.Route,
		Direction: createdCorridorRoute.Direction,
		CreatedAt: createdCorridorRoute.CreatedAt,
		UpdatedAt: createdCorridorRoute.UpdatedAt,
	}

	return &response, nil
}

func (r *corridorRouteService) GetCorridorRouteByID(ctx context.Context, id uint) (*dto.GetCorridorRouteByIDResponse, errs.ErrMessage) {
	corridorRoute, err := r.corridorRouteRepo.GetCorridorRouteByID(ctx, id)
	if err != nil {
		return nil, err
	}

	response := dto.GetCorridorRouteByIDResponse{
		Status:    http.StatusOK,
		ID:        corridorRoute.ID,
		Name:      corridorRoute.Name,
		Latitude:  corridorRoute.Latitude,
		Longitude: corridorRoute.Longitude,
		Route:     corridorRoute.Route,
		Direction: corridorRoute.Direction,
		CreatedAt: corridorRoute.CreatedAt,
		UpdatedAt: corridorRoute.UpdatedAt,
	}

	return &response, nil
}

func (r *corridorRouteService) GetCorridorRouteByName(ctx context.Context, name string) (*dto.GetCorridorRouteByNameResponse, errs.ErrMessage) {
	corridorRoute, err := r.corridorRouteRepo.GetCorridorRouteByName(ctx, name)
	if err != nil {
		return nil, err
	}

	response := dto.GetCorridorRouteByNameResponse{
		Status:    http.StatusOK,
		ID:        corridorRoute.ID,
		Name:      corridorRoute.Name,
		Latitude:  corridorRoute.Latitude,
		Longitude: corridorRoute.Longitude,
		Route:     corridorRoute.Route,
		Direction: corridorRoute.Direction,
		CreatedAt: corridorRoute.CreatedAt,
		UpdatedAt: corridorRoute.UpdatedAt,
	}

	return &response, nil
}

func (r *corridorRouteService) UpdateCorridorRoute(ctx context.Context, id uint, corridorRoutePayload dto.UpdateCorridorRouteRequest) (*dto.UpdateCorridorRouteResponse, errs.ErrMessage) {
	// Retrieve the existing corridor route
	existingCorridorRoute, err := r.corridorRouteRepo.GetCorridorRouteByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update corridor route with new data
	if corridorRoutePayload.Name != "" {
		existingCorridorRoute.Name = corridorRoutePayload.Name
	}
	if corridorRoutePayload.Latitude != 0 {
		existingCorridorRoute.Latitude = corridorRoutePayload.Latitude
	}
	if corridorRoutePayload.Longitude != 0 {
		existingCorridorRoute.Longitude = corridorRoutePayload.Longitude
	}
	if corridorRoutePayload.Route != "" {
		existingCorridorRoute.Route = corridorRoutePayload.Route
	}
	if corridorRoutePayload.Direction != "" {
		existingCorridorRoute.Direction = corridorRoutePayload.Direction
	}

	// Save changes
	err = r.corridorRouteRepo.UpdateCorridorRoute(ctx, *existingCorridorRoute)
	if err != nil {
		return nil, err
	}

	// Prepare response
	response := dto.UpdateCorridorRouteResponse{
		Status:    http.StatusOK,
		ID:        existingCorridorRoute.ID,
		Name:      existingCorridorRoute.Name,
		Latitude:  existingCorridorRoute.Latitude,
		Longitude: existingCorridorRoute.Longitude,
		Route:     existingCorridorRoute.Route,
		Direction: existingCorridorRoute.Direction,
		UpdatedAt: existingCorridorRoute.UpdatedAt,
	}

	return &response, nil
}

func (r *corridorRouteService) DeleteCorridorRoute(ctx context.Context, id uint) (*dto.DeleteCorridorRouteResponse, errs.ErrMessage) {
	err := r.corridorRouteRepo.DeleteCorridorRoute(ctx, id)
	if err != nil {
		return nil, err
	}

	response := dto.DeleteCorridorRouteResponse{
		Status:  http.StatusOK,
		Message: "Corridor Route successfully deleted",
	}

	return &response, nil
}

func (r *corridorRouteService) GetAllCorridorRoutes(ctx context.Context) (*dto.GetAllCorridorRoutesResponse, errs.ErrMessage) {
	corridorRoutes, err := r.corridorRouteRepo.GetAllCorridorRoutes(ctx)
	if err != nil {
		return nil, err
	}

	var response []dto.CorridorRouteInfo
	for _, corridorRoute := range corridorRoutes {
		response = append(response, dto.CorridorRouteInfo{
			ID:        corridorRoute.ID,
			Name:      corridorRoute.Name,
			Latitude:  corridorRoute.Latitude,
			Longitude: corridorRoute.Longitude,
			Route:     corridorRoute.Route,
			Direction: corridorRoute.Direction,
			CreatedAt: corridorRoute.CreatedAt,
			UpdatedAt: corridorRoute.UpdatedAt,
		})
	}

	return &dto.GetAllCorridorRoutesResponse{
		Status:         http.StatusOK,
		Message:        "All corridor routes retrieved successfully",
		CorridorRoutes: response,
	}, nil
}
