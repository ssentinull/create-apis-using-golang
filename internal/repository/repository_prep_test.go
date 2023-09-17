package repository

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/ssentinull/create-apis-using-golang/internal/model/mock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type mockedRepoDependency struct {
	db        *gorm.DB
	sql       sqlmock.Sqlmock
	cacheRepo *mock.MockCacheRepository
}

func newMockedDependency(ctrl *gomock.Controller) mockedRepoDependency {
	dep := mockedRepoDependency{}
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
	dep.cacheRepo = mock.NewMockCacheRepository(ctrl)

	return dep
}
