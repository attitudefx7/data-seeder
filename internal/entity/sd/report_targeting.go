package sd

import (
	_const "data-seeder/internal/const"
	"data-seeder/internal/pkg"
)

type ReportTargeting struct {
	BaseMode
	Id int64 `gorm:"column:id" db:"column:id" json:"id"`
	ReportDate string `gorm:"column:report_date" db:"column:report_date" json:"report_date"`
	Currency string `gorm:"column:currency" db:"column:currency" json:"currency"`
	CampaignId string `gorm:"column:campaign_id" db:"column:campaign_id" json:"campaign_id"`
	AdGroupId string `gorm:"column:ad_group_id" db:"column:ad_group_id" json:"ad_group_id"`
	CampaignName string `gorm:"column:campaign_name" db:"column:campaign_name" json:"campaign_name"`
	AdGroupName string `gorm:"column:ad_group_name" db:"column:ad_group_name" json:"ad_group_name"`
	TargetId string `gorm:"column:target_id" db:"column:target_id" json:"target_id"`
	TargetingType string `gorm:"column:targeting_type" db:"column:targeting_type" json:"targeting_type"`
	TargetingText string `gorm:"column:targeting_text" db:"column:targeting_text" json:"targeting_text"`
	TargetingExpression string `gorm:"column:targeting_expression" db:"column:targeting_expression" json:"targeting_expression"`
	BaseReport
	UnitsOrder1d
}

func (rt *ReportTargeting) TableName() string {
	return "ads_sd_report_targeting"
}

func GenReportTargeting(n int, targetMaps []pkg.TargetMap) []ReportTargeting {
	length := len(targetMaps)

	var reportTargetings []ReportTargeting

	for i := 0; i < n; i++ {
		index := pkg.GetIndex(length)
		targetMap := targetMaps[index]

		reportTarget := ReportTargeting{
			ReportDate: pkg.GenDate2(),
			Currency: "USD",
			CampaignId: targetMap.CampaignId,
			AdGroupId: targetMap.GroupId,
			CampaignName: "campaignName",
			AdGroupName: "groupName",
			TargetId: targetMap.TargetId,
			TargetingType: "",
			TargetingText: "d41d8cd98f00b204e9800998ecf8427e",
			TargetingExpression: "d41d8cd98f00b204e9800998ecf8427e",
		}

		reportTarget.Zid = _const.ZID
		reportTarget.Sid = _const.SID
		reportTarget.ProfileId = _const.PROFILE_ID
		reportTarget.GmtModified = pkg.GenDate()
		reportTarget.GmtCreate = pkg.GenDate()

		reportTarget.Impressions = pkg.Rand()
		reportTarget.Clicks = pkg.Rand()
		reportTarget.Cost = 3.14
		reportTarget.AttributedConversions1d = pkg.Rand()
		reportTarget.AttributedConversions7d = pkg.Rand()
		reportTarget.AttributedConversions14d = pkg.Rand()
		reportTarget.AttributedConversions30d = pkg.Rand()
		reportTarget.AttributedConversions1dSameSku = pkg.Rand()
		reportTarget.AttributedConversions7dSameSku = pkg.Rand()
		reportTarget.AttributedConversions14dSameSku = pkg.Rand()
		reportTarget.AttributedConversions30dSameSku = pkg.Rand()
		reportTarget.AttributedUnitsOrdered1d = pkg.Rand()
		reportTarget.AttributedUnitsOrdered7d = pkg.Rand()
		reportTarget.AttributedUnitsOrdered14d = pkg.Rand()
		reportTarget.AttributedUnitsOrdered30d = pkg.Rand()
		reportTarget.AttributedSales1d = pkg.RandFloat32()
		reportTarget.AttributedSales7d = pkg.RandFloat32()
		reportTarget.AttributedSales14d = pkg.RandFloat32()
		reportTarget.AttributedSales30d = pkg.RandFloat32()
		reportTarget.AttributedSales1dSameSku = pkg.RandFloat32()
		reportTarget.AttributedSales7dSameSku = pkg.RandFloat32()
		reportTarget.AttributedSales14dSameSku = pkg.RandFloat32()
		reportTarget.AttributedSales30dSameSku = pkg.RandFloat32()

		reportTargetings = append(reportTargetings, reportTarget)
	}

	return reportTargetings
}
