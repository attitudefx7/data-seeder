package sd

import (
	"data-seeder/internal/repository/sd"
	"fmt"
	"net/http"
	"time"
)

type Invoke struct {

}

func (i *Invoke) Invoke(w http.ResponseWriter, r *http.Request)  {
	start := time.Now()
	fmt.Fprintln(w, "seeder start")

	fmt.Fprintln(w, "开始生成 sd campaign 数据")
	campaignRepository := sd.NewCampaignRepository()

	// 返回创建的广告活动ID
	campaignMaps := campaignRepository.Seed(10)
	fmt.Println(len(campaignMaps))

	fmt.Fprintln(w, "开始生成 sd group 数据")
	groupRepository := sd.NewGroupRepository()

	// 返回创建的广告组ID => 广告活动ID
	groupMaps := groupRepository.Seed(20, campaignMaps)
	fmt.Println(len(groupMaps))

	fmt.Fprintln(w, "开始生成 sd product 数据")
	productRepository := sd.NewProductRepository()

	// 返回创建的广告ID
	productMaps := productRepository.Seed(50, groupMaps)
	fmt.Println(len(productMaps))

	fmt.Fprintln(w, "开始生成 sd targeting 数据")
	targetingRepository := sd.NewTargetingRepository()

	// 返回创建的广告ID
	targetingMaps := targetingRepository.Seed(50, groupMaps)
	fmt.Println(len(targetingMaps))

	//fmt.Fprintln(w, "开始生成 sd targeting_expression 数据")
	//expressionRepository := sd.NewTargetingExpressionRepository()
	//
	//fmt.Fprintln(w, "开始生成 sd targeting_resolved_expression 数据")
	//resolvedRepository := sd.NewTargetingResolvedExpressionRepository()
	//
	fmt.Fprintln(w, "开始生成 sd report campaign 数据")
	reportCampaignRe := sd.NewReportCampaignRepository()

	reportCampaignRe.Seed(100, campaignMaps)

	fmt.Fprintln(w, "开始生成 sd report group 数据")
	reportGroupRe := sd.NewReportGroupRepository()

	reportGroupRe.Seed(100, groupMaps)

	fmt.Fprintln(w, "开始生成 sd report product 数据")
	reportProductRe := sd.NewReportProductRepository()

	reportProductRe.Seed(100, productMaps)

	fmt.Fprintln(w, "开始生成 sd report targeting 数据")
	reportTargetingRe := sd.NewReportTargetingRepository()

	reportTargetingRe.Seed(100, targetingMaps)

	fmt.Fprintln(w, "seeder end")
	end := time.Now()

	fmt.Print("总耗时：")
	fmt.Println(end.Sub(start))
}