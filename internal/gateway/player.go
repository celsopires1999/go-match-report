package gateway

import (
	"context"

	"github.com/celsopires1999/matchreport/internal/entity"
)

type PlayerRepository interface {
	FindPlayerById(ctx context.Context, id string) (*entity.Tenant, error)
	CreatePlayer(ctx context.Context, tenant *entity.Tenant) error
	UpdatePlayer(ctx context.Context, tenant *entity.Tenant) error
	DeletePlayer(ctx context.Context, id string) error
}
