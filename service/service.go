package service

import (
	"article-manager/conf"
	"article-manager/dao"
)

type Service struct {
	c   *conf.Config
	dao *dao.Dao
}

func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:   c,
		dao: dao.New(c),
	}
	return
}

func (s *Service) Close() {

}
