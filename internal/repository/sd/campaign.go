package sd

import (
	"data-seeder/internal/db"
	"data-seeder/internal/entity/sd"
	"data-seeder/internal/pkg"
	"gorm.io/gorm"
)

type campaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository() *campaignRepository {
	return &campaignRepository{
		db: db.GetDB(),
	}
}

func (c *campaignRepository) Seed(n int) []pkg.CampaignMap {
	campaigns := sd.GenCampaign(n)

	db.GetDB().Model(sd.Campaign{}).Create(&campaigns)

	var campaignMaps []pkg.CampaignMap

	for _, campaign := range campaigns{
		campaignMap := pkg.CampaignMap{
			CampaignId: campaign.CampaignId,
			CampaignName: campaign.Name,
			CampaignStatus: campaign.ServingStatus,
		}
		campaignMaps = append(campaignMaps, campaignMap)
	}

	return campaignMaps
}

