package container

import (
	"fmt"
	"github.com/palantir/stacktrace"
	"go-microservice/domain/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// EntityTypes holds a list of all entity types that need to be auto-migrated by GORM.
var EntityTypes = []interface{}{
	&entities.User{},
	&entities.Post{},
}

// NewGormDBConnection initializes a new GORM database connection using individual components and auto-migrates the defined entity types.
func NewGormDBConnection(user, password, host, port, dbName, params string) (*gorm.DB, error) {
	// construct DSN from individual components
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s", user, password, host, port, dbName, params)

	// open a new database connection using GORM with the constructed DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, stacktrace.Propagate(
			err,
			"error on opening database with GORM")
	}

	// auto-migrate tables for each entity type defined in EntityTypes
	for _, entityType := range EntityTypes {
		if err := db.AutoMigrate(entityType); err != nil {
			return nil, stacktrace.Propagate(
				err,
				fmt.Sprintf("error on auto-migrating table for %T", entityType),
			)
		}
	}

	// return the initialized database connection
	return db, nil
}
