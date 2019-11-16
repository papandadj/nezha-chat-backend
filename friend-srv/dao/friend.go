package dao

// Post .
func (d *Dao) Post(user1, user2 string) (friend *ModelFriend, err error) {
	friend = new(ModelFriend)
	friend.UserID1 = user1
	friend.UserID2 = user2
	err = d.DB.Create(friend).Error
	return
}

//DeleteByUserID .
func (d *Dao) DeleteByUserID(user1, user2 string) (err error) {
	friend := new(ModelFriend)
	err = d.DB.Where("user_id1 = ? and user_id2 = ?", user1, user2).Delete(&friend).Error
	return
}

//GetList .
func (d *Dao) GetList(tokenID string) (friends []*ModelFriend, err error) {
	err = d.DB.Where("user_id1 = ? or user_id2 = ?", tokenID, tokenID).Find(&friends).Error
	return
}

//checkIsFriend .
func (d *Dao) checkIsFriend(tokenID, userID string) (result bool, err error) {
	return
}
