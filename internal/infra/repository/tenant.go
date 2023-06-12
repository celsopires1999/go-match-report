package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/celsopires1999/matchreport/internal/entity"
	"github.com/celsopires1999/matchreport/internal/infra/db"
)

type TenantRepositoryDb struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewTenantRepositoryDb(dbt *sql.DB) *TenantRepositoryDb {
	return &TenantRepositoryDb{
		DB:      dbt,
		Queries: db.New(dbt),
	}
}

func (r *TenantRepositoryDb) Create(ctx context.Context, tenant *entity.Tenant) error {
	return r.Queries.CreateTenant(
		ctx,
		db.CreateTenantParams{
			ID:        tenant.ID,
			Name:      tenant.Name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	)
}

func (r *TenantRepositoryDb) Update(ctx context.Context, tenant *entity.Tenant) error {
	return r.Queries.UpdateTenant(
		ctx,
		db.UpdateTenantParams{
			ID:        tenant.ID,
			Name:      tenant.Name,
			UpdatedAt: time.Now(),
		},
	)
}

func (r *TenantRepositoryDb) Delete(ctx context.Context, id string) error {
	return r.Queries.DeleteTenant(
		ctx,
		id,
	)
}

func (r *TenantRepositoryDb) FindById(ctx context.Context, id string) (*entity.Tenant, error) {
	res, err := r.Queries.FindTenantById(ctx, id)

	if err != nil {
		return nil, errors.New("tenant not found")
	}

	return entity.NewTenantWithId(res.ID, res.Name)
}
