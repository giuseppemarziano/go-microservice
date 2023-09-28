package container

import (
	"fmt"
	"github.com/palantir/stacktrace"
	"go-microservice/domain/entities"
	"go-microservice/domain/valueobject"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var EntityTypes = []interface{}{
	&entities.User{},
	&valueobject.Post{},
}

func NewGormDBConnection(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, stacktrace.Propagate(
			err,
			"error on opening database with GORM")
	}

	for _, entityType := range EntityTypes {
		if err := db.AutoMigrate(entityType); err != nil {
			return nil, stacktrace.Propagate(
				err,
				fmt.Sprintf("error on auto-migrating table for %T", entityType),
			)
		}
	}

	return db, nil
}
