package services

import (
	"lottery-system/dao"
	"lottery-system/models"
)

type ResultService interface {
	Get(id int) *models.LtResult
	GetAll() []models.LtResult
	CountAll() int64
	SearchByGift(giftId, page, size int) []models.LtResult
	SearchByUser(uid, page, size int) []models.LtResult
	CountByGift(giftId int) int64
	CountByUser(uid int) int64
	Delete(id int) error
	Update(data *models.LtResult, columns []string) error
	Create(data *models.LtResult) error
}

type resultService struct {
	dao *dao.ResultDao
}

func NewResultService() ResultService {
	return &resultService{
		dao: dao.NewResultDao(nil),
	}

}
func (s *resultService) Get(id int) *models.LtResult {
	return s.dao.Get(id)
}
func (s *resultService) GetAll() []models.LtResult {
	return s.dao.GetAll()
}
func (s *resultService) CountAll() int64 {
	return s.dao.CountAll()
}
func (s *resultService) SearchByGift(giftId, page, size int) []models.LtResult {
	return s.dao.SearchByGift(giftId, page, size)
}
func (s *resultService) SearchByUser(uid, page, size int) []models.LtResult {
	return s.dao.SearchByUser(uid, page, size)
}
func (s *resultService) CountByGift(giftId int) int64 {
	return s.dao.CountByGift(giftId)
}
func (s *resultService) CountByUser(uid int) int64 {
	return s.dao.CountByUser(uid)
}
func (s *resultService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *resultService) Update(data *models.LtResult, columns []string) error {
	return s.dao.Update(data, columns)
}
func (s *resultService) Create(data *models.LtResult) error {
	return s.dao.Create(data)
}
