package database

import (
	"homework/domain"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	var err error

	if err = dropTables(db); err != nil {
		return err
	}

	if err = setupJoinTables(db); err != nil {
		return err
	}

	if err = autoMigrates(db); err != nil {
		return err
	}

	return createViews(db)
}

func autoMigrates(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.User{},
		&domain.PasswordResetToken{},
	)
}

func dropTables(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&domain.PasswordResetToken{},
		&domain.User{},
	)
}

func setupJoinTables(db *gorm.DB) error {
	var err error

	return err
}

func createViews(db *gorm.DB) error {
	var err error

	return err
}
