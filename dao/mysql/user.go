package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

const secret = "liwenzhou.com"

// 把每一步数据库操作封装成函数，
// 待logic层根据业务需求调用

// CheckUserExist	检查指定用户名的用户是否存在
func CheckUserExist(username string) error {
	sqlstr := "select count(user_id) from user where username=?"
	var count int
	if err := db.Get(&count, sqlstr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

// InsertUser	向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	//对密码进行加密
	user.Password = encryptPassword(user.Password)
	//执行sql语句入库
	sqlstr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlstr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
	//EncodeToString转换成16进制
}

func Login(user *models.User) error {
	oPassword := user.Password
	sqlstr := `select user_id,username,password from user where username=?`
	err := db.Get(user, sqlstr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return nil
}

// GetUserById 根据id获取用户信息
func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, username from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	return
}
