package gateway

import (
	"context"

	"github.com/celsopires1999/matchreport/internal/entity"
)

type TenantRepository interface {
	FindTenantById(ctx context.Context, id string) (*entity.Tenant, error)
	CreateTenant(ctx context.Context, tenant *entity.Tenant) error
	UpdateTenant(ctx context.Context, tenant *entity.Tenant) error
	DeleteTenant(ctx context.Context, id string) error
}
