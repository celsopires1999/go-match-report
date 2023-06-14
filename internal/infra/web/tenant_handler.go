package web

import (
	"github.com/celsopires1999/matchreport/internal/usecase/create_tenant"
	"github.com/gin-gonic/gin"
)

type WebTenantHandler struct {
	CreateTenantUseCase create_tenant.UseCase
}

func NewWebTenantHandler(createTenantUseCase *create_tenant.UseCase) *WebTenantHandler {
	return &WebTenantHandler{
		CreateTenantUseCase: *createTenantUseCase,
	}
}

func (h *WebTenantHandler) CreateTenant(c *gin.Context) {
	var input create_tenant.InputDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, err.Error())
		return
	}

	output, err := h.CreateTenantUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		c.JSON(400, err.Error())
	}

	c.JSON(200, output)
}
