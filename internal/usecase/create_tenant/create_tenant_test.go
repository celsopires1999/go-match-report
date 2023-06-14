package create_tenant_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/celsopires1999/matchreport/configs"
	"github.com/celsopires1999/matchreport/internal/infra/repository"
	"github.com/celsopires1999/matchreport/internal/usecase/create_tenant"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	MIGRATION_PATH = "file://../../../sql/migrations"
)

type TenantDBTestSuite struct {
	suite.Suite
	db *sql.DB
	m  *migrate.Migrate
}

func (s *TenantDBTestSuite) SetupSuite() {

	configs := configs.LoadConfig("../../../")

	conn, err := sql.Open(configs.DBDriver, configs.DBConn)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	m, err := migrate.New(MIGRATION_PATH, configs.DBConn)
	s.Nil(err)
	s.NotNil(m)

	err = m.Up()
	s.Nil(err)

	db, err := sql.Open(configs.DBDriver, configs.DBConn)

	s.NotNil(db)
	s.Nil(err)
	s.m = m
	s.db = db
}

func (s *TenantDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	err := s.m.Down()
	s.Nil(err)
}

func TestTenantDBTestSuite(t *testing.T) {
	suite.Run(t, new(TenantDBTestSuite))
}

func (s *TenantDBTestSuite) TestCreateTenantUseCase() {
	s.Run("should create tenant", func() {
		name := "Veterans' Football League"
		input := create_tenant.InputDTO{
			Name: name,
		}
		tenantRepo := repository.NewTenantRepositoryDb(s.db)
		tenantUseCase := create_tenant.NewCreateTenantUseCase(tenantRepo)
		ctx := context.Background()
		output, err := tenantUseCase.Execute(ctx, input)
		s.NotNil(output)
		s.Nil(err)
		s.Equal(name, output.Name)
	})
}
