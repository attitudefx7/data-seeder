package sd

import (
	_const "data-seeder/internal/const"
	"data-seeder/internal/pkg"
)

type ReportProduct struct {
	BaseMode
	Id int64 `gorm:"column:id" db:"column:id" json:"id"`
	ReportDate string `gorm:"column:report_date" db:"column:report_date" json:"report_date"`
	CampaignId string `gorm:"column:campaign_id" db:"column:campaign_id" json:"campaign_id"`
	AdGroupId string `gorm:"column:ad_group_id" db:"column:ad_group_id" json:"ad_group_id"`
	AdId string `gorm:"column:ad_id" db:"column:ad_id" json:"ad_id"`
	Sku string `gorm:"column:sku" db:"column:sku" json:"sku"`
	Asin string `gorm:"column:asin" db:"column:asin" json:"asin"`
	BaseReport
}

func (rp *ReportProduct) TableName() string {
	return "ads_sd_report_product"
}

func GenReportProduct(n int, productMaps []pkg.ProductMap) []ReportProduct {
	length := len(productMaps)

	var reportProducts []ReportProduct

	for i := 0; i < n; i++ {
		index := pkg.GetIndex(length)
		productMap := productMaps[index]

		reportProduct := ReportProduct{
			ReportDate: pkg.GenDate2(),
			CampaignId: productMap.CampaignId,
			AdGroupId: productMap.GroupId,
			AdId: productMap.ProductId,
			Sku: productMap.Sku,
			Asin: productMap.Asin,
		}

		reportProduct.Zid = _const.ZID
		reportProduct.Sid = _const.SID
		reportProduct.ProfileId = _const.PROFILE_ID
		reportProduct.GmtModified = pkg.GenDate()
		reportProduct.GmtCreate = pkg.GenDate()

		reportProduct.Impressions = pkg.Rand()
		reportProduct.Clicks = pkg.Rand()
		reportProduct.Cost = 3.14
		reportProduct.AttributedConversions1d = pkg.Rand()
		reportProduct.AttributedConversions7d = pkg.Rand()
		reportProduct.AttributedConversions14d = pkg.Rand()
		reportProduct.AttributedConversions30d = pkg.Rand()
		reportProduct.AttributedConversions1dSameSku = pkg.Rand()
		reportProduct.AttributedConversions7dSameSku = pkg.Rand()
		reportProduct.AttributedConversions14dSameSku = pkg.Rand()
		reportProduct.AttributedConversions30dSameSku = pkg.Rand()
		reportProduct.AttributedUnitsOrdered7d = pkg.Rand()
		reportProduct.AttributedUnitsOrdered14d = pkg.Rand()
		reportProduct.AttributedUnitsOrdered30d = pkg.Rand()
		reportProduct.AttributedSales1d = pkg.RandFloat32()
		reportProduct.AttributedSales7d = pkg.RandFloat32()
		reportProduct.AttributedSales14d = pkg.RandFloat32()
		reportProduct.AttributedSales30d = pkg.RandFloat32()
		reportProduct.AttributedSales1dSameSku = pkg.RandFloat32()
		reportProduct.AttributedSales7dSameSku = pkg.RandFloat32()
		reportProduct.AttributedSales14dSameSku = pkg.RandFloat32()
		reportProduct.AttributedSales30dSameSku = pkg.RandFloat32()

		reportProducts = append(reportProducts, reportProduct)
	}

	return reportProducts
}
