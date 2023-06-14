package webserver

import (
	"log"

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

func (s *WebServer) RegisterRoutes() {

}

func (s *WebServer) Start() {
	if err := s.Router.Run(s.WebServerPort); err != nil {
		log.Fatal(err)
	}
}
