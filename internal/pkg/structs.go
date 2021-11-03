package pkg


type CampaignMap struct {
	CampaignId string
	CampaignName string
	CampaignStatus string
}


type GroupMap struct {
	CampaignId string
	CampaignName string
	GroupId string
	GroupName string
}


type ProductMap struct {
	CampaignId string
	GroupId string
	ProductId string
	Sku string
	Asin string
}


type TargetMap struct {
	CampaignId string
	GroupId string
	TargetId string
}