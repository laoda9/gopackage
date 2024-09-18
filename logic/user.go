package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	_ "errors"
)

func SignUp(p *models.ParamSignUp) (err error) {
	//判断用户存不存在
	if err = mysql.CheckUserExist(p.UserName); err != nil {
		return err
	}
	//生成uid
	userID := snowflake.GenID()
	//构建一个user实例
	user := &models.User{
		UserID:   userID,
		Username: p.UserName,
		Password: p.Password,
	}
	//保存进数据库（保存进dao层）
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.UserName,
		Password: p.Password,
	}
	//传递的是指针,就能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	//生成JWT
	return jwt.GenToken(user.UserID, user.Username)
}
