package dependency

import (
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/domain/model"
	contractRepository "github.com/naeemaei/golang-clean-web-api/domain/repository"
	database "github.com/naeemaei/golang-clean-web-api/infra/persistence/database"
	infraRepository "github.com/naeemaei/golang-clean-web-api/infra/persistence/repository"
)

func GetUserRepository(cfg *config.Config) contractRepository.UserRepository {
	return infraRepository.NewUserRepository(cfg)
}

func GetPersianYearRepository(cfg *config.Config) contractRepository.PersianYearRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{}
	return infraRepository.NewBaseRepository[model.PersianYear](cfg, preloads)
}

func GetCountryRepository(cfg *config.Config) contractRepository.CountryRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{{Entity: "Cities"}, {Entity: "Companies"}}
	return infraRepository.NewBaseRepository[model.Country](cfg, preloads)
}

func GetCarModelColorRepository(cfg *config.Config) contractRepository.CarModelColorRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{{Entity: "Color"}}
	return infraRepository.NewBaseRepository[model.CarModelColor](cfg, preloads)
}

func GetCarModelCommentRepository(cfg *config.Config) contractRepository.CarModelCommentRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{{Entity: "User"}}
	return infraRepository.NewBaseRepository[model.CarModelComment](cfg, preloads)
}

func GetCarModelImageRepository(cfg *config.Config) contractRepository.CarModelImageRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{{Entity: "Image"}}
	return infraRepository.NewBaseRepository[model.CarModelImage](cfg, preloads)
}

func GetCarModelPriceHistoryRepository(cfg *config.Config) contractRepository.CarModelPriceHistoryRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{}
	return infraRepository.NewBaseRepository[model.CarModelPriceHistory](cfg, preloads)
}

func GetCarModelPropertyRepository(cfg *config.Config) contractRepository.CarModelPropertyRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{{Entity: "Property.Category"}}
	return infraRepository.NewBaseRepository[model.CarModelProperty](cfg, preloads)
}

func GetCarModelRepository(cfg *config.Config) contractRepository.CarModelRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{
		{Entity: "Company.Country"},
		{Entity: "CarType"},
		{Entity: "Gearbox"},
		{Entity: "CarModelColors.Color"},
		{Entity: "CarModelYears.PersianYear"},
		{Entity: "CarModelYears.CarModelPriceHistories"},
		{Entity: "CarModelProperties.Property.Category"},
		{Entity: "CarModelImages.Image"},
		{Entity: "CarModelComments.User"},
	}
	return infraRepository.NewBaseRepository[model.CarModel](cfg, preloads)
}

func GetCarModelYearRepository(cfg *config.Config) contractRepository.CarModelYearRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{{Entity: "PersianYear"}}
	return infraRepository.NewBaseRepository[model.CarModelYear](cfg, preloads)
}

func GetCarTypeRepository(cfg *config.Config) contractRepository.CarTypeRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{}
	return infraRepository.NewBaseRepository[model.CarType](cfg, preloads)
}

func GetCityRepository(cfg *config.Config) contractRepository.CityRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{{Entity: "Country"}}
	return infraRepository.NewBaseRepository[model.City](cfg, preloads)
}

func GetColorRepository(cfg *config.Config) contractRepository.ColorRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{}
	return infraRepository.NewBaseRepository[model.Color](cfg, preloads)
}

func GetCompanyRepository(cfg *config.Config) contractRepository.CompanyRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{{Entity: "Country"}}
	return infraRepository.NewBaseRepository[model.Company](cfg, preloads)
}

func GetFileRepository(cfg *config.Config) contractRepository.FileRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{}
	return infraRepository.NewBaseRepository[model.File](cfg, preloads)
}

func GetGearboxRepository(cfg *config.Config) contractRepository.GearboxRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{}
	return infraRepository.NewBaseRepository[model.Gearbox](cfg, preloads)
}

func GetPropertyCategoryRepository(cfg *config.Config) contractRepository.PropertyCategoryRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{{Entity: "Properties"}}
	return infraRepository.NewBaseRepository[model.PropertyCategory](cfg, preloads)
}

func GetPropertyRepository(cfg *config.Config) contractRepository.PropertyRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{{Entity: "Category"}}
	return infraRepository.NewBaseRepository[model.Property](cfg, preloads)
}

func GetRoleRepository(cfg *config.Config) contractRepository.RoleRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{}
	return infraRepository.NewBaseRepository[model.Role](cfg, preloads)
}
