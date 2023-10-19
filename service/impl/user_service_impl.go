package impl

import (
	"context"
	"strconv"
	"warrantyapi/common"
	"warrantyapi/constant"
	"warrantyapi/entity"
	"warrantyapi/exception"
	"warrantyapi/model"
	"warrantyapi/repository"
	"warrantyapi/service"

	"golang.org/x/crypto/bcrypt"
)

func NewUserServiceImpl(userRepository *repository.UserRepository, logRepository *repository.LogRepository) service.UserService {
	return &userServiceImpl{
		UserRepository: *userRepository,
		LogRepository:  *logRepository,
	}
}

type userServiceImpl struct {
	repository.UserRepository
	repository.LogRepository
}

// Create implements repository.UserRepository
func (service *userServiceImpl) Create(ctx context.Context, userInput model.User, createdBy string) model.UserResponse {
	common.Validate(userInput)

	_, err := service.UserRepository.FindDuplicateUserName(ctx, userInput.UserName)
	if err != nil {
		panic(exception.DuplicateError{
			Message: constant.MESSAGE_USER_DUPLICATE,
		})
	}

	hashPass, _ := service.HashPassword(userInput.UserName + userInput.Passwords)
	user := entity.User{
		CreatedBy:  createdBy,
		UserName:   userInput.UserName,
		Passwords:  hashPass,
		UserType:   userInput.UserType,
		UserStatus: userInput.UserStatus,
		FirstName:  userInput.FirstName,
		LastName:   userInput.LastName,
	}
	service.UserRepository.Insert(ctx, user)
	service.LogRepository.Insert(ctx, entity.Log{
		CreatedBy: createdBy,
		Module:    "user",
		Detail:    "สร้าง : ผู้ใช้งาน รหัส  " + strconv.FormatUint(uint64(user.ID), 10) + " ชื่อ " + user.UserName,
	})
	return model.UserResponse{
		ID:         user.ID,
		UserType:   user.UserType,
		UserStatus: user.UserStatus,
		UserName:   user.UserName,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
	}
}

// Update implements repository.UserRepository
func (service *userServiceImpl) Update(ctx context.Context, userInput model.User, updatedBy string) model.UserResponse {
	common.Validate(userInput)

	user := entity.User{
		ID:         userInput.ID,
		UserType:   userInput.UserType,
		UserStatus: userInput.UserStatus,
		FirstName:  userInput.FirstName,
		LastName:   userInput.LastName,
	}
	service.UserRepository.Update(ctx, user)
	service.LogRepository.Insert(ctx, entity.Log{
		CreatedBy: updatedBy,
		Module:    "user",
		Detail:    "แก้ไข : ผู้ใช้งาน รหัส  " + strconv.FormatUint(uint64(user.ID), 10) + " ชื่อ " + user.UserName,
	})
	return model.UserResponse{
		ID:         user.ID,
		UserType:   user.UserType,
		UserStatus: user.UserStatus,
		UserName:   user.UserName,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
	}
}

// Delete implements repository.UserRepository
func (service *userServiceImpl) Delete(ctx context.Context, userId string, createdBy string) {
	user, err := service.UserRepository.FindById(ctx, userId)
	if err != nil {
		panic(exception.NotFoundError{
			Message: constant.MESSAGE_USER_NOT_FOUND,
		})
	}
	service.UserRepository.Delete(ctx, user)
	service.LogRepository.Insert(ctx, entity.Log{
		CreatedBy: createdBy,
		Module:    "user",
		Detail:    "ลบ : ผู้ใช้งาน รหัส  " + strconv.FormatUint(uint64(user.ID), 10) + " ชื่อ " + user.UserName,
	})
}

// FindById implements repository.UserRepository
func (service *userServiceImpl) FindById(ctx context.Context, userId string) model.UserResponse {
	user, err := service.UserRepository.FindById(ctx, userId)
	if err != nil {
		panic(exception.NotFoundError{
			Message: constant.MESSAGE_USER_NOT_FOUND,
		})
	}
	//test send google chat
	return model.UserResponse{
		ID:         user.ID,
		UserType:   user.UserType,
		UserStatus: user.UserStatus,
		UserName:   user.UserName,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
	}
}

// FindAll implements repository.UserRepository
func (service *userServiceImpl) FindAll(ctx context.Context, limit int, offset int, orderBy string, sort string) (users []model.UserResponse) {
	responses := service.UserRepository.FindAll(ctx, limit, offset, orderBy, sort)
	for _, user := range responses {
		// log.Debug().Str("FirstName", user.FirstName).Send()
		users =
			append(users, model.UserResponse{
				ID:         user.ID,
				UserType:   user.UserType,
				UserStatus: user.UserStatus,
				UserName:   user.UserName,
				FirstName:  user.FirstName,
				LastName:   user.LastName,
			})
	}
	return users
}

// FindByUserName implements repository.UserRepository
func (service *userServiceImpl) FindByUserName(ctx context.Context, userName string) model.UserResponse {
	user, err := service.UserRepository.FindByUserName(ctx, userName)
	if err != nil {
		panic(exception.NotFoundError{
			Message: constant.MESSAGE_USER_NOT_FOUND,
		})
	}
	return model.UserResponse{
		ID:         user.ID,
		UserType:   user.UserType,
		UserStatus: user.UserStatus,
		UserName:   user.UserName,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
	}
}

func (service *userServiceImpl) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}
