package sd

type BaseMode struct {
	Zid int64 `gorm:"column:zid" db:"column:zid" json:"zid"`
	Sid int64 `gorm:"column:sid" db:"column:sid" json:"sid"`
	ProfileId string `gorm:"column:profile_id" db:"column:profile_id" json:"profile_id"`
	GmtModified string `gorm:"column:gmt_modified" db:"column:gmt_modified" json:"gmt_modified"`
	GmtCreate string `gorm:"column:gmt_create" db:"column:gmt_create" json:"gmt_create"`
}
