package models

type LtUser struct {
	Id         int    `xorm:"not null pk autoincr comment('Primary Key') INT"`
	Username   string `xorm:"comment('用户名') VARCHAR(50)"`
	Blacktime  int    `xorm:"comment('黑名单限制时间') INT"`
	Realname   string `xorm:"comment('联系人') VARCHAR(50)"`
	Mobile     string `xorm:"comment('手机号') VARCHAR(50)"`
	Address    string `xorm:"comment('联系地址') VARCHAR(255)"`
	SysCreated int    `xorm:"comment('创建时间') INT"`
	SysUpdated int    `xorm:"comment('修改时间') INT"`
	SysIp      string `xorm:"comment('IP 地址') VARCHAR(50)"`
}
