package webserver

import (
	"log"

	"github.com/celsopires1999/matchreport/internal/infra/web"
	"github.com/gin-gonic/gin"
)

type WebServer struct {
	Router        *gin.Engine
	WebServerPort string
}

func NewWebServer(webServerPort string) *WebServer {
	return &WebServer{
		Router:        gin.Default(),
		WebServerPort: webServerPort,
	}
}

func (s *WebServer) RegisterRoutes(handler *web.WebTenantHandler) {
	s.Router.GET("/health", web.HealthHandler)
	s.Router.POST("/tenants", handler.CreateTenant)
}

func (s *WebServer) Start() {
	if err := s.Router.Run(s.WebServerPort); err != nil {
		log.Fatal(err)
	}
}
