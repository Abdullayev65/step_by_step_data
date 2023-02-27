package service

import "github.com/Abdullayev65/step_by_step_data/pkg/models"

func (s *Service) UserByUsername(username string) (*models.User, error) {
	user := new(models.User)
	err := s.DB.NewSelect().Model(user).
		Where("username = ?", username).Scan(s.ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) UserAdd(user *models.User) error {
	_, err := s.DB.NewInsert().Model(user).Exec(s.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) IsAdmin(userID int) bool {
	user := &models.User{ID: userID}
	s.DB.NewSelect().Model(user).Column("admin").WherePK().Scan(s.ctx)
	return user.Admin
}
