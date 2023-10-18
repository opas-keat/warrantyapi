package impl

import (
	"context"
	"warrantyapi/constant"
	"warrantyapi/entity"
	"warrantyapi/model"
	"warrantyapi/repository"
	"warrantyapi/service"
)

func NewConfigServiceImpl(
	configRepository *repository.ConfigRepository,
	logRepository *repository.LogRepository,
) service.ConfigService {
	return &configServiceImpl{
		ConfigRepository: *configRepository,
		LogRepository:    *logRepository,
	}
}

type configServiceImpl struct {
	repository.ConfigRepository
	repository.LogRepository
}

// Update implements service.ConfigService
func (service *configServiceImpl) Update(ctx context.Context, searchRequest []model.ConfigRequest, updatedBy string) []model.ConfigResponse {
	searchEntity := entity.Config{
		ConfigCode: searchRequest[0].ConfigCode,
	}
	results := service.ConfigRepository.List(ctx, 0, 1, "", searchEntity)
	if len(results) == 0 {
		return []model.ConfigResponse{}
	}

	var configs []entity.Config
	for _, rs := range results {
		configs = append(configs, entity.Config{
			ID:           rs.ID,
			ConfigCode:   rs.ConfigCode,
			ConfigDetail: rs.ConfigDetail,
			ConfigValue:  searchRequest[0].ConfigValue,
		})
	}
	service.ConfigRepository.Update(ctx, configs)
	var configsResponse []model.ConfigResponse
	for _, rs := range configs {
		service.LogRepository.Insert(ctx, entity.Log{
			CreatedBy: updatedBy,
			Module:    constant.ModuleConfig,
			Detail:    "แก้ไข : การตั้งค่า รหัส  " + rs.ID.String() + " " + rs.ConfigCode + " " + rs.ConfigValue,
		})

		configsResponse = append(configsResponse, model.ConfigResponse{
			ID:           rs.ID.String(),
			ConfigCode:   rs.ConfigCode,
			ConfigDetail: rs.ConfigDetail,
			ConfigValue:  rs.ConfigValue,
		})
	}
	return configsResponse
}

// List implements service.ConfigService
func (service *configServiceImpl) List(ctx context.Context, offset int, limit int, order string, searchRequest model.ConfigRequest) (responses []model.ConfigResponse) {
	searchEntity := entity.Config{
		ConfigCode: searchRequest.ConfigCode,
	}
	results := service.ConfigRepository.List(ctx, offset, limit, order, searchEntity)
	if len(results) == 0 {
		return []model.ConfigResponse{}
	}

	for _, rs := range results {
		responses = append(responses, model.ConfigResponse{
			ID:           rs.ID.String(),
			ConfigCode:   rs.ConfigCode,
			ConfigDetail: rs.ConfigDetail,
			ConfigValue:  rs.ConfigValue,
		})
	}
	return responses
}
