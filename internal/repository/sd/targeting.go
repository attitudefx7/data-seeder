package sd

import (
	"data-seeder/internal/db"
	"data-seeder/internal/entity/sd"
	"data-seeder/internal/pkg"
	"gorm.io/gorm"
)

type targetingRepository struct {
	db *gorm.DB
}


func NewTargetingRepository() *targetingRepository {
	return &targetingRepository{
		db: db.GetDB(),
	}
}

func (tp *targetingRepository) Seed(n int, groupMaps []pkg.GroupMap) []pkg.TargetMap {
	targetings := sd.GenTargeting(n, groupMaps)

	db.GetDB().Model(sd.Targeting{}).Create(&targetings)

	var targetingMaps []pkg.TargetMap

	for _, target := range targetings{
		newMap := pkg.TargetMap{
			CampaignId: target.CampaignId,
			GroupId: target.AdGroupId,
			TargetId: target.TargetId,
		}

		targetingMaps = append(targetingMaps, newMap)
	}

	return targetingMaps
}
