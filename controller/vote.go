package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PostVoteController(c *gin.Context) {
	//1.参数校验
	p := new(models.VoteData)
	if err := c.ShouldBindJSON(&p); err != nil {
		//类型断言
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}

	//需要登录

	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	//赞成票是1，反对票是-1
	logic.VoteForPost(userID, p)
	ResponseSuccess(c, nil)
}
