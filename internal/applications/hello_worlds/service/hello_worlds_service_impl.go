package service

import (
	"context"
	"github.com/labstack/gommon/log"
	"myapp/exceptions"
	"myapp/internal/applications/hello_worlds/repository"
	"strings"
)

type HelloWorldsServiceImpl struct {
	repository repository.HelloWorldsRepository
}

func NewHelloWorldsService(repository repository.HelloWorldsRepository) HelloWorldService {
	return &HelloWorldsServiceImpl{
		repository: repository,
	}
}

func (service *HelloWorldsServiceImpl) Hello(ctx context.Context, message string, errorFlag string) (string, error) {
	log.Info("ctx debug", ctx)

	messageService := message + " hello from service-impl layer -"
	result, err := service.repository.Hello(ctx, messageService, errorFlag)
	if strings.EqualFold(errorFlag, "service") {
		return "", exceptions.NewBusinessLogicError(exceptions.EBL10002, err)
	}

	return result, err
}
