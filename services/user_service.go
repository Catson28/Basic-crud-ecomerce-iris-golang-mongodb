package services

import (
	"fmt"
	//"sync"
	//"util"

	"tentativa/datamodels"
	"tentativa/datamodels/request"

	"tentativa/repo"
	"tentativa/util"
)

// UserService is responsible for User CRUD operations,
// however, for the sake of the example we only implement the Read one.
type UserService interface {
	Create(user datamodels.User) (response datamodels.Response)
	GetSinger(singer request.SignRequest) (user datamodels.User, booler bool)
}

type userService struct {
	repo repo.UsersRepository
}

var userRepo = repo.NewUsersRepository()

func NewUserService() UserService {
	return &userService{
		repo: userRepo,
	}
}

func (u *userService) Create(user datamodels.User) (response datamodels.Response) {
	err := u.repo.Save(user)
	if err != nil {
		response.Code = 30001
		response.Msg = fmt.Sprintf("保存数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"

	return
}

func (u *userService) GetSinger(singer request.SignRequest) (user datamodels.User, booler bool) {
	user, err := u.repo.GetSignerByName(singer.Username)
	if err != nil {
		return user, true
	}

	// we compare the user input and the stored hashed password.
	ok := util.ValidatePassword(singer.Password, user.Password)
	if ok {
		//panic(ok)
		return user, true
	}

	return user, false
}
