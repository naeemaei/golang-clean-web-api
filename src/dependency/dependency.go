package dependency

import (
	"github.com/naeemaei/golang-clean-web-api/config"
	"github.com/naeemaei/golang-clean-web-api/domain/model"
	contractRepository "github.com/naeemaei/golang-clean-web-api/domain/repository"
	infraRepository "github.com/naeemaei/golang-clean-web-api/infra/persistence/repository"
)

func GetUserRepository(cfg *config.Config) contractRepository.UserRepository {
	return infraRepository.NewUserRepository(cfg)
}

func GetPersianYearRepository(cfg *config.Config) contractRepository.PersianYearRepository {
	return infraRepository.NewBaseRepository[model.PersianYear](cfg)
}
