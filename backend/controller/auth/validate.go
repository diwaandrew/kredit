package auth

import (
	"fmt"
	"time"

	"github.com/diwaandrew/kredit/model"
)

const (
	YYYYMM = "200601"
)

//1. Duplicate PPK
func (r *repository) CheckingDuplicate(ppk string) bool {
	var customer model.Customer_Data_Tabs
	//Search top 1
	res2 := r.db.Where("ppk= ?", ppk).First(&customer)
	if res2.Error != nil {
		return true
	} else {
		return false
	}
}

//2. Company Already
func (r *repository) CheckingCompany(ScCompany string) bool {
	var company model.Mst_Company_Tabs

	res3 := r.db.Where("company_short_name = ?", ScCompany).First(&company)

	if res3.Error != nil {
		return false
	} else {
		return true
	}
}

//3. Branch Code
func (r *repository) CheckingBranch(BranchCode string) bool {
	var branch model.Branch_Tabs

	res4 := r.db.Where("code = ?", BranchCode).First(&branch)

	if res4.Error != nil {
		return false
	} else {
		return true
	}
}

//7. VEHICLE_ENGINE_NO
func (r *repository) CheckingDuplicateVEN(engineNo string) bool {
	var engine model.Vehicle_Data_Tabs

	res5 := r.db.Where("engine_no= ?", engineNo).First(&engine)
	if res5.Error != nil {
		return true
	} else {
		return false
	}
}

//8. VEHICLE_ENGINE_CHASIS
func (r *repository) CheckingDuplicateVEC(chasisNo string) bool {
	var chasis model.Vehicle_Data_Tabs

	res6 := r.db.Where("engine_no= ?", chasisNo).First(&chasis)
	if res6.Error != nil {
		return true
	} else {
		return false
	}
}

//CREATE CUSTCODE
func (r *repository) GenerateCustcode(ScCompany string) string {
	var ID_Tab model.ID_Tabs
	var Company model.Mst_Company_Tabs

	r.db.Where("company_short_name = ? ", ScCompany).First(&Company)
	companyCode := Company.Company_Code

	dateget := time.Now()

	r.db.Where("company_code = ? ", companyCode).First(&ID_Tab)
	appCode := ID_Tab.Code
	codeMid := ID_Tab.MiddleCode
	CustCodeSeq := ID_Tab.Value
	// CustCodeLen := ID_Tab.Digit
	counter := codeMid + fmt.Sprintf("%d", CustCodeSeq)

	ResultCust := appCode + companyCode + dateget.Format(YYYYMM) + counter
	CustCodeSeq += 1
	r.db.Model(&ID_Tab).Where("code=?", appCode).Update("value", CustCodeSeq)

	return ResultCust

}

//IF ERROR UPDATE FLAG 8
func (r *repository) CheckingError(ID int64, desc string, ScReff string, ScCreateDate time.Time, ScBranchCode string, ScCompany string,
	CustomerPpk string, CustomerName string) {

	var nasabah model.Staging_Customers
	r.db.Model(&nasabah).Where("id=?", ID).Update("sc_flag", "8")

	nasabahFailed := model.Staging_Errors{
		Id:           ID,
		SeReff:       ScReff,
		SeCreateDate: ScCreateDate,
		BranchCode:   ScBranchCode,
		Company:      ScCompany,
		Ppk:          CustomerPpk,
		Name:         CustomerName,
		ErrorDesc:    desc,
	}
	r.db.Create(&nasabahFailed)
}
