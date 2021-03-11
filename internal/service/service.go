package service

import "halfway_demo/internal/dao"

type Service struct {
	dao *dao.Dao
}

func New(d *dao.Dao) *Service {
	return &Service{dao: d}
}