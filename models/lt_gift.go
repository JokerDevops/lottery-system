package models

type LtGift struct {
	Id           int    `xorm:"not null pk autoincr comment('Primary Key') INT"`
	Title        string `xorm:"comment('奖品名称') VARCHAR(255)"`
	PrizeNum     int    `xorm:"comment('奖品数量，0 无限量，>0 限量，<0 无奖品') INT"`
	LeftNum      int    `xorm:"comment('剩余奖品数量') INT"`
	PrizeCode    string `xorm:"comment('0-9999 表示百分之一百，0-0 表示万分之一的中奖概率') VARCHAR(50)"`
	PrizeTime    int    `xorm:"comment('发奖周期，D 天') INT"`
	Img          string `xorm:"comment('奖品图片') VARCHAR(255)"`
	Displayorder int    `xorm:"comment('位置序号，小的排在前面') INT"`
	Gtype        int    `xorm:"comment('奖品类型，0 虚拟币，1 虚拟卷，2 实物-小奖，3实物-大奖') INT"`
	Gdata        string `xorm:"comment('拓展数据，如：虚拟币数量') VARCHAR(255)"`
	TimeBegin    int    `xorm:"comment('开始时间') INT"`
	TimeEnd      int    `xorm:"comment('结束时间') INT"`
	PrizeData    string `xorm:"comment('发奖计划，[时间1，数量1],[时间2，数量2]') TEXT"`
	PrizeBegin   int    `xorm:"comment('发奖周期的开始') INT"`
	PrizeEnd     int    `xorm:"comment('发奖周期的结束') INT"`
	SysStatus    int    `xorm:"comment('状态，0 正常，1 删除') INT"`
	SysCreated   int    `xorm:"comment('创建时间') INT"`
	SysUpdated   int    `xorm:"comment('修改时间') INT"`
	SysIp        string `xorm:"comment('操作人 IP') VARCHAR(50)"`
}
