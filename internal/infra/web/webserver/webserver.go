package webserver

import (
	"log"

	"github.com/celsopires1999/matchreport/internal/infra/web"
	"github.com/gin-contrib/cors"
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

func headerCors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
}

func (s *WebServer) RegisterRoutes(handler *web.WebTenantHandler) {
	s.Router.GET("/health", headerCors, web.HealthHandler)
	s.Router.POST("/tenants", headerCors, handler.CreateTenant)
}

func (s *WebServer) Start() {

	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"*"}

	s.Router.Use(cors.New(corsConfig))

	if err := s.Router.Run(s.WebServerPort); err != nil {
		log.Fatal(err)
	}
}
