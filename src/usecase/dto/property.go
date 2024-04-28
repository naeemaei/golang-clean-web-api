package dto

type CreatePropertyCategory struct {
	Name string
	Icon string
}

type UpdatePropertyCategory struct {
	Name string
	Icon string
}

type PropertyCategory struct {
	IdName
	Icon       string
	Properties []Property
}

type CreateProperty struct {
	Name        string
	CategoryId  int
	Icon        string
	Description string
	DataType    string
	Unit        string
}

type UpdateProperty struct {
	Name        string
	CategoryId  int
	Icon        string
	Description string
	DataType    string
	Unit        string
}

type Property struct {
	IdName
	Icon        string
	Description string
	DataType    string
	Unit        string
	Category    PropertyCategory
}
