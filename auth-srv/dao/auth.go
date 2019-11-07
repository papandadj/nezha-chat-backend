package dao

import (
	"errors"

	"github.com/patrickmn/go-cache"
)

// AuthCheck check  for expire and token is equal in memory .
func (d *Dao) AuthCheck(token string) (success bool, err error) {
	return
}

// AuthSaveToken .
func (d *Dao) AuthSaveToken(username, token string) (err error) {
	d.Mem.Set(username, token, cache.DefaultExpiration)
	return
}

// AuthDelToken .
func (d *Dao) AuthDelToken(token string) (err error) {
	return
}

// AuthGetToken get token .
func (d *Dao) AuthGetToken(username, id string) (token string, ok bool, err error) {
	tokenI, found := d.Mem.Get(username)
	if !found {
		return
	}

	token, ok = tokenI.(string)
	if !ok {
		err = errors.New("memory cache sava format error")
		return
	}
	return
}
