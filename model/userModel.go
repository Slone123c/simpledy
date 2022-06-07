package model

type User struct {
	Id       int64  `gorm:"size:64;not null"`
	Username string `gorm:"size:32;not null"`
	Password string `gorm:"size:32;not null"`
}

type UserInformation struct {
	Id            int64  `gorm:"size:64;not null"`
	Name          string `gorm:"size:32;not null"` // 用户名称
	FollowCount   int64  `gorm:"size:64;not null"`
	FollowerCount int64  `gorm:"size:64;not null"`
	// 使得 UserId 成为外键并关联至 User 表
	User   User `gorm:"ForeignKey:UserId"`
	UserId int64
}

/*一共需要以下几个数据表
前四个必须表的详细字段见本文件最后
User表
UserInformation表
Video表
Comment表
favorite表 （可要可不要）
follow表 （可要可不要）
follower表 （可要可不要）
*/

/*各接口需求，不是很详细的详解：
	1.Feed(视频流接口):
	1）用户不登录也可以刷抖音，返回按投稿时间倒序的视频列表
	2）视频数由服务端控制，单次最多30个

	2.user/Register(用户注册):
	————某些需求已完成，说点补充的：
	1)用户注册成功后，账号密码要存储进User表，并自动生成随机ID
	2)注册成功后UserInformation表要自动添加该ID，并且自动初始化UserInformation表的其他字段

	3.user/Login(用户登录)：
	1)登录时查询User表是否存在该UserName，存在登录，不存在显示未注册

	4.User(用户信息)：
	两种情况会调用此接口
	1) 登录后，跳转到个人页面，调用此接口后将User表的id，name，follow_count，follower_count
	2) 到他人主页后，会加上多调用一个is_follow
       (结合接口11，比如到B用户主页后，会在UserInformation表中查询B用户的follower，
	    查不到返回False，否则为True)

	5.Publish/Action(视频发布接口)：
	1) =_=...有个疑问，视频封面怎么上传......
	2) 注意是请求Body参数(multipart/form-data)(所有接口除了这一个，都是请求Query参数)
	3) 发布的视频信息要添加到Video表中，并在UserInformation表中附加"publish_video"字段，
	   存储用户所发布的Video的ID，以关联Video表方便接口6调用

	6.Publish/List(用户视频列表)：
	1) 到个人主页后调用此接口，查询UserInformation表的"publish_video"字段，找到
	   所有个人发布的Video的ID

	7.favorite/Action(点赞/取消点赞)
  和8.favorite/List(喜欢的视频)
	1) 点赞后，Video表中该视频的ID，favorite_count要加1，
	   取消点赞后，Video表中该视频的ID，favorite_count要减1
	2) 点赞后，要能够在“我喜欢的视频”页面加上该视频，
	   取消点赞后，要能够在“我喜欢的视频”页面移除该视频
	————实现方法：在Video表中加一个“favorite_user"字段，存储喜欢该视频的用户的用户ID，
	             用户调用"Publish/List"接口后可以查询到喜欢的视频
		开个玩笑，上面的方法不可行，因为每一次调用都要查询所有的视频再判断该用户是否喜欢，
		时间复杂度太高，所以要换成下面这个：
		————真正的实现方法：在UserInformation表中加一个"favorite_video"字段，存储该用户喜欢
		的视频的视频ID，调用"favorite/List"接口后只需查询UserInformation表的favorite_video
		字段即可查询到用户所有喜欢的视频的视频ID
		————其实也可以新建一个Favorite表存储每一个用户喜欢的视频ID，但好像没啥必要
	(注意，这个方法与下面的Comment实现方法不同，与Relation实现方法相同)

	9.comment/Action(评论/删除评论)
  和10.comment/List(评论列表)
    ————需要注意的是，评论与点赞不同，喜欢的视频的列表是一个用户所有喜欢的视频在同一页面展示，
	    而评论列表是所有评论用户的评论在同一页面展示，所以数据库的Table要不同处理
	————因此，新建Comment列表，每一个Video对应一个Comment列表
    1) 用户评论后，在评论页面增加该条评论，并存储进Comment数据表
	2) 用户删除评论，要先判断这条评论是不是用户自己的，不是则删评失败，
	   删除成功后对应的Comment数据表要删除这条评论信息
	3) 调用"comment/List"接口后，要查询数据库的Video列表，获取该视频的ID，由该视频ID
	   到Comment表中找到对应评论

	11.relation/Action(关注/取消关注)
  和12.relation/follow/List(关注列表)
  和13.relation/follower/List(粉丝列表)
    1) A用户进入B用户的主页，点击关注后，UserInformation表中，B用户的follower_count加1，
	   A用户的follow_count加1;同理取消关注后，UserInformation表中B用户follower_count减1，
	   A用户的follow_count减1
	2) 在UserInformation表中新建"follow_user"字段,调用接口12后显示所有关注的用户
	3) 在UserInformation表中新建"follower_user"字段,调用接口13后显示所有关注的用户
	————或者新建Follow表和Follower表

*/

/*四个必须表的详细字段：
1) User表：
	username (主键，唯一标识)
	password
	userid

2) UserInfomation表：
	id(=User.userid)  ---用户ID  (主键，唯一标识)
	name(=User.username)   ---用户名
	follow_count     ---关注总数
	follower_count   ---粉丝总数
	publish_video(数组/切片)   ---上传视频的ID
	favorite_video(数组/切片)  ---喜欢视频的ID
	follow_user(数组/切片)     ---关注的用户的ID
	follower_user(数组/切片)   ---粉丝的ID
//因此：如何实现is_favorite和is_follow?
  查询UserInformation表中favorite_video是否含有该Video的ID
  查询UserInformation表中follow_user是否含有该用户的ID

3) Video表：
	id    ---视频ID  (主键，唯一标识)
	author(=UserInformation.id + UserInformation.name +
	        UserInformation.follow_count + UserInformation.follower_count)
		  (实际是存储视频上传者的ID，关联UserInformation表)
    play_url    ---视频播放地址
	cover_url   ---视频封面地址
	favorite_count  ---视频点赞总数
	comment_count   ---视频点赞总数
	title       ---视频标题

4) Comment表：
	(注意，是每条评论占据Comment表的一行，即Comment表中可以重复出现id一样的)
	(查询一个Video的所有评论只需把所有video_id一样的找出来即可)
	id                    ---评论ID  (主键，唯一标识)
	video_id(=Video.id)   ---视频ID
	user(=UserInformation.id + UserInformation.name +
	      UserInformation.follow_count + UserInformation.follower_count)
		(实际上是存储评论用户的ID，关联UserInformation表)
	content               ---评论内容
	create_data           ---评论创建时间

*/
