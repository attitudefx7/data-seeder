package sd

import (
	"data-seeder/internal/db"
	"data-seeder/internal/entity/sd"
	"data-seeder/internal/pkg"
	"fmt"
	"gorm.io/gorm"
)

type reportTargetingRepository struct {
	db *gorm.DB
}

func NewReportTargetingRepository() *reportTargetingRepository {
	return &reportTargetingRepository{
		db: db.GetDB(),
	}
}

func (rt *reportTargetingRepository) Seed(n int, targetMaps []pkg.TargetMap) bool {
	// 获取数据
	reportTargets := sd.GenReportTargeting(n, targetMaps)
	fmt.Println(len(reportTargets))

	rt.db.Model(&sd.ReportTargeting{}).Create(&reportTargets)

	return true
}
