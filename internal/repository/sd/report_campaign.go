package sd

import (
	"data-seeder/internal/db"
	"data-seeder/internal/entity/sd"
	"data-seeder/internal/pkg"
	"fmt"
	"gorm.io/gorm"
)

type reportCampaignRepository struct {
	db *gorm.DB
}

func NewReportCampaignRepository() *reportCampaignRepository {
	return &reportCampaignRepository{
		db: db.GetDB(),
	}
}

func (rc *reportCampaignRepository) Seed(n int, campaignMaps []pkg.CampaignMap)  {
	// 获取数据
	reportCampaigns := sd.GenReportCampaign(n, campaignMaps)
	fmt.Println(len(reportCampaigns))

	rc.db.Model(&sd.ReportCampaign{}).Create(&reportCampaigns)

	return
}
