package auth

import (
	"log"
	"net/http"

	"github.com/diwaandrew/kredit/model"
)

type Service interface {
	GetNasabah() ([]model.Staging_Customers, int, error)
}

type service struct {
	repo NasabahRepository
}

func NewService(repo NasabahRepository) *service {
	return &service{repo}
}

func (s *service) GetNasabah() ([]model.Staging_Customers, int, error) {
	nasabah, err := s.repo.GetNasabah()
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}

	return nasabah, http.StatusOK, nil
}
