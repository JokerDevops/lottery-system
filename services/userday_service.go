package services

import (
	"lottery-system/dao"
	"lottery-system/models"
)

type UserdayService interface {
	Get(id int) *models.LtUserday
	GetAll() []models.LtUserday
	CountAll() int64
	Update(data *models.LtUserday, columns []string) error
	Create(data *models.LtUserday) error
}

type userdayService struct {
	dao *dao.UserdayDao
}

func NewUserdayService() UserdayService {
	return &userdayService{
		dao: dao.NewUserdayDao(nil),
	}
}

func (s *userdayService) Get(id int) *models.LtUserday {
	return s.dao.Get(id)
}

func (s *userdayService) GetAll() []models.LtUserday {
	return s.dao.GetAll()
}

func (s *userdayService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *userdayService) Update(data *models.LtUserday, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *userdayService) Create(data *models.LtUserday) error {
	return s.dao.Create(data)
}
