package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	var users []*User

	// working
	if err := DB.Where(`id in (?)`, []string{`1`, `2`}).
		Where(
			`EXISTS (?)`,
			DB.Table(`accounts`).Where(`"accounts".user_id = "users".id`),
		).Find(&users).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// not working
	ids := []string{`1`, `2`}
	accountID := `3`
	db := DB
	db = whereIn(db, ids)
	db = whereExists(db, accountID)
	if err := db.Find(&users).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

func whereIn(db *gorm.DB, ids []string) *gorm.DB {
	return DB.Where(`id in (?)`, ids)
}

func whereExists(db *gorm.DB, accountID string) *gorm.DB {
	return db.Where(
		`EXISTS (?)`,
		db.Table(`accounts`).Where(`"accounts".user_id = "users".id`),
	)
}
