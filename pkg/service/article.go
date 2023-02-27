package service

import (
	"errors"
	"github.com/Abdullayev65/step_by_step_data/pkg/models"
)

func (s *Service) ArticleAdd(article *models.Article) error {
	_, err := s.DB.NewInsert().Model(article).Exec(s.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ArticleGet(id int) (*models.Article, error) {
	article := &models.Article{ID: id}
	err := s.DB.NewSelect().Model(article).WherePK().Scan(s.ctx)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (s *Service) ArticlesByUserID(userID int) *[]models.Article {
	articles := make([]models.Article, 0)
	s.DB.NewSelect().Model(&articles).
		Where("user_id = ?", userID).Scan(s.ctx)
	return &articles
}

func (s *Service) ArticlesByStep(step int) *[]models.Article {
	articles := make([]models.Article, 0)
	s.DB.NewSelect().Model(&articles).
		Where("step = ?", step).Scan(s.ctx)
	return &articles
}

func (s *Service) AllArticles() *[]models.Article {
	articles := make([]models.Article, 0)
	s.DB.NewSelect().Model(&articles).Scan(s.ctx)
	return &articles
}

func (s *Service) ArticleAccept(id int, step int) error {
	article := models.Article{ID: id}
	err := s.DB.NewSelect().Model(&article).Column("step").
		WherePK().Scan(s.ctx)
	if err != nil {
		return err
	}
	if step < int(article.Step) {
		return errors.New("user bu bosqichdan o'tkan")
	}
	article.Step++
	_, err = s.DB.NewUpdate().Model(&article).Column("step").
		WherePK().Exec(s.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ArticleReject(id int) error {
	article := models.Article{ID: id, Active: false}
	_, err := s.DB.NewUpdate().Model(&article).
		Set("active = ?", false).WherePK().Exec(s.ctx)

	if err != nil {
		return err
	}
	return nil
}
