package sd

import (
	_const "data-seeder/internal/const"
	"data-seeder/internal/pkg"
)

type Targeting struct {
	BaseMode
	AstId int64 `gorm:"column:ast_id" db:"column:ast_id" json:"ast_id"`
	CampaignId string `gorm:"column:campaign_id" db:"column:campaign_id" json:"campaign_id"`
	AdGroupId string `gorm:"column:ad_group_id" db:"column:ad_group_id" json:"ad_group_id"`
	TargetId string `gorm:"column:target_id" db:"column:target_id" json:"target_id"`
	Bid float32 `gorm:"column:bid" db:"column:bid" json:"bid"`
	ExpressionType string `gorm:"column:expression_type" db:"column:expression_type" json:"expression_type"` // "manual", "auto"
	State string `gorm:"column:state" db:"column:state" json:"state"` // "enabled", "paused", "archived"
	// "ADVERTISER_STATUS_ENABLED", "STATUS_UNAVAILABLE", "ADVERTISER_PAUSED", "ACCOUNT_OUT_OF_BUDGET", "ADVERTISER_PAYMENT_FAILURE", "CAMPAIGN_PAUSED", "CAMPAIGN_ARCHIVED", "PENDING_START_DATE", "ENDED", "CAMPAIGN_OUT_OF_BUDGET", "AD_GROUP_STATUS_ENABLED", "AD_GROUP_PAUSED", "AD_GROUP_ARCHIVED", "AD_GROUP_INCOMPLETE", "AD_GROUP_LOW_BID", "TARGET_STATUS_LIVE", "TARGET_STATUS_PAUSED", "TARGET_STATUS_ARCHIVED"
	ServingStatus string `gorm:"column:serving_status" db:"column:serving_status" json:"serving_status"`
	CreationDate int64 `gorm:"column:creation_date" db:"column:creation_date" json:"creation_date"`
	LastUpdatedDate int64 `gorm:"column:last_updated_date" db:"column:last_updated_date" json:"last_updated_date"`
	DataSyncVersion int `gorm:"column:data_sync_version" db:"column:data_sync_version" json:"data_sync_version"`
}

func (t *Targeting) TableName() string {
	return "ads_sd_targeting"
}

func GenTargeting(n int, groupMaps []pkg.GroupMap) []Targeting {
	var targetings []Targeting

	length := len(groupMaps)

	for i := 0; i < n; i++ {

		index := pkg.GetIndex(length)

		targeting :=  Targeting{
			CampaignId: groupMaps[index].CampaignId,
			AdGroupId: groupMaps[index].GroupId,
			TargetId: pkg.RandTargetId(),
			Bid: 3.00,
			ExpressionType: "manual", // "manual"
			State: "enabled", //  || "paused" || "archived",
			ServingStatus: "CAMPAIGN_STATUS_ENABLED",
			CreationDate: 0,
			LastUpdatedDate: 0,
			DataSyncVersion: 0,
		}

		targeting.Zid = _const.ZID
		targeting.Sid = _const.SID
		targeting.ProfileId = _const.PROFILE_ID
		targeting.GmtModified = pkg.GenDate()
		targeting.GmtCreate = pkg.GenDate()

		targetings = append(targetings, targeting)
	}

	return targetings
}
