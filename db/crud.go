package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"tt/model"
)

func CreateUser(user *model.User) (err error) {
	return Conn.Create(user).Error
}

func UpdateUser(user *model.User) (err error) {
	if err = Conn.First(&model.User{}, user.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("id: %d doesn't exist", user.ID)
		}
		return
	}

	var updates map[string]interface{}
	var data []byte

	if data, err = json.Marshal(user); err != nil {
		return
	}
	if err = json.Unmarshal(data, &updates); err != nil {
		return
	}

	return Conn.Model(user).Updates(updates).Error
}

func DeleteUser(id int64) (err error) {
	if err = Conn.First(&model.User{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("id: %d doesn't exist", id)
		}
		return
	}
	return Conn.Delete(&model.User{}, id).Error
}

func SearchUsersOrdered(emailSubstr, phoneSubstr, orderBy string) (users []model.User, err error) {
	err = Conn.Order(orderBy).Where("email LIKE ?", emailSubstr).Where("phone LIKE ?", phoneSubstr).Find(&users).Error
	return
}
