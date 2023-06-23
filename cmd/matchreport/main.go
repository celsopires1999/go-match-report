package main

import (
	"database/sql"
	"log"

	"github.com/celsopires1999/matchreport/configs"
	"github.com/celsopires1999/matchreport/internal/infra/repository"
	"github.com/celsopires1999/matchreport/internal/infra/web"
	"github.com/celsopires1999/matchreport/internal/infra/web/webserver"
	"github.com/celsopires1999/matchreport/internal/usecase/create_tenant"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("Match Report as of 22-06-2023 at 22:08")

	configs := configs.LoadConfig(".")
	conn, err := sql.Open(configs.DBDriver, configs.DBConn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	handler := initDependencies(conn)
	ws := webserver.NewWebServer(":" + configs.WebServerPort)
	ws.RegisterRoutes(handler)
	ws.Start()
}

func initDependencies(conn *sql.DB) *web.WebTenantHandler {
	tenantRepository := repository.NewTenantRepositoryDb(conn)
	createTenantUseCase := create_tenant.NewCreateTenantUseCase(tenantRepository)
	return web.NewWebTenantHandler(createTenantUseCase)
}
