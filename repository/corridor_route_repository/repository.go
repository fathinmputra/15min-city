package corridor_route_repository

import (
	"15min-city/entity"
	"15min-city/pkg/errs"
	"context"
)

type CorridorRouteRepository interface {
	CreateCorridorRoute(ctx context.Context, corridorRoute entity.Corridor_Route) (*entity.Corridor_Route, errs.ErrMessage)
	GetCorridorRouteByID(ctx context.Context, id uint) (*entity.Corridor_Route, errs.ErrMessage)
	GetCorridorRouteByName(ctx context.Context, name string) (*entity.Corridor_Route, errs.ErrMessage)
	UpdateCorridorRoute(ctx context.Context, corridorRoute entity.Corridor_Route) errs.ErrMessage
	DeleteCorridorRoute(ctx context.Context, id uint) errs.ErrMessage
	GetAllCorridorRoutes(ctx context.Context) ([]entity.Corridor_Route, errs.ErrMessage)
}
