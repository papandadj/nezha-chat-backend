package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// UserPost create user
func (d *Dao) UserPost(username, password string) (user *ModelUser, err error) {
	user = new(ModelUser)
	user.Username = username
	user.Password = password
	err = d.DB.Create(user).Error
	return
}

//UserGetByUsername get user by name
func (d *Dao) UserGetByUsername(username string) (user *ModelUser, ok bool, err error) {
	user = new(ModelUser)
	db := d.DB.Where("username = ?", username)
	ok, err = d.DB.FindOne(db, user)
	return
}

//UserCheckPassword .
func (d *Dao) UserCheckPassword(username, password string) (user *ModelUser, ok bool, err error) {
	user = new(ModelUser)
	db := d.DB.Where("username = ?", username).Where("password = ?", password)
	ok, err = d.DB.FindOne(db, user)
	return
}

//UserGetList .
func (d *Dao) UserGetList(name string, ids []string) (user []*ModelUser, err error) {
	var db *gorm.DB
	if name != "" {
		db = d.DB.Where("username LIKE ?", fmt.Sprintf("%s%s%s", "%", name, "%"))
	} else {
		db = d.DB.Where("id IN (?)", ids)
	}
	err = db.Find(&user).Error
	return
}

//UserGet .
func (d *Dao) UserGet(id string) (user *ModelUser, exist bool, err error) {
	user = new(ModelUser)
	db := d.DB.Where("id = ?", id)
	exist, err = d.DB.FindOne(db, user)

	return
}
