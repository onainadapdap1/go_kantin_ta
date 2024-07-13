package menumakanan

import (
	"github.com/onainadapdap1/golang_kantin/internal/repository/menumakanan"
	"github.com/onainadapdap1/golang_kantin/models"
)

type MenuMakananService interface {
	CreateMenuMakanan(menu models.MenuMakanan) error
	GetAllMenuMakanan() ([]models.MenuMakanan, error)
	DeleteMenuMakanan(id uint) error 
}

type menuMakananService struct {
	repo menumakanan.MenuMakananRepository
}

func NewMenuMakananServ(repo menumakanan.MenuMakananRepository) MenuMakananService {
	return &menuMakananService{repo: repo}
}

func (s *menuMakananService) DeleteMenuMakanan(id uint) error {
	err := s.repo.DeleteMenuMakanan(id)
	if err != nil {
		return err
	}

	return nil
} 

func (s *menuMakananService) CreateMenuMakanan(menu models.MenuMakanan) error {
	return s.repo.CreateMenuMakanan(menu)
}

func (s *menuMakananService) GetAllMenuMakanan() ([]models.MenuMakanan, error) {
	menusMakanans, err := s.repo.GetAllMenuMakanan()
	if err != nil {
		return nil, err
	}
	return menusMakanans, nil
}