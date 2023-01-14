package skalaangsuran

import (
	"log"
	"net/http"

	"github.com/diwaandrew/kredit/model"
)

type Service interface {
	GenerateSkalaAngsuran() ([]model.Customer_Data_Tabs, int, error)
}

type service struct {
	repo CustomerRepository
}

func NewService(repo CustomerRepository) *service {
	return &service{repo}
}

func (s *service) GenerateSkalaAngsuran() ([]model.Customer_Data_Tabs, int, error) {
	customer, err := s.repo.GenerateSkalaAngsuran()
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}

	return customer, http.StatusOK, nil
}
