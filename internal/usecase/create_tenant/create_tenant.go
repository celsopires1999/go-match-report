package create_tenant

import (
	"context"

	"github.com/celsopires1999/matchreport/internal/entity"
	"github.com/celsopires1999/matchreport/internal/gateway"
)

type InputDTO struct {
	Name string `json:"name"`
}

type OutputDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UseCase struct {
	TenantGateway gateway.TenantRepository
}

func NewCreateTenantUseCase(tenantGateway gateway.TenantRepository) *UseCase {
	return &UseCase{
		TenantGateway: tenantGateway,
	}
}

func (uc *UseCase) Execute(ctx context.Context, input InputDTO) (*OutputDTO, error) {
	tenant, err := entity.NewTenant(input.Name)
	if err != nil {
		return nil, err
	}
	err = uc.TenantGateway.Create(context.Background(), tenant)
	if err != nil {
		return nil, err
	}

	output := &OutputDTO{
		ID:   tenant.ID,
		Name: tenant.Name,
	}
	return output, nil
}
