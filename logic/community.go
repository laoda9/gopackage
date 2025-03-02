package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

// GetCommunityList
func GetCommunityList() ([]*models.Community, error) {
	// 查数据库，查找到所有的数据并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
