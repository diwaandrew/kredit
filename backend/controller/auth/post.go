package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/diwaandrew/kredit/model"
)

func (r *repository) PostCDT(Custcode string, ppk string, name string, address1 string, address2 string,
	city string, zip string, birthPlace string, BirthDate string, ID_Type string, ID_Number string, MobileNo string,
	LoanTglPk string, TglPkChanneling string, MotherMaidenName string, ChannelingCompany string) {

	//FORMAT
	date, err := time.Parse("2006-01-02", BirthDate)
	if err != nil {
		fmt.Println(err)
	}
	idType, err := strconv.ParseInt(ID_Type, 10, 8)
	if err != nil {
		fmt.Println(err)
	}
	drawdownDate, err := time.Parse("2006-01-02", LoanTglPk)
	if err != nil {
		fmt.Println(err)
	}
	tglPkChanneling, err := time.Parse("2006-01-02", TglPkChanneling)
	if err != nil {
		fmt.Println(err)
	}

	CDT := model.Customer_Data_Tabs{
		Custcode:           Custcode,
		PPK:                ppk,
		Name:               name,
		Address1:           address1,
		Address2:           address2,
		City:               city,
		Zip:                zip,
		Birth_Place:        birthPlace,
		Birth_Date:         date,
		ID_Type:            int8(idType),
		ID_Number:          ID_Number,
		Mobile_No:          MobileNo,
		Drawdown_date:      drawdownDate,
		Tgl_pk_channelling: tglPkChanneling,
		Mother_maiden_name: MotherMaidenName,
		Channeling_Company: ChannelingCompany,
		Approval_Status:    "9",
	}
	r.db.Create(&CDT)
}

func (r *repository) PostLDT(Custcode string, Branch string, OTR string, DownPayment string, LoanAmount string, LoanPeriod string,
	LoanInterestFlatChanneling string, LoanInterestEffectiveChanneling string, LoanEffectivePaymentType string, MonthlyPayment string,
	inputDate time.Time, InputDate2 time.Time) {

	OTR64, err := strconv.ParseFloat(OTR, 64)
	if err != nil {
		fmt.Println(err)
	}

	DownPayment64, err := strconv.ParseFloat(DownPayment, 64)
	if err != nil {
		fmt.Println(err)
	}

	LoanAmount64, err := strconv.ParseFloat(LoanAmount, 64)
	if err != nil {
		fmt.Println(err)
	}

	InterestFlat, err := strconv.ParseFloat(LoanInterestFlatChanneling, 32)
	if err != nil {
		fmt.Println(err)
	}

	InterestEffective, err := strconv.ParseFloat(LoanInterestEffectiveChanneling, 32)
	if err != nil {
		fmt.Println(err)
	}

	EffectivePaymentType, err := strconv.ParseInt(LoanEffectivePaymentType, 10, 8)
	if err != nil {
		fmt.Println(err)
	}

	MonthlyPayment64, err := strconv.ParseFloat(MonthlyPayment, 64)
	if err != nil {
		fmt.Println(err)
	}

	LDT := model.Loan_Data_Tabs{
		Custcode:             Custcode,
		Branch:               Branch,
		OTR:                  float64(OTR64),
		DownPayment:          float64(DownPayment64),
		LoanAmount:           float64(LoanAmount64),
		LoanPeriod:           LoanPeriod,
		InterestType:         1,
		InterestFlat:         float32(InterestFlat),
		InterestEffective:    float32(InterestEffective),
		EffectivePaymentType: int8(EffectivePaymentType),
		AdminFee:             30,
		MonthlyPayment:       float64(MonthlyPayment64),
		InputDate:            inputDate,
		LastModified:         time.Now(),
		ModifiedBy:           "system",
		InputDate2:           InputDate2,
		InputBy:              "system",
		LastModified2:        time.Now(),
		ModifiedBy2:          "system",
	}
	r.db.Create(&LDT)
}

func (r *repository) PostVDT(Custcode string, Brand string, Type string, Year string, Golongan int8, Jenis string, Status string,
	Color string, PoliceNo string, EngineNo string, ChasisNo string, Bpkb string, RegisterNo string, Stnk string, StnkAddress1 string,
	StnkAddress2 string, StnkCity string, DealerID string, InputDate time.Time, Inputby string,
	LastModified time.Time, Modifiedby string, TglStnk string, TglBpkb string, TglPolis time.Time, PolisNo string, CollateralID string,
	Ketagunan string, AgunanLbu string, Dealer string, AddressDealer1 string, AddressDealer2 string, CityDealer string) {

	VehicleBrand8, err := strconv.ParseInt(Brand, 10, 8)
	if err != nil {
		fmt.Println(err)
	}

	Status8, err := strconv.ParseInt(Status, 10, 8)
	if err != nil {
		fmt.Println(err)
	}

	VehicleDealerID, err := strconv.ParseInt(DealerID, 10, 8)
	if err != nil {
		fmt.Println(err)
	}

	VehicleTglStnk, err := time.Parse("2006-01-02 15:04:05", TglStnk)
	if err != nil {
		fmt.Println(err)
	}

	VehicleTglBpkb, err := time.Parse("2006-01-02 15:04:05", TglBpkb)
	if err != nil {
		fmt.Println(err)
	}

	CollateralTypeID, err := strconv.ParseInt(CollateralID, 10, 8)
	if err != nil {
		fmt.Println(err)
	}

	VDT := model.Vehicle_Data_Tabs{
		Custcode:       Custcode,
		Brand:          int(VehicleBrand8),
		Type:           Type,
		Year:           Year,
		Golongan:       int8(Golongan),
		Jenis:          Jenis,
		Status:         int8(Status8),
		Color:          Color,
		PoliceNo:       PoliceNo,
		EngineNo:       EngineNo,
		ChasisNo:       ChasisNo,
		Bpkb:           Bpkb,
		RegisterNo:     RegisterNo,
		Stnk:           Stnk,
		StnkAddress1:   StnkAddress1,
		StnkAddress2:   StnkAddress2,
		StnkCity:       StnkCity,
		DealerID:       int(VehicleDealerID),
		Inputdate:      InputDate,
		Inputby:        Inputby,
		Lastmodified:   LastModified,
		Modifiedby:     Modifiedby,
		TglStnk:        VehicleTglStnk,
		TglBpkb:        VehicleTglBpkb,
		TglPolis:       TglPolis,
		PolisNo:        PolisNo,
		CollateralID:   int64(CollateralTypeID),
		Ketagunan:      Ketagunan,
		AgunanLbu:      AgunanLbu,
		Dealer:         Dealer,
		AddressDealer1: AddressDealer1,
		AddressDealer2: AddressDealer2,
		CityDealer:     CityDealer,
	}
	r.db.Create(&VDT)
}
