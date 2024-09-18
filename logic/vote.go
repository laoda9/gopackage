package logic

import "bluebell/models"

/*
投票功能
用户投票
投一票就加432分	86400/200->200票可以给帖子续上一天
投票几种情况：
direction=1时：
1.之前没用投过票，现在投赞成票	//2.之前投过反对票，现在投赞成票
direction=0时：
1.之前投过赞成票	//2.之前投过反对票，现在取消投票
direction：-1时：
1.之前没投过票，现在投反对票	//2.之前投赞成票，现在投反对票

投票限制：每个帖子自一个星期内允许投票，超过一星期就不允许再投票

	1.到期之后redis中存储的票数保存到mysql中
	2.到期之后删除KeyPostVotedZSetPrefix
*/
func VoteForPost(userID int64, p *models.VoteData) (err error) {
	return VoteForPost(userID, p)
}
