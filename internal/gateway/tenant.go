package gateway

import (
	"context"

	"github.com/celsopires1999/matchreport/internal/entity"
)

type TenantRepository interface {
	FindById(ctx context.Context, id string) (*entity.Tenant, error)
	Create(ctx context.Context, tenant *entity.Tenant) error
	Update(ctx context.Context, tenant *entity.Tenant) error
	Delete(ctx context.Context, id string) error
}
