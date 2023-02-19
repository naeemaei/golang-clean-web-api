package dto

type CreateUpdatePropertyCategoryRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=20"`
}

type PropertyCategoryResponse struct {
	Id         int                `json:"id"`
	Name       string             `json:"name" binding:"required,min=3,max=15"`
	Icon       string             `json:"icon" binding:"min=1,max=1000"`
	Properties []PropertyResponse `json:"properties,omitempty"`
}

type CreateUpdatePropertyRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=20"`
}

type PropertyResponse struct {
	Id          int                      `json:"id"`
	Name        string                   `json:"name" binding:"required,min=3,max=40"`
	Icon        string                   `json:"icon" binding:"min=1,max=1000"`
	Description string                   `json:"description" binding:"min=1,max=200"`
	DataType    string                   `json:"dataType" binding:"min=0,max=15"`
	Unit        string                   `json:"unit" binding:"min=0,max=15"`
	Category    PropertyCategoryResponse `json:"category,omitempty"`
}
