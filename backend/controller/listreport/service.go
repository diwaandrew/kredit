package listreport

import (
	"log"
	"net/http"

	"github.com/diwaandrew/kredit/model"
)

type Service interface {
	GetListReport(data GetNasabah) ([]response, int, error)
	GetBranch() ([]model.Branch_Tabs, int, error)
	GetCompany() ([]model.Mst_Company_Tabs, int, error)
	SearchListReport(data GetSearchRequest) ([]response, int, error)
	UpdateCustomer(req []requestbody) (int, error)
}
type service struct {
	repo ListRepository
}

func NewService(repo ListRepository) *service {
	return &service{repo}
}

func (s *service) GetListReport(data GetNasabah) ([]response, int, error) {
	listReport, err := s.repo.GetListReport(data.StatusTrx)
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}
	return listReport, http.StatusOK, nil
}

func (s *service) GetBranch() ([]model.Branch_Tabs, int, error) {
	Branch, err := s.repo.GetBranch()
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}

	return Branch, http.StatusOK, nil
}

func (s *service) GetCompany() ([]model.Mst_Company_Tabs, int, error) {
	Company, err := s.repo.GetCompany()
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}

	return Company, http.StatusOK, nil
}

func (s *service) UpdateCustomer(req []requestbody) (int, error) {
	err := s.repo.UpdateCustomer(req)
	if err != nil {
		log.Println("Internal server error : ", err)
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (s *service) SearchListReport(data GetSearchRequest) ([]response, int, error) {
	listReport, err := s.repo.SearchListReport(data.Branch, data.Company, data.StartDate, data.EndDate, data.StatusTrx)
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}
	return listReport, http.StatusOK, nil
}
