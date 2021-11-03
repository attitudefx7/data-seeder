package sd

import (
	_const "data-seeder/internal/const"
	"data-seeder/internal/pkg"
)

type Product struct {
	BaseMode
	AspId int64 `gorm:"column:asp_id" db:"column:asp_id" json:"asp_id"`
	CampaignId string `gorm:"column:campaign_id" db:"column:campaign_id" json:"campaign_id"`
	AdGroupId string `gorm:"column:ad_group_id" db:"column:ad_group_id" json:"ad_group_id"`
	AdId string `gorm:"column:ad_id" db:"column:ad_id" json:"ad_id"`
	Asin string `gorm:"column:asin" db:"column:asin" json:"asin"`
	Sku string `gorm:"column:sku" db:"column:sku" json:"sku"`
	State string `gorm:"column:state" db:"column:state" json:"state"` // "enabled", "paused", "archived"
	// "ADVERTISER_STATUS_ENABLED", "STATUS_UNAVAILABLE", "ADVERTISER_PAUSED", "ACCOUNT_OUT_OF_BUDGET", "ADVERTISER_PAYMENT_FAILURE", "CAMPAIGN_PAUSED", "CAMPAIGN_ARCHIVED", "PENDING_START_DATE", "ENDED", "CAMPAIGN_OUT_OF_BUDGET", "AD_GROUP_STATUS_ENABLED", "AD_GROUP_PAUSED", "AD_GROUP_ARCHIVED", "AD_GROUP_INCOMPLETE", "AD_GROUP_LOW_BID", "AD_STATUS_LIVE", "AD_STATUS_PAUSED", "AD_STATUS_ARCHIVED", "MISSING_IMAGE", "MISSING_DECORATION", "NOT_BUYABLE", "NOT_IN_BUYBOX", "OUT_OF_STOCK", "NOT_IN_POLICY"
	ServingStatus string `gorm:"column:serving_status" db:"column:serving_status" json:"serving_status"`
	CreationDate int64 `gorm:"column:creation_date" db:"column:creation_date" json:"creation_date"`
	LastUpdatedDate int64 `gorm:"column:last_updated_date" db:"column:last_updated_date" json:"last_updated_date"`
	DataSyncVersion int `gorm:"column:data_sync_version" db:"column:data_sync_version" json:"data_sync_version"`
}

func (p *Product) TableName() string {
	return "ads_sd_product"
}

func GenProduct(n int, groupIdMapCampaignIds []pkg.GroupMap) []Product {
	var products []Product

	length := len(groupIdMapCampaignIds)

	for i := 0; i < n; i++ {

		index := pkg.GetIndex(length)

		product :=  Product{
			CampaignId: groupIdMapCampaignIds[index].CampaignId,
			AdGroupId: groupIdMapCampaignIds[index].GroupId,
			AdId: pkg.RandProductId(),
			Asin: pkg.GenName("asin"),
			Sku: pkg.GenName("sku"), // "T00030"
			State: "enabled", //  || "paused" || "archived",
			ServingStatus: "ADVERTISER_STATUS_ENABLED",
			CreationDate: 0,
			LastUpdatedDate: 0,
			DataSyncVersion: 0,
		}

		product.Zid = _const.ZID
		product.Sid = _const.SID
		product.ProfileId = _const.PROFILE_ID
		product.GmtModified = pkg.GenDate()
		product.GmtCreate = pkg.GenDate()

		products = append(products, product)
	}

	return products
}
