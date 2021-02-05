package factory

import (
	"github.com/erickrodrigs/codepix/codepix-go/application/usecase"
	"github.com/erickrodrigs/codepix/codepix-go/infrastructure/repository"
	"github.com/jinzhu/gorm"
)

func transactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	transactionRepository := repository.TransactionRepositoryDb{Db: database}

	transactionUsecase := usecase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixRepository:         pixRepository,
	}

	return transactionUsecase
}
