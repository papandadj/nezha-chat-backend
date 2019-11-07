package dao

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

//UserLogin .
func (d *Dao) UserLogin(username, password string) (user *ModelUser, ok bool, err error) {
	user = new(ModelUser)
	db := d.DB.Where("username = ?", username).Where("password = ?", password)
	ok, err = d.DB.FindOne(db, user)
	return
}
