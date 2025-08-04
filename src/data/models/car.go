package models

import "time"

type Gearbox struct {
	BaseModel
	Name string `gorm:"size:15;not null;unique"`
}

type CarType struct {
	BaseModel
	Name      string     `gorm:"size:15;not null;unique"`
	CarModels []CarModel `gorm:"foreignKey:CarTypeId"`
}

type Company struct {
	BaseModel
	Name      string `gorm:"size:15;not null;unique"`
	CountryId int
	Country   Country `gorm:"foreignKey:CountryId"`
	CarModels []CarModel
}

type CarModel struct {
	BaseModel
	Name               string             `gorm:"size:15;not null;unique" json:"name"`
	CompanyId          int                `json:"companyId"`
	Company            Company            `gorm:"foreignKey:CompanyId" json:"company"`
	CarTypeId          int                `json:"carTypeId"`
	CarType            CarType            `gorm:"foreignKey:CarTypeId" json:"carType"`
	GearboxId          int                `json:"gearboxId"`
	Gearbox            Gearbox            `gorm:"foreignKey:GearboxId" json:"gearbox"`
	CarModelColors     []CarModelColor    `gorm:"foreignKey:CarModelId" json:"carModelColors"`
	CarModelYears      []CarModelYear     `gorm:"foreignKey:CarModelId" json:"carModelYears"`
	CarModelProperties []CarModelProperty `gorm:"foreignKey:CarModelId" json:"carModelProperties"`
	CarModelImages     []CarModelImage    `gorm:"foreignKey:CarModelId" json:"carModelImages"`
	CarModelComments   []CarModelComment  `gorm:"foreignKey:CarModelId" json:"carModelComments"`
}

type CarModelColor struct {
	BaseModel
	CarModelId int
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	ColorId    int
	Color      Color `gorm:"foreignKey:ColorId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
}

type CarModelYear struct {
	BaseModel
	CarModelId             int                    `gorm:"uniqueIndex:idx"`
	CarModel               CarModel               `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	PersianYearId          int                    `gorm:"uniqueIndex:idx"`
	PersianYear            PersianYear            `gorm:"foreignKey:PersianYearId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	CarModelPriceHistories []CarModelPriceHistory `gorm:"foreignKey:CarModelYearId"`
}

type CarModelImage struct {
	BaseModel
	CarModelId  int
	CarModel    CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	ImageId     int
	Image       File `gorm:"foreignKey:ImageId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	IsMainImage bool `gorm:"not null"`
}

type CarModelPriceHistory struct {
	BaseModel
	CarModelYear   CarModelYear `gorm:"foreignKey:CarModelYearId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	CarModelYearId int
	Price          float64   `gorm:"type:decimal(10,2);not null"`
	PriceAt        time.Time `gorm:"type:timestamp with time zone;not null"`
}

type CarModelProperty struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	CarModelId int
	Property   Property `gorm:"foreignKey:PropertyId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	PropertyId int
	Value      string `gorm:"size:100;not null"`
}

type CarModelComment struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	CarModelId int
	UserId     int  `gorm:"column:user_id;not null"` 
	User       User `gorm:"foreignKey:UserId;references:Id"`
	Message string `gorm:"size:500;not null"`
}
