package handler

// import "simpledy/model"

// func HandleFavoriteListPost(user_id string, token string) (model.FavoriteListResponseAll, error) {
// 	var statusCode = -1     //0-成功，其他值-失败
// 	var statusMsg = ""      //返回状态描述
// 	var video_id = 0        //视频唯一标识
// 	var user_id = 0         //用户id
// 	var name = ""           //用户名
// 	var follow_count = 0    //关注总数
// 	var follower_count = 0  //粉丝总数
// 	var is_follow = false   //true-已关注，false-未关注
// 	var play_url = ""       //视频播放地址
// 	var cover_url = ""      //视频封面地址
// 	var favorite_count = 0  //视频点赞总数
// 	var comment_count = 0   //视频评论总数
// 	var is_favorite = false //true-已点赞，false-未点赞
// 	var title = ""          //视频标题

// }

// func HandleRegisterPost(username string, password string) (model.UserLoginResponse, error) {
// 	// 检查用户是否重复
// 	var statusCode = -1
// 	var statusMsg = ""
// 	var userId = 0
// 	var token = ""
// 	var err error
// 	exist := repository.IfUserExistsByUsername(username)
// 	if exist {
// 		statusCode = -1
// 		err = errors.New("用户名已存在")
// 		statusMsg = err.Error()
// 	} else {
// 		// 设置响应消息
// 		statusCode = 1
// 		statusMsg = "用户注册成功"
// 		// 创建新用户
// 		user := model.User{
// 			Username: username,
// 			Password: password,
// 		}
// 		// 数据库插入新用户
// 		repository.InsertNewUser(user)
// 		newUser := repository.FindUserByUserName(username)
// 		userId = int(newUser.ID)
// 		token, _ = utils.CreateToken(newUser)
// 	}
// 	// 更新并返回响应
// 	resp := generateUserLoginResponse(int32(statusCode), statusMsg, userId, token)
// 	return resp, err
// }

// func generateUserLoginResponse(statusCode int32, statusMsg string, userId int, token string) model.UserLoginResponse {
// 	return model.UserLoginResponse{
// 		StatusCode: statusCode,
// 		StatusMsg:  statusMsg,
// 		UserId:     userId,
// 		Token:      token,
// 	}
// }

// func generateFavoriteListResponse(statusCode int32, statusMsg string,
// 	video_id int32, user_id int32, name string, follow_count int32,
// 	follower_count int32, is_follow bool, play_url string, cover_url string,
// 	favorite_count int32, comment_count int32, is_favorite bool,
// 	title string) model.FavoriteListResponseAll {
// 	return model.FavoriteListResponseAll{
// 		StatusCode: statusCode,
// 		StatusMsg:  statusMsg,
// 		VideoList: model.FavoriteListResponseVideoList{
// 			IsFavorite: is_favorite,
// 			PlayUrl:    "http://cftysi.lc/wdduk",
// 			author: {
// 				user_id:        31,
// 				name:           "委行六",
// 				follow_count:   29,
// 				follower_count: 62,
// 				is_follow:      false,
// 			},
// 			title:          "口只主外认真",
// 			Id:             72,
// 			CoverUrl:       "http://gjjidfo.pr/oqxj",
// 			favorite_count: 30,
// 			comment_count:  42,
// 		},
// 	}
// }

// func generateFavoriteListResponse(statusCode int32, statusMsg string,
// 	VideoList *[]model.FavoriteListResponseVideoList) model.FavoriteListResponseAll {
// 	return model.FavoriteListResponseAll {
// 		StatusCode: statusCode,
// 		StatusMsg:  &statusMsg,
// 		VideoList : {
//             is_favorite: false,
//             play_url: "http://cftysi.lc/wdduk",
//             author: {
//                 user_id: 31,
//                 name: "委行六",
//                 follow_count: 29,
//                 follower_count: 62,
//                 is_follow: false,
//             },
//         	title: "口只主外认真",
//             video_id: 72,
//             cover_url: "http://gjjidfo.pr/oqxj",
//             favorite_count: 30,
//             comment_count: 42,
// 		},
// 	}
// }
