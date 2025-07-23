package dto

type CreatePropertyCategoryRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=20"`
	Icon string `json:"icon" binding:"min=1,max=1000"`
}
type UpdatePropertyCategoryRequest struct {
	Name string `json:"name,omitempty"`
	Icon string `json:"icon,omitempty"`
}

type PropertyCategoryResponse struct {
	Id         int                `json:"id"`
	Name       string             `json:"name"`
	Icon       string             `json:"icon"`
	Properties []PropertyResponse `json:"properties,omitempty"`
}

type CreatePropertyRequest struct {
	Name        string `json:"name" binding:"required,alpha,min=3,max=60"`
	Icon        string `json:"icon" binding:"min=1,max=1000"`
	CategoryId  int    `json:"categoryId" binding:"required"`
	Description string `json:"description" binding:"max=1000"`
	DataType    string `json:"dataType" binding:"max=15"`
	Unit        string `json:"unit" binding:"max=15"`
}

type UpdatePropertyRequest struct {
	Name        string `json:"name,omitempty"`
	Icon        string `json:"icon,omitempty" binding:"max=1000"`
	CategoryId  int    `json:"categoryId,omitempty"`
	Description string `json:"description,omitempty" binding:"max=1000"`
	DataType    string `json:"dataType,omitempty" binding:"max=15"`
	Unit        string `json:"unit,omitempty" binding:"max=15"`
}

type PropertyResponse struct {
	Id          int                      `json:"id"`
	Name        string                   `json:"name"`
	Icon        string                   `json:"icon"`
	Category    PropertyCategoryResponse `json:"category,omitempty"`
	Description string                   `json:"description"`
	DataType    string                   `json:"dataType"`
	Unit        string                   `json:"unit"`
}
