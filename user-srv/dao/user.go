package dao

// UserPost 创建用户
func (d *Dao) UserPost(username, password string) (user *ModelUser, err error) {
	user = new(ModelUser)
	user.Username = username
	user.Password = password

	return
}
