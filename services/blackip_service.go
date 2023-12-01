package services

import (
	"lottery-system/dao"
	"lottery-system/models"
)

type BlackipService interface {
	Get(id int) *models.LtBlackip
	GetAll() []models.LtBlackip
	CountAll() int64
	Update(data *models.LtBlackip, columns []string) error
	Create(data *models.LtBlackip) error
	GetByIp(ip string) *models.LtBlackip
}

type blackipService struct {
	dao *dao.BlackipDao
}

func NewBlackipService() BlackipService {
	return &blackipService{
		dao: dao.NewBlackipDao(nil),
	}
}

func (s *blackipService) Get(id int) *models.LtBlackip {
	return s.dao.Get(id)
}
func (s *blackipService) GetAll() []models.LtBlackip {
	return s.dao.GetAll()
}
func (s *blackipService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *blackipService) Update(data *models.LtBlackip, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *blackipService) Create(data *models.LtBlackip) error {
	return s.dao.Create(data)
}
func (s *blackipService) GetByIp(ip string) *models.LtBlackip {
	return s.dao.GetByIp(ip)
}
