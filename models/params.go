package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	UserName   string `json:"userName" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"`   // 可以为空
	Page        int64  `json:"page" form:"page" example:"1"`       // 页码
	Size        int64  `json:"size" form:"size" example:"10"`      // 每页数据量
	Order       string `json:"order" form:"order" example:"score"` // 排序依据
}

type VoteData struct {
	//UserID从请求中获取当前的ID
	PostID    int64 `json:"postid" binding:"required"`                 //帖子id
	Direction int   `json:"direction" binding:"required" oneof=1 0 -1` //赞成票1反对票-1取消投票0,使用oneof
}
