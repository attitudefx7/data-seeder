package sd

type BaseReport struct {
	Impressions int `gorm:"column:impressions" db:"column:impressions" json:"impressions"`
	Clicks int `gorm:"column:clicks" db:"column:clicks" json:"clicks"`
	Cost float32 `gorm:"column:cost" db:"column:cost" json:"cost"`
	AttributedConversions1d int `gorm:"column:attributed_conversions1d" db:"column:attributed_conversions1d" json:"attributed_conversions1d"`
	AttributedConversions7d int `gorm:"column:attributed_conversions7d" db:"column:attributed_conversions7d" json:"attributed_conversions7d"`
	AttributedConversions14d int `gorm:"column:attributed_conversions14d" db:"column:attributed_conversions14d" json:"attributed_conversions14d"`
	AttributedConversions30d int `gorm:"column:attributed_conversions30d" db:"column:attributed_conversions30d" json:"attributed_conversions30d"`
	AttributedConversions1dSameSku int `gorm:"column:attributed_conversions1d_same_sku" db:"column:attributed_conversions1d_same_sku" json:"attributed_conversions1d_same_sku"`
	AttributedConversions7dSameSku int `gorm:"column:attributed_conversions7d_same_sku" db:"column:attributed_conversions7d_same_sku" json:"attributed_conversions7d_same_sku"`
	AttributedConversions14dSameSku int `gorm:"column:attributed_conversions14d_same_sku" db:"column:attributed_conversions14d_same_sku" json:"attributed_conversions14d_same_sku"`
	AttributedConversions30dSameSku int `gorm:"column:attributed_conversions30d_same_sku" db:"column:attributed_conversions30d_same_sku" json:"attributed_conversions30d_same_sku"`
	AttributedUnitsOrdered7d int `gorm:"column:attributed_units_ordered7d" db:"column:attributed_units_ordered7d" json:"attributed_units_ordered7d"`
	AttributedUnitsOrdered14d int `gorm:"column:attributed_units_ordered14d" db:"column:attributed_units_ordered14d" json:"attributed_units_ordered14d"`
	AttributedUnitsOrdered30d int `gorm:"column:attributed_units_ordered30d" db:"column:attributed_units_ordered30d" json:"attributed_units_ordered30d"`
	AttributedSales1d int `gorm:"column:attributed_sales1d" db:"column:attributed_sales1d" json:"attributed_sales1d"`
	AttributedSales7d int `gorm:"column:attributed_sales7d" db:"column:attributed_sales7d" json:"attributed_sales7d"`
	AttributedSales14d int `gorm:"column:attributed_sales14d" db:"column:attributed_sales14d" json:"attributed_sales14d"`
	AttributedSales30d int `gorm:"column:attributed_sales30d" db:"column:attributed_sales30d" json:"attributed_sales30d"`
	AttributedSales1dSameSku int `gorm:"column:attributed_sales1d_same_sku" db:"column:attributed_sales1d_same_sku" json:"attributed_sales1d_same_sku"`
	AttributedSales7dSameSku int `gorm:"column:attributed_sales7d_same_sku" db:"column:attributed_sales7d_same_sku" json:"attributed_sales7d_same_sku"`
	AttributedSales14dSameSku int `gorm:"column:attributed_sales14d_same_sku" db:"column:attributed_sales14d_same_sku" json:"attributed_sales14d_same_sku"`
	AttributedSales30dSameSku int `gorm:"column:attributed_sales30d_same_sku" db:"column:attributed_sales30d_same_sku" json:"attributed_sales30d_same_sku"`
}

type UnitsOrder1d struct {
	AttributedUnitsOrdered1d int `gorm:"column:attributed_units_ordered1d" db:"column:attributed_units_ordered1d" json:"attributed_units_ordered1d"`
}

