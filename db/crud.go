package db

import (
	"encoding/json"
	"gorm.io/gorm"

	"tt/model"
)

func CreateUser(user *model.User) (err error) {
	return Conn.Create(user).Error
}

func UpdateUser(user *model.User) (err error) {
	var updates map[string]interface{}
	var data []byte
	if data, err = json.Marshal(user); err != nil {
		return
	}
	if err = json.Unmarshal(data, &updates); err != nil {
		return
	}
	return Conn.Transaction(func(tx *gorm.DB) (err error) { return tx.Model(user).Updates(updates).Error })
}

func DeleteUser(user *model.User) (err error) {
	return Conn.Delete(user).Error
}

func SearchUsersOrdered(emailSubstr, phoneSubstr, orderBy string) (users []model.User, err error) {
	err = Conn.Order(orderBy).Where("email LIKE ?", emailSubstr).Where("phone LIKE ?", phoneSubstr).Find(&users).Error
	return
}
