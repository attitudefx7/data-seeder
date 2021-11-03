package sd

import (
	"data-seeder/internal/db"
	"data-seeder/internal/entity/sd"
	"data-seeder/internal/pkg"
	"gorm.io/gorm"
)

type groupRepository struct {
	db *gorm.DB
}


func NewGroupRepository() *groupRepository {
	return &groupRepository{
		db: db.GetDB(),
	}
}

func (g *groupRepository) Seed(n int, campaignMaps []pkg.CampaignMap) []pkg.GroupMap {
	groups := sd.GenGroup(n, campaignMaps)

	db.GetDB().Model(sd.Group{}).Create(&groups)

	var groupMaps []pkg.GroupMap

	campaignIdMapName := make(map[string]string)

	for _, campaignMap := range campaignMaps {
		campaignIdMapName[campaignMap.CampaignId] = campaignMap.CampaignName
	}

	for _, group := range groups{
		newMap := pkg.GroupMap{
			CampaignId: group.CampaignId,
			CampaignName: campaignIdMapName[group.CampaignId],
			GroupId: group.AdGroupId,
			GroupName: group.Name,
		}

		groupMaps = append(groupMaps, newMap)
	}

	return groupMaps
}
