package repository

import (
	"context"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-redis/redismock/v9"
	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
	"github.com/ssentinull/create-apis-using-golang/internal/model/mock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type mockedRepoDependency struct {
	ctx       context.Context
	ctrl      *gomock.Controller
	db        *gorm.DB
	sql       sqlmock.Sqlmock
	redis     *redis.Client
	redisCmd  redismock.ClientMock
	cacheRepo *mock.MockCacheRepository
}

func (d mockedRepoDependency) close() {
	d.ctx.Done()
	d.ctrl.Finish()
	d.redis.Close()
}

func newMockedDependency(t *testing.T) mockedRepoDependency {
	dep := mockedRepoDependency{}

	dep.ctrl = gomock.NewController(t)
	dep.ctx = context.Background()

	mockDBConn, mockSQL, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dep.sql = mockSQL
	dialector := postgres.New(postgres.Config{
		Conn:       mockDBConn,
		DriverName: "postgres",
	})

	mockGormConn, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a gorm connection", err)
	}

	dep.db = mockGormConn
	dep.cacheRepo = mock.NewMockCacheRepository(dep.ctrl)
	dep.redis, dep.redisCmd = redismock.NewClientMock()

	return dep
}
