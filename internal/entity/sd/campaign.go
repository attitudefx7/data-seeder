package sd

import (
	_const "data-seeder/internal/const"
	"data-seeder/internal/pkg"
)

type Campaign struct {
	BaseMode
	Id int64 `gorm:"column:asc_id" db:"column:asc_id" json:"asc_id"`
	//Zid int64 `gorm:"column:zid" db:"column:zid" json:"zid"`
	//Sid int64 `gorm:"column:sid" db:"column:sid" json:"sid"`
	//ProfileId string `gorm:"column:profile_id" db:"column:profile_id" json:"profile_id"`
	//GmtModified string `gorm:"column:gmt_modified" db:"column:gmt_modified" json:"gmt_modified"`
	//GmtCreate string `gorm:"column:gmt_create" db:"column:gmt_create" json:"gmt_create"`
	CampaignId string `gorm:"column:campaign_id" db:"column:campaign_id" json:"campaign_id"`
	Name string `gorm:"column:name" db:"column:name" json:"name"`
	Tactic string `gorm:"column:tactic" db:"column:tactic" json:"tactic"`
	StartDate string `gorm:"column:start_date" db:"column:start_date" json:"start_date"` // YYYYMMDD
	EndDate string `gorm:"column:end_date" db:"column:end_date" json:"end_date"` // YYYYMMDD
	State string `gorm:"column:state" db:"column:state" json:"state"` // "enabled" || "paused" || "archived"
	Budget float32 `gorm:"column:budget" db:"column:budget" json:"budget"`
	BudgetType string `gorm:"column:budget_type" db:"column:budget_type" json:"budget_type"` // "daily"
	// "ADVERTISER_STATUS_ENABLED","STATUS_UNAVAILABLE","ADVERTISER_PAUSED","ACCOUNT_OUT_OF_BUDGET","ADVERTISER_PAYMENT_FAILURE","CAMPAIGN_PAUSED","CAMPAIGN_ARCHIVED","PENDING_START_DATE","ENDED","CAMPAIGN_OUT_OF_BUDGET"
	ServingStatus string `gorm:"column:serving_status" db:"column:serving_status" json:"serving_status"`
	CreationDate int64 `gorm:"column:creation_date" db:"column:creation_date" json:"creation_date"` // 创建时间戳，毫秒
	LastUpdatedDate int64 `gorm:"column:last_updated_date" db:"column:last_updated_date" json:"last_updated_date"` // 最后更新时间戳，毫秒
	CostType string `gorm:"column:cost_type" db:"column:cost_type" json:"cost_type"`
	DeliveryProfile string `gorm:"column:delivery_profile" db:"column:delivery_profile" json:"delivery_profile"`
	DataSyncVersion int `gorm:"column:data_sync_version" db:"column:data_sync_version" json:"data_sync_version"`
}

func (c *Campaign) TableName() string {
	return "ads_sd_campaign"
}

func GenCampaign(n int) []Campaign {
	var campaigns []Campaign

	for i := 0; i < n; i++ {
		campaign :=  Campaign{
			CampaignId: pkg.RandCampaignId(),
			Name: pkg.GenName("campaign"),
			Tactic: "T00020", // "T00030"
			StartDate: "20211101",
			EndDate: "", // nullable
			State: "enabled", //  || "paused" || "archived",
			Budget: 3.00,
			BudgetType: "daily",
			// "ADVERTISER_STATUS_ENABLED","STATUS_UNAVAILABLE","ADVERTISER_PAUSED","ACCOUNT_OUT_OF_BUDGET","ADVERTISER_PAYMENT_FAILURE","CAMPAIGN_PAUSED","CAMPAIGN_ARCHIVED","PENDING_START_DATE","ENDED","CAMPAIGN_OUT_OF_BUDGET"
			ServingStatus: "ADVERTISER_STATUS_ENABLED",
			CreationDate: 0,
			LastUpdatedDate: 0,
			CostType: "cpc",
			DeliveryProfile: "as_soon_as_possible",
			DataSyncVersion: 0,
			//Zid: 1,
			//Sid: 16,
			//ProfileId: "121923590660074",
			//GmtModified: pkg.GenDate(),
			//GmtCreate: pkg.GenDate(),
		}

		campaign.Zid = _const.ZID
		campaign.Sid = _const.SID
		campaign.ProfileId = _const.PROFILE_ID
		campaign.GmtModified = pkg.GenDate()
		campaign.GmtCreate = pkg.GenDate()

		campaigns = append(campaigns, campaign)
	}

	return campaigns
}