package impl

import (
	"context"
	"warrantyapi/repository"
	"warrantyapi/service"
)

func NewLogServiceImpl(logRepository *repository.LogRepository) service.LogService {
	return &logServiceImpl{
		LogRepository: *logRepository,
	}
}

type logServiceImpl struct {
	repository.LogRepository
}

// Create implements repository.LogService
func (service *logServiceImpl) Create(ctx context.Context, module string, detail string) {
	panic("unimplemented")
}
