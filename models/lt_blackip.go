package models

type LtBlackip struct {
	Id         int    `xorm:"not null pk autoincr comment('Primary Key') INT"`
	Ip         string `xorm:"comment('IP 地址') VARCHAR(50)"`
	Blacktime  int    `xorm:"comment('黑名单限制到期时间') INT"`
	SysCreated int    `xorm:"comment('创建时间') INT"`
	SysUpdated int    `xorm:"comment('修改时间') INT"`
}
