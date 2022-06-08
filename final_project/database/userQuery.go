package database

import (
	"final_project/models"
	"fmt"
)

func (u *DB) CreateUser(user models.User) (models.User, error) {
	err := u.db.Create(&user).Error
	return user, err
}

func (u *DB) FindUser(where map[string]interface{}) (models.User, error) {
	user := models.User{}
	fmt.Println(where)
	err := u.db.Where(where).First(&user).Error
	fmt.Println(user)
	return user, err
}

func (u *DB) UpdateUser(user models.User) (models.User, error) {
	err := u.db.Save(&user).Error
	return user, err
}

func (u *DB) DeleteUser(id interface{}) error {
	err := u.db.Delete(&models.User{}, id).Error
	return err
}
