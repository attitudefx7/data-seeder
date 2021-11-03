package sd

import (
	"data-seeder/internal/db"
	"data-seeder/internal/entity/sd"
	"data-seeder/internal/pkg"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}


func NewProductRepository() *productRepository {
	return &productRepository{
		db: db.GetDB(),
	}
}

func (pp *productRepository) Seed(n int, groupIdMapCampaignIds []pkg.GroupMap) []pkg.ProductMap {
	products := sd.GenProduct(n, groupIdMapCampaignIds)

	db.GetDB().Model(sd.Product{}).Create(&products)

	var productMaps []pkg.ProductMap

	for _, product := range products{
		newMap := pkg.ProductMap{
			CampaignId: product.CampaignId,
			GroupId: product.AdGroupId,
			ProductId: product.AdId,
			Sku: product.Sku,
			Asin: product.Asin,
		}

		productMaps = append(productMaps, newMap)
	}

	return productMaps
}
