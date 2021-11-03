package sd

type TargetingExpression struct {
	BaseMode
	AsteId int64 `gorm:"column:aste_id" db:"column:aste_id" json:"aste_id"`
	CampaignId string `gorm:"column:campaign_id" db:"column:campaign_id" json:"campaign_id"`
	AdGroupId string `gorm:"column:ad_group_id" db:"column:ad_group_id" json:"ad_group_id"`
	TargetId string `gorm:"column:target_id" db:"column:target_id" json:"target_id"`
	ExpressionType string `gorm:"column:expression_type" db:"column:expression_type" json:"expression_type"`
	ExpressionValue string `gorm:"column:expression_value" db:"column:expression_value" json:"expression_value"`
	ExpressionIndex int `gorm:"column:expression_index" db:"column:expression_index" json:"expression_index"`
	DataSyncVersion int `gorm:"column:data_sync_version" db:"column:data_sync_version" json:"data_sync_version"`
}

func (te *TargetingExpression) TableName() string {
	return "ads_sd_targeting_expression"
}
