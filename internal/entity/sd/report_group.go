package sd

import (
	_const "data-seeder/internal/const"
	"data-seeder/internal/pkg"
)

type ReportGroup struct {
	BaseMode
	Id int64 `gorm:"column:id" db:"column:id" json:"id"`
	ReportDate string `gorm:"column:report_date" db:"column:report_date" json:"report_date"`
	CampaignId string `gorm:"column:campaign_id" db:"column:campaign_id" json:"campaign_id"`
	AdGroupId string `gorm:"column:ad_group_id" db:"column:ad_group_id" json:"ad_group_id"`
	CampaignName string `gorm:"column:campaign_name" db:"column:campaign_name" json:"campaign_name"`
	AdGroupName string `gorm:"column:ad_group_name" db:"column:ad_group_name" json:"ad_group_name"`
	BaseReport
	UnitsOrder1d
}

func (rg *ReportGroup) TableName() string {
	return "ads_sd_report_group"
}

func GenReportGroup(n int, groupMaps []pkg.GroupMap) []ReportGroup {
	length := len(groupMaps)

	var reportGroups []ReportGroup

	for i := 0; i < n; i++ {
		index := pkg.GetIndex(length)
		groupMap := groupMaps[index]

		reportGroup := ReportGroup{
			ReportDate: pkg.GenDate2(),
			CampaignId: groupMap.CampaignId,
			CampaignName: groupMap.CampaignName,
			AdGroupId: groupMap.GroupId,
			AdGroupName: groupMap.GroupName,
		}

		reportGroup.Zid = _const.ZID
		reportGroup.Sid = _const.SID
		reportGroup.ProfileId = _const.PROFILE_ID
		reportGroup.GmtModified = pkg.GenDate()
		reportGroup.GmtCreate = pkg.GenDate()

		reportGroup.Impressions = pkg.Rand()
		reportGroup.Clicks = pkg.Rand()
		reportGroup.Cost = 3.14
		reportGroup.AttributedConversions1d = pkg.Rand()
		reportGroup.AttributedConversions7d = pkg.Rand()
		reportGroup.AttributedConversions14d = pkg.Rand()
		reportGroup.AttributedConversions30d = pkg.Rand()
		reportGroup.AttributedConversions1dSameSku = pkg.Rand()
		reportGroup.AttributedConversions7dSameSku = pkg.Rand()
		reportGroup.AttributedConversions14dSameSku = pkg.Rand()
		reportGroup.AttributedConversions30dSameSku = pkg.Rand()
		reportGroup.AttributedUnitsOrdered1d = pkg.Rand()
		reportGroup.AttributedUnitsOrdered7d = pkg.Rand()
		reportGroup.AttributedUnitsOrdered14d = pkg.Rand()
		reportGroup.AttributedUnitsOrdered30d = pkg.Rand()
		reportGroup.AttributedSales1d = pkg.RandFloat32()
		reportGroup.AttributedSales7d = pkg.RandFloat32()
		reportGroup.AttributedSales14d = pkg.RandFloat32()
		reportGroup.AttributedSales30d = pkg.RandFloat32()
		reportGroup.AttributedSales1dSameSku = pkg.RandFloat32()
		reportGroup.AttributedSales7dSameSku = pkg.RandFloat32()
		reportGroup.AttributedSales14dSameSku = pkg.RandFloat32()
		reportGroup.AttributedSales30dSameSku = pkg.RandFloat32()

		reportGroups = append(reportGroups, reportGroup)
	}

	return reportGroups
}