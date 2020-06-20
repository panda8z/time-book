package model

import "errors"

// Auth is the basic auth data model in database
type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth from database for given username and password
func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}

// AddUser used by register func to create new user
func AddUser(username, password string) error {
	if CheckAuth(username, password) {
		return errors.New("User already have")
	}
	// insert into tb_auth("username","password") value("xxxx","123456")
	user := &Auth{
		Username: username,
		Password: password,
	}
	// INSERT INTO `tb_auth` (`username`,`password`) VALUES ('panda1','123456')
	db.Create(user)
	if user.ID != 0 {
		return nil
	} else {
		return errors.New("Unkown error")
	}
}
