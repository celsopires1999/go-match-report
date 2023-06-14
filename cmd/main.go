package main

import (
	"database/sql"

	"github.com/celsopires1999/matchreport/configs"
	"github.com/celsopires1999/matchreport/internal/infra/repository"
	"github.com/celsopires1999/matchreport/internal/infra/web"
	"github.com/celsopires1999/matchreport/internal/infra/web/webserver"
	"github.com/celsopires1999/matchreport/internal/usecase/create_tenant"
	_ "github.com/lib/pq"
)

func main() {
	configs := configs.LoadConfig(".")
	conn, err := sql.Open(configs.DBDriver, configs.DBConn)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	tenantRepository := repository.NewTenantRepositoryDb(conn)
	createTenantUseCase := create_tenant.NewCreateTenantUseCase(tenantRepository)
	handler := web.NewWebTenantHandler(createTenantUseCase)
	ws := webserver.NewWebServer(":8888")
	ws.Router.POST("/tenants", handler.CreateTenant)
	ws.Start()
}
