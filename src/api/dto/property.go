package dto

import "github.com/naeemaei/golang-clean-web-api/usecase/dto"

type CreatePropertyCategoryRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=50"`
	Icon string `json:"icon" binding:"max=1000"`
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
	Name        string `json:"name" binding:"required,alpha,min=3,max=50"`
	CategoryId  int    `json:"categoryId" binding:"required"`
	Icon        string `json:"icon" binding:"max=1000"`
	Description string `json:"description" binding:"max=1000"`
	DataType    string `json:"dataType" binding:"max=15"`
	Unit        string `json:"unit" binding:"max=15"`
}

type UpdatePropertyRequest struct {
	Name        string `json:"name,omitempty"`
	CategoryId  int    `json:"categoryId,omitempty"`
	Icon        string `json:"icon,omitempty" binding:"max=1000"`
	Description string `json:"description,omitempty" binding:"max=1000"`
	DataType    string `json:"dataType,omitempty" binding:"max=15"`
	Unit        string `json:"unit,omitempty" binding:"max=15"`
}

type PropertyResponse struct {
	Id          int                      `json:"id"`
	Name        string                   `json:"name"`
	Icon        string                   `json:"icon"`
	Description string                   `json:"description"`
	DataType    string                   `json:"dataType"`
	Unit        string                   `json:"unit"`
	Category    PropertyCategoryResponse `json:"category,omitempty"`
}

func ToPropertyResponse(from dto.Property) PropertyResponse {
	return PropertyResponse{
		Id:          from.Id,
		Name:        from.Name,
		Icon:        from.Icon,
		DataType:    from.DataType,
		Unit:        from.Unit,
		Category:    ToPropertyCategoryResponse(from.Category),
		Description: from.Description,
	}
}

func ToCreateProperty(from CreatePropertyRequest) dto.CreateProperty {
	return dto.CreateProperty{
		Name:        from.Name,
		Icon:        from.Icon,
		DataType:    from.DataType,
		Unit:        from.Unit,
		CategoryId:  from.CategoryId,
		Description: from.Description,
	}
}

func ToUpdateProperty(from UpdatePropertyRequest) dto.UpdateProperty {
	return dto.UpdateProperty{
		Name:        from.Name,
		Icon:        from.Icon,
		DataType:    from.DataType,
		Unit:        from.Unit,
		CategoryId:  from.CategoryId,
		Description: from.Description,
	}
}

func ToPropertyCategoryResponse(from dto.PropertyCategory) PropertyCategoryResponse {
	properties := []PropertyResponse{}
	for _, item := range from.Properties {
		properties = append(properties, ToPropertyResponse(item))
	}
	return PropertyCategoryResponse{
		Id:         from.Id,
		Name:       from.Name,
		Icon:       from.Icon,
		Properties: properties,
	}
}

func ToCreatePropertyCategory(from CreatePropertyCategoryRequest) dto.CreatePropertyCategory {
	return dto.CreatePropertyCategory{
		Name: from.Name,
		Icon: from.Icon,
	}
}

func ToUpdatePropertyCategory(from UpdatePropertyCategoryRequest) dto.UpdatePropertyCategory {
	return dto.UpdatePropertyCategory{
		Name: from.Name,
		Icon: from.Icon,
	}
}
