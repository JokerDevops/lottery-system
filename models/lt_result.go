package models

type LtResult struct {
	Id         int    `xorm:"not null pk autoincr comment('Primary Key') INT"`
	GiftId     int    `xorm:"comment('奖品 ID，关联 lt_gift表') index INT"`
	GiftName   string `xorm:"comment('奖品名称') VARCHAR(50)"`
	GiftType   int    `xorm:"comment('奖品类型，同 lt_gift.gtype') INT"`
	Uid        int    `xorm:"comment('用户 id') INT"`
	PrizeCode  int    `xorm:"comment('抽奖编号（4位的随机数）') INT"`
	GiftData   string `xorm:"comment('获奖信息') VARCHAR(255)"`
	SysCreated int    `xorm:"comment('创建时间') INT"`
	SysIp      string `xorm:"comment('用户抽奖的 IP') VARCHAR(50)"`
	SysStatus  int    `xorm:"comment('状态，0 正常，1 删除，2 作弊') INT"`
}
