package auth

import (
	"log"
	"regexp"
	"time"

	"github.com/diwaandrew/kredit/model"
	"gorm.io/gorm"
)

type NasabahRepository interface {
	GetNasabah() ([]model.Staging_Customers, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetNasabah() ([]model.Staging_Customers, error) {
	var nasabah []model.Staging_Customers

	//WHEN
	res1 := r.db.Where("sc_flag= ?", "0").Find(&nasabah)

	if res1.Error != nil {
		log.Println("Get Data error : ", res1.Error)
		return nil, res1.Error
	}

	for _, item := range nasabah {

		//VALIDATION CHECK 1
		if !r.CheckingDuplicate(item.CustomerPpk) {
			r.CheckingError(item.ID, "Duplicate Field PPK", item.ScReff, item.ScCreateDate, item.ScBranchCode, item.ScCompany, item.CustomerPpk, item.CustomerName)
			continue
		}
		//VALIDATION CHECK 2
		if !r.CheckingCompany(item.ScCompany) {
			r.CheckingError(item.ID, "Company Tidak Terdaftar", item.ScReff, item.ScCreateDate, item.ScBranchCode, item.ScCompany, item.CustomerPpk, item.CustomerName)
			continue
		}
		//VALIDATION CHECK 3
		if !r.CheckingBranch(item.ScBranchCode) {
			r.CheckingError(item.ID, "Branch Code Tidak Terdaftar", item.ScReff, item.ScCreateDate, item.ScBranchCode, item.ScCompany, item.CustomerPpk, item.CustomerName)
			continue
		}
		//VALIDATION CHECK 4
		if item.CustomerIDType == "1" && item.CustomerIDNumber == "" {
			r.CheckingError(item.ID, "CUSTOMER_ID_NUMBER Kosong", item.ScReff, item.ScCreateDate, item.ScBranchCode, item.ScCompany, item.CustomerPpk, item.CustomerName)
			continue
		}
		//VALIDATION CHECK 5
		regex := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~]+`)
		if matched := regex.MatchString(item.CustomerName); matched {
			r.CheckingError(item.ID, "Nama Mengandung Simbol", item.ScReff, item.ScCreateDate, item.ScBranchCode, item.ScCompany, item.CustomerPpk, item.CustomerName)
			continue
		}

		//VALIDATION CHECK 6
		if item.VehicleBpkb == "" || item.VehicleStnk == "" || item.VehicleEngineNo == "" || item.VehicleChasisNo == "" {
			r.CheckingError(item.ID, "Data Kosong", item.ScReff, item.ScCreateDate, item.ScBranchCode, item.ScCompany, item.CustomerPpk, item.CustomerName)
			continue
		}
		//VALIDATION CHECK 7
		if !r.CheckingDuplicateVEN(item.VehicleEngineNo) {
			r.CheckingError(item.ID, "Duplicate VEHICLE_Engine_NO", item.ScReff, item.ScCreateDate, item.ScBranchCode, item.ScCompany, item.CustomerPpk, item.CustomerName)
			continue
		}
		//VALIDATION CHECK 8
		if !r.CheckingDuplicateVEC(item.VehicleChasisNo) {
			r.CheckingError(item.ID, "Duplicate VEHICLE_CHASIS_NO", item.ScReff, item.ScCreateDate, item.ScBranchCode, item.ScCompany, item.CustomerPpk, item.CustomerName)
			continue
		}

		CustcodeNew := r.GenerateCustcode(item.ScCompany)

		//POST DATA
		r.PostCDT(CustcodeNew, item.CustomerPpk, item.CustomerName, item.CustomerAddress1, item.CustomerAddress2, item.CustomerCity,
			item.CustomerZip, item.CustomerBirthPlace, item.CustomerBirthDate, item.CustomerIDType, item.CustomerIDNumber, item.CustomerMobileNo,
			item.LoanTglPk, item.LoanTglPkChanneling, item.CustomerMotherMaidenName, item.ScCompany)

		r.PostLDT(CustcodeNew, item.ScBranchCode, item.LoanOtr, item.LoanDownPayment, item.LoanLoanAmountChanneling, item.LoanLoanPeriodChanneling,
			item.LoanInterestFlatChanneling, item.LoanInterestEffectiveChanneling, item.LoanEffectivePaymentType,
			item.LoanMonthlyPaymentChanneling, item.ScCreateDate, item.ScCreateDate)

		r.PostVDT(CustcodeNew, item.VehicleBrand, item.VehicleType, item.VehicleYear, 1, item.VehicleJenis, item.VehicleStatus, item.VehicleColor,
			item.VehiclePoliceNo, item.VehicleEngineNo, item.VehicleChasisNo, item.VehicleBpkb, "1", item.VehicleStnk, "", "", "", item.VehicleDealerID,
			time.Now(), "system", time.Now(), "system", item.VehicleTglStnk, item.VehicleTglStnk, time.Now(), item.VehiclePoliceNo, item.CollateralTypeID, "", "",
			item.VehicleDealer, item.VehicleAddressDealer1, item.VehicleAddressDealer2, item.VehicleCityDealer)

		//UPDATE WHEN Checking VALID
		r.db.Model(&nasabah).Where("id=?", item.ID).Update("sc_flag", "1")

	}
	return nasabah, nil
}
