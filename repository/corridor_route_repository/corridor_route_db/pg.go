package corridor_route_db

import (
	"15min-city/entity"
	"15min-city/pkg/errs"
	"15min-city/repository/corridor_route_repository"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type corridorRouteRepository struct {
	db *gorm.DB
}

// NewCorridorRouteRepository creates a new instance of CorridorRouteRepository.
func NewCorridorRouteRepository(db *gorm.DB) corridor_route_repository.CorridorRouteRepository {
	return &corridorRouteRepository{
		db: db,
	}
}

func (r *corridorRouteRepository) CreateCorridorRoute(ctx context.Context, corridorRoute entity.Corridor_Route) (*entity.Corridor_Route, errs.ErrMessage) {
	if err := r.db.WithContext(ctx).Create(&corridorRoute).Error; err != nil {
		return nil, errs.NewInternalServerError("failed to create corridor route")
	}
	return &corridorRoute, nil
}

func (r *corridorRouteRepository) GetCorridorRouteByID(ctx context.Context, id uint) (*entity.Corridor_Route, errs.ErrMessage) {
	var corridorRoute entity.Corridor_Route
	if err := r.db.WithContext(ctx).First(&corridorRoute, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("corridor route not found")
		}
		return nil, errs.NewInternalServerError("failed to get corridor route")
	}
	return &corridorRoute, nil
}

func (r *corridorRouteRepository) GetCorridorRouteByName(ctx context.Context, name string) ([]entity.Corridor_Route, errs.ErrMessage) {
	var corridorRoute []entity.Corridor_Route
	if err := r.db.WithContext(ctx).Where("name LIKE ?", "%"+name+"%").Find(&corridorRoute).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("corridor route not found")
		}
		return nil, errs.NewInternalServerError("failed to get corridor route by name")
	}
	return corridorRoute, nil
}

func (r *corridorRouteRepository) GetCorridorRouteByRoute(ctx context.Context, route string) ([]entity.Corridor_Route, errs.ErrMessage) {
	var corridorRoute []entity.Corridor_Route
	if err := r.db.WithContext(ctx).Where("route LIKE ?", "%"+route+"%").Find(&corridorRoute).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("corridor route not found")
		}
		return nil, errs.NewInternalServerError("failed to get corridor route by route")
	}
	return corridorRoute, nil
}

func (r *corridorRouteRepository) GetCorridorRouteByDirection(ctx context.Context, direction string) ([]entity.Corridor_Route, errs.ErrMessage) {
	var corridorRoute []entity.Corridor_Route
	if err := r.db.WithContext(ctx).Where("direction LIKE ?", "%"+direction+"%").Find(&corridorRoute).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("corridor route not found")
		}
		return nil, errs.NewInternalServerError("failed to get corridor route by direction")
	}
	return corridorRoute, nil
}

func (r *corridorRouteRepository) UpdateCorridorRoute(ctx context.Context, corridorRoute entity.Corridor_Route) errs.ErrMessage {
	if err := r.db.WithContext(ctx).Model(&entity.Corridor_Route{}).Where("id = ?", corridorRoute.ID).Updates(corridorRoute).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError("corridor route not found")
		}
		return errs.NewInternalServerError("failed to update corridor route")
	}
	return nil
}

func (r *corridorRouteRepository) DeleteCorridorRoute(ctx context.Context, id uint) errs.ErrMessage {
	var corridorRoute entity.Corridor_Route
	fmt.Println("Attempting to hard delete corridor route with ID:", id)

	// Check if the corridor route with the given ID exists
	if err := r.db.WithContext(ctx).First(&corridorRoute, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError("Corridor Route not found")
		}
		return errs.NewInternalServerError("Failed to fetch corridor route for deletion")
	}

	// Hard delete the corridor route if found
	if err := r.db.WithContext(ctx).Unscoped().Delete(&corridorRoute).Error; err != nil {
		fmt.Println("Error deleting corridor route:", err)
		return errs.NewInternalServerError("Failed to delete corridor route")
	}

	fmt.Println("Corridor Route hard deleted successfully")
	return nil
}

func (r *corridorRouteRepository) GetAllCorridorRoutes(ctx context.Context) ([]entity.Corridor_Route, errs.ErrMessage) {
	var corridorRoutes []entity.Corridor_Route
	if err := r.db.WithContext(ctx).Find(&corridorRoutes).Error; err != nil {
		return nil, errs.NewInternalServerError("failed to retrieve corridor routes")
	}
	return corridorRoutes, nil
}
