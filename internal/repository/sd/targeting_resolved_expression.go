package sd

import (
	"data-seeder/internal/db"
	"gorm.io/gorm"
)

type targetingResolvedExpressionRepository struct {
	db *gorm.DB
}

func NewTargetingResolvedExpressionRepository() *targetingResolvedExpressionRepository {
	return &targetingResolvedExpressionRepository{
		db: db.GetDB(),
	}
}
