package models

import "time"

type Country struct {
	BaseModel
	Name      string `gorm:"size:15;type:string;not null"`
	Cities    []City `gorm:"foreignKey:CountryId"`
	Companies []Company
}

type City struct {
	BaseModel
	Name      string `gorm:"size:40;type:string;not null"`
	CountryId int
	Country   Country `gorm:"foreignKey:CountryId;constraint:OnUpdate:No ACTION:OnDelete:NO ACTION"`
}

type PersianYear struct {
	BaseModel
	PersianTitle  string    `gorm:"size:10;not null;unique"`
	Year          int       `gorm:"type:int;uniqueIndex;not null"`
	StartAt       time.Time `gorm:"type:TIMESTAMP with time zone; not null; unique"`
	EndAt         time.Time `gorm:"type:TIMESTAMP with time zone; not null; unique"`
	CarModelYears []CarModelYear
}

type Color struct {
	BaseModel
	Name          string `gorm:"size:15;not null;unique"`
	HexCode       string `gorm:"size:7;not null;unique"`
	CarModelColor []CarModelColor
}

type File struct {
	BaseModel
	Name        string `gorm:"size:100;not null"`
	Directory   string `gorm:"size:100;not null"`
	Description string `gorm:"size:500;not null"`
	MimeType    string `gorm:"size:20;not null"`
}
