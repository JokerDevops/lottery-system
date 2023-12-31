package dao

import (
	"log"
	"lottery-system/models"

	"github.com/go-xorm/xorm"
)

type CodeDao struct {
	engine *xorm.Engine
}

func NewCodeDao(engine *xorm.Engine) *CodeDao {
	return &CodeDao{
		engine: engine,
	}
}

func (d *CodeDao) Get(id int) *models.LtCode {
	data := &models.LtCode{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *CodeDao) GetAll() []models.LtCode {
	datalist := make([]models.LtCode, 0)
	err := d.engine.Asc("sys_status").Asc("displayorder").Find(&datalist)
	if err != nil {
		log.Println("code_dao.GetAll error=", err)
		return datalist
	}
	return datalist
}

func (d *CodeDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtCode{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *CodeDao) Delete(id int) error {
	data := &models.LtCode{Id: id, SysStatus: 1}
	_, err := d.engine.ID(data.Id).Update(data)
	return err
}

func (d *CodeDao) Update(data *models.LtCode, columns []string) error {
	_, err := d.engine.ID(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *CodeDao) Create(data *models.LtCode) error {
	_, err := d.engine.Insert(data)
	return err
}

// 找到下一个可用的最小的优惠卷
func (d *CodeDao) NextUsingCode(giftId, codeId int) *models.LtCode {
	datalist := make([]models.LtCode, 0)
	err := d.engine.Where("gift_id=?", giftId).Where("sys_status=?", 0).Where("id>?", codeId).Asc("id").Limit(1).Find(&datalist)
	if err != nil || len(datalist) < 1 {
		return nil
	} else {
		return &datalist[0]
	}
}

// 根据唯一的 code 来更新
func (d *CodeDao) UpdateByCode(data *models.LtCode, columns []string) error {
	_, err := d.engine.Where("code=?", data.Code).MustCols(columns...).Update(data)
	return err
}
