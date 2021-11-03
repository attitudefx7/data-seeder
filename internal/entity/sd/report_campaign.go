package sd

import (
	_const "data-seeder/internal/const"
	"data-seeder/internal/pkg"
)

type ReportCampaign struct {
	BaseMode
	Id int64 `gorm:"column:id" db:"column:id" json:"id"`
	ReportDate string `gorm:"column:report_date" db:"column:report_date" json:"report_date"`
	CampaignId string `gorm:"column:campaign_id" db:"column:campaign_id" json:"campaign_id"`
	CampaignName string `gorm:"column:campaign_name" db:"column:campaign_name" json:"campaign_name"`
	CampaignStatus string `gorm:"column:campaign_status" db:"column:campaign_status" json:"campaign_status"`
	BaseReport
	UnitsOrder1d
}

func (rc *ReportCampaign) TableName() string {
	return "ads_sd_report_campaign"
}

func GenReportCampaign(n int, campaignMaps []pkg.CampaignMap) []ReportCampaign {
	length := len(campaignMaps)

	var reportCampaigns []ReportCampaign

	for i := 0; i < n; i++ {
		index := pkg.GetIndex(length)
		campaignMap := campaignMaps[index]

		reportCampaign := ReportCampaign{
			ReportDate: pkg.GenDate2(),
			CampaignId: campaignMap.CampaignId,
			CampaignName: campaignMap.CampaignName,
			CampaignStatus: campaignMap.CampaignStatus,
		}

		reportCampaign.Zid = _const.ZID
		reportCampaign.Sid = _const.SID
		reportCampaign.ProfileId = _const.PROFILE_ID
		reportCampaign.GmtModified = pkg.GenDate()
		reportCampaign.GmtCreate = pkg.GenDate()

		reportCampaign.Impressions = pkg.Rand()
		reportCampaign.Clicks = pkg.Rand()
		reportCampaign.Cost = 3.14
		reportCampaign.AttributedConversions1d = pkg.Rand()
		reportCampaign.AttributedConversions7d = pkg.Rand()
		reportCampaign.AttributedConversions14d = pkg.Rand()
		reportCampaign.AttributedConversions30d = pkg.Rand()
		reportCampaign.AttributedConversions1dSameSku = pkg.Rand()
		reportCampaign.AttributedConversions7dSameSku = pkg.Rand()
		reportCampaign.AttributedConversions14dSameSku = pkg.Rand()
		reportCampaign.AttributedConversions30dSameSku = pkg.Rand()
		reportCampaign.AttributedUnitsOrdered1d = pkg.Rand()
		reportCampaign.AttributedUnitsOrdered7d = pkg.Rand()
		reportCampaign.AttributedUnitsOrdered14d = pkg.Rand()
		reportCampaign.AttributedUnitsOrdered30d = pkg.Rand()
		reportCampaign.AttributedSales1d = pkg.RandFloat32()
		reportCampaign.AttributedSales7d = pkg.RandFloat32()
		reportCampaign.AttributedSales14d = pkg.RandFloat32()
		reportCampaign.AttributedSales30d = pkg.RandFloat32()
		reportCampaign.AttributedSales1dSameSku = pkg.RandFloat32()
		reportCampaign.AttributedSales7dSameSku = pkg.RandFloat32()
		reportCampaign.AttributedSales14dSameSku = pkg.RandFloat32()
		reportCampaign.AttributedSales30dSameSku = pkg.RandFloat32()

		reportCampaigns = append(reportCampaigns, reportCampaign)
	}

	return reportCampaigns
}
