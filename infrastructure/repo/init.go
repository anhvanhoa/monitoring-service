package repo

import (
	"monitoring_service/domain/repository"

	"github.com/anhvanhoa/service-core/utils"
	"github.com/go-pg/pg/v10"
)

type Repositories interface {
	EnvironmentalAlertRepository() repository.EnvironmentalAlertRepository
}

type RepositoriesImpl struct {
	eaRepository repository.EnvironmentalAlertRepository
}

func InitRepositories(db *pg.DB) Repositories {
	helper := utils.NewHelper()
	return &RepositoriesImpl{
		eaRepository: NewEnvironmentalAlertRepository(db, helper),
	}
}

func (r *RepositoriesImpl) EnvironmentalAlertRepository() repository.EnvironmentalAlertRepository {
	return r.eaRepository
}
