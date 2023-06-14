package repository_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/celsopires1999/matchreport/configs"
	"github.com/celsopires1999/matchreport/internal/entity"
	"github.com/celsopires1999/matchreport/internal/infra/repository"
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

	db, err := sql.Open("postgres", configs.DBConn)

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

func (s *TenantDBTestSuite) TestCreateTenant() {
	s.Run("should create tenant", func() {
		name := "Veterans' Football League"
		tenant, _ := entity.NewTenant(name)
		tenantRepo := repository.NewTenantRepositoryDb(s.db)
		ctx := context.Background()
		err := tenantRepo.Create(ctx, tenant)
		s.Nil(err)
	})
}
