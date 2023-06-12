package create_tenant

import (
	"context"

	"github.com/celsopires1999/matchreport/internal/entity"
	"github.com/celsopires1999/matchreport/internal/gateway"
)

type CreateTenantInputDTO struct {
	Name string
}

type CreateTenantOutputDTO struct {
	ID   string
	Name string
}

type CreateTenantUseCase struct {
	TenantGateway gateway.TenantRepository
}

func NewCreateTenantUseCase(tenantGateway gateway.TenantRepository) *CreateTenantUseCase {
	return &CreateTenantUseCase{
		TenantGateway: tenantGateway,
	}
}

func (uc *CreateTenantUseCase) Execute(input CreateTenantInputDTO) (*CreateTenantOutputDTO, error) {
	tenant, err := entity.NewTenant(input.Name)
	if err != nil {
		return nil, err
	}
	err = uc.TenantGateway.Create(context.Background(), tenant)
	if err != nil {
		return nil, err
	}

	output := &CreateTenantOutputDTO{
		ID:   tenant.ID,
		Name: tenant.Name,
	}
	return output, nil
}
