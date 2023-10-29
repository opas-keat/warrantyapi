package impl

import (
	"context"
	"warrantyapi/entity"
	"warrantyapi/exception"
	"warrantyapi/model"
	"warrantyapi/repository"
	"warrantyapi/service"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func NewAuthenServiceImpl(authenRepository *repository.AuthenRepository) service.AuthenService {
	return &authenServiceImpl{
		AuthenRepository: *authenRepository,
	}
}

type authenServiceImpl struct {
	repository.AuthenRepository
}

func (service *authenServiceImpl) Login(ctx context.Context, userName string, passWord string) entity.UserAuth {
	user, err := service.AuthenRepository.Login(ctx, userName)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	checkPassLogin := false
	checkPassLogin = service.CheckPasswordHash(userName+passWord, user.Passwords)
	log.Debug().Bool("checkPassLogin", checkPassLogin).Send()
	if !checkPassLogin {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	return user
	// role := "ROLE_USER"
	// if user.UserType == "1" {
	// 	role = "ROLE_ADMIN"
	// }
	// var userRoles []map[string]interface{}
	// userRoles = append(userRoles, map[string]interface{}{
	// 	"role": role,
	// })

	// tokenJwtResult := common.GenerateToken(user.UserName, userRoles,)
	// resultWithToken := map[string]interface{}{
	// 	"token":    tokenJwtResult,
	// 	"username": user.UserName,
	// 	"role":     userRoles,
	// }
	// // Create token
	// token := jwt.New(jwt.SigningMethodHS256)
	// // Set claims
	// claims := token.Claims.(jwt.MapClaims)
	// claims["name"] = user.UserName
	// claims["role"] = role
	// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// // Generate encoded token and send it as response.
	// t, err := token.SignedString([]byte("omsoft"))
	// // token, err := CreateToken(authD)
	// if err != nil {
	// 	println("login %s\n", err)
	// }

	// return resultWithToken
}

func (service *authenServiceImpl) LogOut(ctx context.Context) model.AuthResponse {
	return model.AuthResponse{}
}

func (service *authenServiceImpl) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
