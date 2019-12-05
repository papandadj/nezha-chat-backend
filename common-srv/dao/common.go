package dao

// UserImgGetList .
func (d *Dao) UserImgGetList() (userImages []ModelUserImage, err error) {
	err = d.DB.Find(&userImages).Error
	return
}
