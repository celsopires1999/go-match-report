package entity

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type Player struct {
	ID     string  `valid:"uuid"`
	Name   string  `valid:"required~Name is required"`
	Tenant *Tenant `valid:"required~Tenant is required"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewPlayer(name string, tenant *Tenant) (*Player, error) {
	player := &Player{
		ID:     uuid.New().String(),
		Name:   name,
		Tenant: tenant,
	}

	if err := player.validate(); err != nil {
		return nil, err
	}
	return player, nil
}

func (p *Player) validate() error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return err
	}
	return nil
}
