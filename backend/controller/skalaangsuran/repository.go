package skalaangsuran

import (
	"log"
	"math"
	"strconv"
	"time"

	"github.com/diwaandrew/kredit/model"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	GenerateSkalaAngsuran() ([]model.Customer_Data_Tabs, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GenerateSkalaAngsuran() ([]model.Customer_Data_Tabs, error) {
	var CDT []model.Customer_Data_Tabs
	var LDT model.Loan_Data_Tabs

	res := r.db.Where("approval_status = ?", "9").Find(&CDT)
	if res.Error != nil {
		log.Println("Get Data Error CDT : ", res.Error)
		return nil, res.Error
	}

	for _, item := range CDT {
		res := r.db.Where("custcode = ? ", item.Custcode).Find(&LDT)
		if res.Error != nil {
			log.Print("Get Data Error LDT :", res.Error)
			return nil, res.Error
		}

		Periode, err := strconv.Atoi(LDT.LoanPeriod)
		if err != nil {
			panic(err)
		}

		OSBalance := LDT.LoanAmount
		MonthlyPayment := LDT.MonthlyPayment
		CurrentDate := time.Now()

		for i := 0; i <= Periode; i++ {
			if i == 0 {
				r.PostSkala(item.Custcode, i, OSBalance, OSBalance, CurrentDate, LDT.InterestEffective, MonthlyPayment, 0, 0, CurrentDate)
			} else {
				Interest := math.Round(OSBalance * float64(LDT.InterestEffective) * 30 / 36000)
				Principle := math.Round(MonthlyPayment - Interest)
				EndBalance := math.Round(OSBalance - Principle)
				DueDate := CurrentDate.AddDate(0, i, 0)

				if EndBalance < 0 {
					EndBalance = 0
				}

				r.PostSkala(item.Custcode, i, OSBalance, EndBalance, DueDate, LDT.InterestEffective, MonthlyPayment, Principle, Interest, CurrentDate)
				OSBalance = EndBalance
			}
		}
	}
	return CDT, nil
}

func (r *repository) PostSkala(custcode string, counter int, osbalance float64, endbalance float64,
	duedate time.Time, eff float32, rental float64, principle float64, interest float64, datenow time.Time) {

	Skala := model.Skala_Rental_Tabs{
		Custcode:    custcode,
		Counter:     int8(counter),
		Osbalance:   float64(osbalance),
		End_Balance: float64(endbalance),
		Due_Date:    duedate,
		Eff_Rate:    float32(eff),
		Rental:      float64(rental),
		Principle:   float64(principle),
		Interest:    float64(interest),
		InputDate:   datenow,
	}
	r.db.Create(&Skala)
}
