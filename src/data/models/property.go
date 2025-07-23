package models

type PropertyCategory struct {
	BaseModel
	Name       string     `gorm:"size:60;not null;unique"`
	Icon       string     `gorm:"size:1000;not null"`
	Properties []Property `gorm:"foreignKey:CategoryId"`
}

type Property struct {
	BaseModel
	Name        string `gorm:"size:60;not null;unique"` // اسم ویژگی باید یکتا باشه
	Icon        string `gorm:"size:1000;not null"`      // آیکون نمادین، نه یکتا
	CategoryId  int
	Category    PropertyCategory `gorm:"foreignKey:CategoryId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	Description string           `gorm:"size:1000;not null"` // توضیح ممکنه برای چند ویژگی مشابه باشه
	DataType    string           `gorm:"size:15;not null"`   // مثل "bool" یا "string"، تکراری مجازه
	Unit        string           `gorm:"size:15;not null"`   // مثل "mm" یا "inch"، همون‌طور
}
