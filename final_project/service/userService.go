package service

import (
	"final_project/database"
	"final_project/errorhandler"
	"final_project/helpers"
	"final_project/models"
	"fmt"
)

type UserService struct {
	userQuery database.UserQuery
}

// var db = gorm.DB{}
// var query = database.NewUserQuery(&db)

func NewUserService(query database.UserQuery) *UserService {
	return &UserService{
		userQuery: query,
	}
}

func (u *UserService) CreateUser(user models.User) (models.User, error) {
	encPass, err := helpers.PasswordEncryption(user.Password)
	errorhandler.Check(err)
	user.Password = encPass
	res, err := u.userQuery.CreateUser(user)
	return res, err
}

func (u *UserService) UserLogin(email, password string) (string, error) {
	where := make(map[string]interface{})
	where["email"] = email
	user, err := u.userQuery.FindUser(where)
	if err != nil {
		return "User Not Found", err
	}
	fmt.Println(user)
	err = helpers.ComparePassword(user.Password, password)
	if err != nil {
		fmt.Println("alhfoeuwhofewubg")
		return "password salah", err
	}

	token, err := helpers.GenerateToken(user.ID, user.Email, user.UserName)
	if err != nil {
		fmt.Println("eufhewuofhoewgfb")
		return "gagal generate token", err
	}

	// helpers.SaveTokenToDB(token, email)

	return token, nil
}

func (u *UserService) UpdateUser(user models.User, id interface{}) (models.User, error) {
	where := make(map[string]interface{})
	where["id"] = id

	prevUserData, err := u.userQuery.FindUser(where)
	if err != nil {
		return prevUserData, err
	}

	user.ID = prevUserData.ID
	user.CreatedAt = prevUserData.CreatedAt
	user.Password = prevUserData.Password

	updatedUser, err := u.userQuery.UpdateUser(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (u *UserService) Deleteuser(id interface{}) (string, error) {
	err := u.userQuery.DeleteUser(id)
	if err != nil {
		return err.Error(), err
	}

	return "Your account has successfully deleted", nil
}
