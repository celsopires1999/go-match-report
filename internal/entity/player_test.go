package entity_test

import (
	"testing"

	"github.com/celsopires1999/matchreport/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewPlayer(t *testing.T) {
	t.Run("should create a new player", func(t *testing.T) {
		name := "Joe Doe"
		tenant, _ := entity.NewTenant("Veterans' Football League")
		player, _ := entity.NewPlayer(name, tenant)
		assert.NotNil(t, player)
		assert.NotEmpty(t, player.Id)
		assert.Equal(t, name, player.Name)
		assert.Equal(t, tenant, player.Tenant)
	})

	t.Run("should return error if name is empty", func(t *testing.T) {
		name := ""
		tenant, _ := entity.NewTenant("Veterans' Football League")
		player, err := entity.NewPlayer(name, tenant)
		assert.Nil(t, player)
		assert.Error(t, err)
		assert.Equal(t, "Name is required", err.Error())
	})
	t.Run("should return error if tenant is empty", func(t *testing.T) {
		name := "Joe Doe"
		tenant, _ := entity.NewTenant("")
		player, err := entity.NewPlayer(name, tenant)
		assert.Nil(t, player)
		assert.Error(t, err)
		assert.Equal(t, "Tenant is required", err.Error())
	})
	t.Run("should return error if name and tenant are empty", func(t *testing.T) {
		name := ""
		tenant, _ := entity.NewTenant("")
		player, err := entity.NewPlayer(name, tenant)
		assert.Nil(t, player)
		assert.Error(t, err)
		assert.Equal(t, "Name is required;Tenant is required", err.Error())
	})
}
