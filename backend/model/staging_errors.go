package model

import "time"

type Staging_Errors struct {
	Id           int64     `json:"id" gorm:"type: bigint ;not null"`
	SeReff       string    `json:"se_reff" gorm:"type: varchar(50)"`
	SeCreateDate time.Time `json:"se_create_date" gorm:"type: timestamp"`
	BranchCode   string    `json:"branch_code" gorm:"type: varchar(50)"`
	Company      string    `json:"company" gorm:"type: varchar(50)"`
	Ppk          string    `json:"ppk" gorm:"type: varchar(50)"`
	Name         string    `json:"name" gorm:"type: varchar(50)"`
	ErrorDesc    string    `json:"error_desc" gorm:"type: varchar(3000)"`
}

func (m *Staging_Errors) TableName() string {
	return "staging_error"
}
