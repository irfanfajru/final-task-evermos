package mysql

import (
	"fmt"
	"tugas_akhir_example/internal/daos"
	"tugas_akhir_example/internal/helper"

	"gorm.io/gorm"
)

func RunMigration(mysqlDB *gorm.DB) {
	err := mysqlDB.AutoMigrate(
		&daos.Book{},
		&daos.User{},
	)

	var count int64
	if mysqlDB.Migrator().HasTable(&daos.Book{}) {
		mysqlDB.Model(&daos.Book{}).Count(&count)
		if count < 1 {
			mysqlDB.CreateInBatches(booksSeed, len(booksSeed))
		}
	}
	// seeder user
	if mysqlDB.Migrator().HasTable(&daos.User{}) {
		mysqlDB.Model(&daos.User{}).Count(&count)
		if count < 1 {
			mysqlDB.CreateInBatches(usersSeed, len(usersSeed))
		}
	}

	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Failed Database Migrated : %s", err.Error()))
	}

	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Database Migrated")
}
