package model

import "time"

type Stemming struct {
	Id                    int       `gorm:"column:id" json:"id"`
	StartTime             string    `gorm:"column:start_time" json:"start_time"`
	EndTime               string    `gorm:"column:end_time" json:"end_time"`
	PreviousPeriodSurplus float64   `gorm:"column:previous_period_surplus" json:"previous_period_surplus"`
	CurrentPeriodAog      float64   `gorm:"column:current_period_aog" json:"current_period_aog"`
	PeriodOfConsumption   float64   `gorm:"column:period_of_consumption" json:"period_of_consumption"`
	CurrentResidue        float64   `gorm:"column:current_residue" json:"current_residue"`
	DailyOfConsumption    float64   `gorm:"column:daily_of_consumption" json:"daily_of_consumption"`
	OnceAmountOfMud       float64   `gorm:"column:once_amount_of_mud" json:"once_amount_of_mud"`
	DailyNumberOfIron     float64   `gorm:"column:daily_number_of_iron" json:"daily_number_of_iron"`
	AverageIronOutTime    float64   `gorm:"column:average_iron_out_time" json:"average_iron_out_time"`
	PeriodIronMudConsume  float64   `gorm:"column:period_iron_mud_consume" json:"period_iron_mud_consume"`
	PeriodSumOfIron       float64   `gorm:"column:period_sum_of_iron" json:"period_sum_of_iron"`
	ContractPrice         float64   `gorm:"column:contract_price" json:"contract_price"`
	ConversionPrice       float64   `gorm:"column:conversion_price" json:"conversion_price"`
	Remark                string    `gorm:"column:remark" json:"remark"`
	OpenId                string    `gorm:"column:openid" json:"openid"`
	ClientId              int       `gorm:"column:client_id" json:"client_id"`
	CreateTime            time.Time `gorm:"column:create_time" json:"create_time"`
	ModifyTime            time.Time `gorm:"column:modify_time" json:"modify_time"`
}

type StemmingResponse struct {
	Id         int       `json:"id"`
	StartTime  string    `json:"start_time"`
	EndTime    string    `json:"end_time"`
	CreateTime time.Time `json:"create_time"`
	ModifyTime time.Time `json:"modify_time"`
}
