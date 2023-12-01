package models

type LtCode struct {
	Id         int `xorm:"not null pk autoincr comment('Primary Key') INT"`
	GiftId     int `xorm:"comment('奖品 ID，关联 lt_gift 表') index INT"`
	Code       int `xorm:"comment('虚拟卷编码') INT"`
	SysCreated int `xorm:"comment('创建时间') INT"`
	SysUpdated int `xorm:"comment('更新时间') INT"`
	SysStatus  int `xorm:"comment('状态，0 正常，1 作废，2 已发放') INT"`
}
