package sd

type TargetingResolvedExpression struct {
	BaseMode
	AstreId int64 `gorm:"column:astre_id" db:"column:astre_id" json:"astre_id"`
	CampaignId string `gorm:"column:campaign_id" db:"column:campaign_id" json:"campaign_id"`
	AdGroupId string `gorm:"column:ad_group_id" db:"column:ad_group_id" json:"ad_group_id"`
	TargetId string `gorm:"column:target_id" db:"column:target_id" json:"target_id"`
	ResolvedExpressionType string `gorm:"column:resolved_expression_type" db:"column:resolved_expression_type" json:"resolved_expression_type"`
	ResolvedExpressionValue string `gorm:"column:resolved_expression_value" db:"column:resolved_expression_value" json:"resolved_expression_value"`
	ResolvedExpressionIndex int `gorm:"column:resolved_expression_index" db:"column:resolved_expression_index" json:"resolved_expression_index"`
	DataSyncVersion int `gorm:"column:data_sync_version" db:"column:data_sync_version" json:"data_sync_version"`
}

func (tre *TargetingResolvedExpression) TableName() string {
	return "ads_sd_targeting_resolved_expression"
}
