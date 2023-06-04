package entity

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type Tenant struct {
	Id   string `valid:"uuid"`
	Name string `valid:"required~Name is required"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewTenant(name string) (*Tenant, error) {
	return buildTenant(uuid.New().String(), name)
}

func NewTenantWithId(id string, name string) (*Tenant, error) {
	return buildTenant(id, name)
}

func buildTenant(id string, name string) (*Tenant, error) {
	tenant := &Tenant{
		Id:   uuid.New().String(),
		Name: name,
	}

	if err := tenant.validate(); err != nil {
		return nil, err
	}

	return tenant, nil
}

func (t *Tenant) validate() error {
	if _, err := govalidator.ValidateStruct(t); err != nil {
		return err
	}
	return nil
}
