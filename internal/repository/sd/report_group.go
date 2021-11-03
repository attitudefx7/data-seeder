package sd

import (
	"data-seeder/internal/db"
	"data-seeder/internal/entity/sd"
	"data-seeder/internal/pkg"
	"fmt"
	"gorm.io/gorm"
)

type reportGroupRepository struct {
	db *gorm.DB
}

func NewReportGroupRepository() *reportGroupRepository {
	return &reportGroupRepository{
		db: db.GetDB(),
	}
}

func (rg *reportGroupRepository) Seed(n int, groupMaps []pkg.GroupMap)  {
	// 获取数据
	reportGroups := sd.GenReportGroup(n, groupMaps)
	fmt.Println(len(reportGroups))

	rg.db.Model(&sd.ReportGroup{}).Create(&reportGroups)

	return
}
