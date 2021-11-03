package sd

import (
	_const "data-seeder/internal/const"
	"data-seeder/internal/pkg"
)

type Group struct {
	BaseMode
	AsgId int64 `gorm:"column:asg_id" db:"column:asg_id" json:"asg_id"`
	CampaignId string `gorm:"column:campaign_id" db:"column:campaign_id" json:"campaign_id"`
	AdGroupId string `gorm:"column:ad_group_id" db:"column:ad_group_id" json:"ad_group_id"`
	Tactic string `gorm:"column:tactic" db:"column:tactic" json:"tactic"`
	Name string `gorm:"column:name" db:"column:name" json:"name"`
	DefaultBid float32 `gorm:"column:default_bid" db:"column:default_bid" json:"default_bid"`
	State string `gorm:"column:state" db:"column:state" json:"state"` // "enabled", "paused", "archived"
	ServingStatus string `gorm:"column:serving_status" db:"column:serving_status" json:"serving_status"`
	// ["ADVERTISER_STATUS_ENABLED", "STATUS_UNAVAILABLE", "ADVERTISER_PAUSED", "ACCOUNT_OUT_OF_BUDGET", "ADVERTISER_PAYMENT_FAILURE", "CAMPAIGN_PAUSED", "CAMPAIGN_ARCHIVED", "PENDING_START_DATE", "ENDED", "CAMPAIGN_OUT_OF_BUDGET", "AD_GROUP_STATUS_ENABLED", "AD_GROUP_PAUSED", "AD_GROUP_ARCHIVED", "AD_GROUP_INCOMPLETE", "AD_GROUP_LOW_BID", "ADGROUP_POLICING_PENDING_REVIEW", "ADGROUP_POLICING_CREATIVE_REJECTED "]',
	CreationDate int64 `gorm:"column:creation_date" db:"column:creation_date" json:"creation_date"` // 广告组的创建时间戳，毫秒
	LastUpdatedDate int64 `gorm:"column:last_updated_date" db:"column:last_updated_date" json:"last_updated_date"` // 广告组的创建时间戳，毫秒
	DataSyncVersion int `gorm:"column:data_sync_version" db:"column:data_sync_version" json:"data_sync_version"`
}

func (g *Group) TableName() string {
	return "ads_sd_group"
}

func GenGroup(n int, campaignMaps []pkg.CampaignMap) []Group {
	var groups []Group

	length := len(campaignMaps)

	for i := 0; i < n; i++ {

		index := pkg.GetIndex(length)

		group :=  Group{
			CampaignId: campaignMaps[index].CampaignId,
			Name: pkg.GenName("group"),
			Tactic: "T00020", // "T00030"
			AdGroupId: pkg.RandGroupId(),
			DefaultBid: 3.00, // nullable
			State: "enabled", //  || "paused" || "archived",
			ServingStatus: "ADVERTISER_STATUS_ENABLED",
			CreationDate: 0,
			LastUpdatedDate: 0,
			DataSyncVersion: 0,
		}

		group.Zid = _const.ZID
		group.Sid = _const.SID
		group.ProfileId = _const.PROFILE_ID
		group.GmtModified = pkg.GenDate()
		group.GmtCreate = pkg.GenDate()

		groups = append(groups, group)
	}

	return groups
}