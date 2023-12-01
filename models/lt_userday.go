package models

type LtUserday struct {
	Id         int `xorm:"not null pk autoincr comment('Primary Key') INT"`
	Uid        int `xorm:"comment('用户id') INT"`
	Day        int `xorm:"comment('日期，如：20230724') INT"`
	Num        int `xorm:"comment('次数') INT"`
	SysCreated int `xorm:"comment('创建时间') INT"`
	SysUpdated int `xorm:"comment('修改时间') INT"`
}
