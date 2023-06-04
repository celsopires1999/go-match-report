package entity_test

import (
	"testing"

	"github.com/celsopires1999/matchreport/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewTenant(t *testing.T) {
	t.Run("should create a new tenant", func(t *testing.T) {
		name := "Veterans' Football League"
		tenant, _ := entity.NewTenant(name)
		assert.NotNil(t, tenant)
		assert.NotEmpty(t, tenant.Id)
		assert.Equal(t, name, tenant.Name)
	})
	t.Run("should return error when name is empty", func(t *testing.T) {
		_, err := entity.NewTenant("")
		assert.Error(t, err)
		assert.Equal(t, "Name is required", err.Error())
	})
}
