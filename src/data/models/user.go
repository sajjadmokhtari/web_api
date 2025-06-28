package models

type User struct {
	BaseModel
	UserName     string `gorm:"type:string;size:20;not null;unique"`
	FirstName    string `gorm:"type:string;size:15;null;unique"`
	LastName     string `gorm:"type:string;size:25;null;unique"`
	MobileNumber string `gorm:"type:string;size:11;null;unique;default:null"`
	Email        string `gorm:"type:string;size:64;null;unique;default:null"`
	Password     string `gorm:"type:string;size:64;not null"`
	Enabled      string `gorm:"default:true"`
	UserRole     *[]UserRole
}
