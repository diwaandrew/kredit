package listreport

import (
	"log"
	"time"

	"github.com/diwaandrew/kredit/model"
	"gorm.io/gorm"
)

type ListRepository interface {
	GetListReport(statustrx string) ([]response, error)
	GetBranch() ([]model.Branch_Tabs, error)
	GetCompany() ([]model.Mst_Company_Tabs, error)
	SearchListReport(branch string, company string, startDate string, endDate string, statustrx string) ([]response, error)
	UpdateCustomer(req []requestbody) error
}

type repository struct {
	db *gorm.DB
}

type response struct {
	Name              string
	Ppk               string
	Otr               string
	Loan_Amount       string
	DrawdownDate      time.Time
	LoanPeriod        string
	InterestEffective float32
	MonthlyPayment    string
	CollateralID      int64
	Branch            string
	Company           string
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) UpdateCustomer(req []requestbody) error {
	for _, req2 := range req {
		res := r.db.Where("ppk=?", req2.Ppk).Updates(model.Customer_Data_Tabs{
			Approval_Status: "0",
		})
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}

func (r *repository) GetListReport(statustrx string) ([]response, error) {

	// r.db.raw("select * from view_dradown where approval_status=@approval", map[string]interface{"approval":"9"}).find(&data)

	res, err := r.db.Raw(`SELECT ppk,name,otr,loan_amount,drawdown_date,Loan_period,interest_effective,monthly_payment,collateral_id,Branch,channeling_company 
	FROM view_drawdown 
	WHERE approval_status = $1`, statustrx).Rows()
	listData := []response{}
	if err != nil {
		panic(err)
	}
	defer res.Close()

	for res.Next() {
		var ppk string
		var name string
		var otr string
		var loan_amount string
		var drawdownDate time.Time
		var LoanPeriod string
		var InterestEffective float32
		var MonthlyPayment string
		var CollateralID int64
		var Branch string
		var Company string
		if err := res.Scan(&ppk, &name, &otr, &loan_amount, &drawdownDate, &LoanPeriod, &InterestEffective, &MonthlyPayment, &CollateralID, &Branch, &Company); err != nil {
			panic(err)
		}
		customer := response{
			Ppk:               ppk,
			Name:              name,
			Otr:               otr,
			Loan_Amount:       loan_amount,
			DrawdownDate:      drawdownDate,
			LoanPeriod:        LoanPeriod,
			InterestEffective: InterestEffective,
			MonthlyPayment:    MonthlyPayment,
			CollateralID:      CollateralID,
			Branch:            Branch,
			Company:           Company,
		}
		listData = append(listData, customer)
	}

	return listData, nil
}

func (r *repository) GetBranch() ([]model.Branch_Tabs, error) {
	var Branch []model.Branch_Tabs
	res := r.db.Order("code asc").Find(&Branch)
	if res.Error != nil {
		log.Println("Get Data error : ", res.Error)
		return nil, res.Error
	}
	return Branch, nil
}

func (r *repository) GetCompany() ([]model.Mst_Company_Tabs, error) {

	var Company []model.Mst_Company_Tabs
	res := r.db.Order("company_code asc").Find(&Company)
	if res.Error != nil {
		log.Println("Get Data error : ", res.Error)
		return nil, res.Error
	}
	return Company, nil
}

func (r *repository) SearchListReport(branch string, company string, startDate string, endDate string, statustrx string) ([]response, error) {
	query1 := ""
	query2 := ""
	if branch == "" {
		query1 = "and ldt.branch like $1 "
		branch = "%%"
	} else {
		query1 = "and ldt.branch = $1 "
	}

	if company == "" {
		query2 = "and cdt.channeling_company like $2 "
		company = "%%"
	} else {
		query2 = "and cdt.channeling_company = $2 "
	}

	res, err := r.db.Raw(`SELECT cdt.ppk, cdt.name, ldt.otr, ldt.loan_amount, cdt.drawdown_date, ldt.loan_period, 
	ldt.interest_effective, ldt.monthly_payment, vdt.collateral_id,ldt.branch,cdt.channeling_company 
	FROM customer_data_tab cdt 
	LEFT JOIN Loan_Data_Tab ldt 
	ON cdt.custcode = ldt.custcode 
	LEFT JOIN vehicle_data_tab vdt 
	ON cdt.custcode = vdt.custcode 
	WHERE cdt.approval_status = $5 
	AND drawdown_date between $3 and $4 `+query1+query2, branch, company, startDate, endDate, statustrx).Rows()
	listData := []response{}
	if err != nil {
		panic(err)
	}
	defer res.Close()

	for res.Next() {
		var ppk string
		var name string
		var otr string
		var loan_amount string
		var drawdownDate time.Time
		var LoanPeriod string
		var InterestEffective float32
		var MonthlyPayment string
		var CollateralID int64
		var Branch string
		var Company string
		if err := res.Scan(&ppk, &name, &otr, &loan_amount, &drawdownDate, &LoanPeriod, &InterestEffective, &MonthlyPayment, &CollateralID, &Branch, &Company); err != nil {
			panic(err)
		}
		customer := response{
			Ppk:               ppk,
			Name:              name,
			Otr:               otr,
			Loan_Amount:       loan_amount,
			DrawdownDate:      drawdownDate,
			LoanPeriod:        LoanPeriod,
			InterestEffective: InterestEffective,
			MonthlyPayment:    MonthlyPayment,
			CollateralID:      CollateralID,
			Branch:            Branch,
			Company:           Company,
		}
		listData = append(listData, customer)
	}

	return listData, nil
}
