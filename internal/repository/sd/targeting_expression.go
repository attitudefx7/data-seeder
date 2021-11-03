package sd

import (
	"data-seeder/internal/db"
	"gorm.io/gorm"
)

type targetingExpressionRepository struct {
	db *gorm.DB
}

func NewTargetingExpressionRepository() *targetingExpressionRepository {
	return &targetingExpressionRepository{
		db: db.GetDB(),
	}
}
