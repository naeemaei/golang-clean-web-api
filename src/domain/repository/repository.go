package repository

import (
	"context"

	"github.com/naeemaei/golang-clean-web-api/domain/filter"
	"github.com/naeemaei/golang-clean-web-api/domain/model"
)

type BaseRepository[TEntity any] interface {
	Create(ctx context.Context, entity TEntity) (TEntity, error)
	Update(ctx context.Context, id int, entity map[string]interface{}) (TEntity, error)
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (TEntity, error)
	GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (int64, *[]TEntity, error)
}
type CountryRepository interface {
	BaseRepository[model.Country]
}

type CityRepository interface {
	BaseRepository[model.City]
	// Create(ctx context.Context, City model.City) (model.City, error)
	// Update(ctx context.Context, id int, City model.City) (model.City, error)
	// Delete(ctx context.Context, id int) error
	// GetById(ctx context.Context, id int) (model.City, error)
	// GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (int64, *[]model.City, error)
}

type PersianYearRepository interface {
	BaseRepository[model.PersianYear]
}

type ColorRepository interface {
	BaseRepository[model.Color]
}

type FileRepository interface {
	BaseRepository[model.File]
}

type GearboxRepository interface {
	BaseRepository[model.Gearbox]
}

type CarTypeRepository interface {
	BaseRepository[model.CarType]
}

type CompanyRepository interface {
	BaseRepository[model.Company]
}

type CarModelRepository interface {
	BaseRepository[model.CarModel]
}

type CarModelColorRepository interface {
	BaseRepository[model.CarModelColor]
}

type CarModelYearRepository interface {
	BaseRepository[model.CarModelYear]
}

type CarModelImageRepository interface {
	BaseRepository[model.CarModelImage]
}

type CarModelPriceHistoryRepository interface {
	BaseRepository[model.CarModelPriceHistory]
}

type CarModelPropertyRepository interface {
	BaseRepository[model.CarModelProperty]
}

type CarModelCommentRepository interface {
	BaseRepository[model.CarModelComment]
}

type PropertyCategoryRepository interface {
	BaseRepository[model.PropertyCategory]
}

type PropertyRepository interface {
	BaseRepository[model.Property]
}

type UserRepository interface {
	ExistsMobileNumber(ctx context.Context, mobileNumber string) (bool, error)
	ExistsUsername(ctx context.Context, username string) (bool, error)
	ExistsEmail(ctx context.Context, email string) (bool, error)
	FetchUserInfo(ctx context.Context, username string, password string) (model.User, error)
	GetDefaultRole(ctx context.Context) (roleId int, err error)
	CreateUser(ctx context.Context, u model.User) (model.User, error)
}

type RoleRepository interface {
	BaseRepository[model.Role]
}
