package sd

import (
	"data-seeder/internal/db"
	"data-seeder/internal/entity/sd"
	"data-seeder/internal/pkg"
	"fmt"
	"gorm.io/gorm"
)

type reportProductRepository struct {
	db *gorm.DB
}

func NewReportProductRepository() *reportProductRepository {
	return &reportProductRepository{
		db: db.GetDB(),
	}
}

func (rp *reportProductRepository) Seed(n int, productMaps []pkg.ProductMap)  {
	// 获取数据
	reportProducts := sd.GenReportProduct(n, productMaps)
	fmt.Println(len(reportProducts))

	rp.db.Model(&sd.ReportProduct{}).Create(&reportProducts)

	return
}
