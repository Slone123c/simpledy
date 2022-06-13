package repository

import (
	"fmt"
	"simpledy/model"
)

var RelationInsert = model.Relation{}

func IfRelationActionYes(userId int64, touser_id int64) (model.Relation, int) {
	var result model.Relation
	res := db.Model(model.Relation{}).Where("user_id = ? AND to_user_id = ?", userId, touser_id).Scan(&result)
	return result, int(res.RowsAffected)
}

//新生成关注信息储存进到Realtion表中
func CreateRelationInformation(newRelationInformation model.Relation) int {
	res := db.Create(&newRelationInformation)
	fmt.Println("新关注信息已生成成功！")
	return int(res.RowsAffected)
}

//从Relation表中删除该条点赞数据
func DeleteRelationInformation(userId int64, touser_id int64) bool {
	res := db.Where("user_id = ? AND to_user_id = ?", userId, touser_id).Unscoped().Delete(&RelationInsert)
	affected := res.RowsAffected
	if affected > 0 {
		fmt.Println("数据成功删除")
		return true
	} else {
		fmt.Println("数据删除失败")
		return false
	}
}

//更改数据库状态
//根据ID找到对应被关注/被取关用户,粉丝数量加/减1
//根据ID找到对应操作用户,关注数量加/减1
func UpdateRelationCount(user_id int64, to_user_id int64, action_type int) (c bool) {
	c = false
	var result1 model.UserInformation
	var result2 model.UserInformation
	//分别获取关注者和被关注者的信息
	db.Model(&model.UserInformation{}).Where("id = ?", user_id).Scan(&result1)
	db.Model(&model.UserInformation{}).Where("id = ?", to_user_id).Scan(&result2)

	//预设结果
	//关注
	result11 := result1.FollowCount + 1
	result21 := result2.FollowerCount + 1
	fmt.Println(result11, result21)
	//取关
	result12 := result1.FollowCount - 1
	result22 := result2.FollowerCount - 1
	fmt.Println(result12, result22)
	//更新UserInformation的Follower和Follow的Count
	if action_type == 1 {
		db.Model(&model.UserInformation{}).Where("id = ?", user_id).Update("follow_count", result11)
		db.Model(&model.UserInformation{}).Where("id = ?", user_id).Update("follower_count", result12)
		c = true
	} else {
		db.Model(&model.UserInformation{}).Where("id = ?", to_user_id).Update("follow_count", result21)
		db.Model(&model.UserInformation{}).Where("id = ?", to_user_id).Update("follower_count", result22)
		c = true
	}
	return c
}

//根据UserId找到所有的touserId
func FindFollowByUserId(userId int) ([]model.Relation, int) {
	var follows []model.Relation
	res := db.Debug().Model(model.Relation{}).Where("user_id = ?", userId).Find(&follows)
	return follows, int(res.RowsAffected)
}

//根据关注者的ID找到关注者的信息
func FindFollowInfoByUserId(follows []model.Relation) []model.UserInformation {
	followInfo := make([]model.UserInformation, len(follows))
	for i := 0; i < len(follows); i++ {
		db.Debug().Model(model.UserInformation{}).Where("id = ?", follows[i].ToUserId).Find(&followInfo[i])
		fmt.Println(followInfo)
	}
	return followInfo
}

//根据UserId找到所有的粉丝ID
func FindFollowerByUserId(userId int) ([]model.Relation, int) {
	var followers []model.Relation
	res := db.Debug().Model(model.Relation{}).Where("to_user_id = ?", userId).Find(&followers)
	return followers, int(res.RowsAffected)
}

//根据粉丝的ID找到所有粉丝信息
func FindFollowerInfoByUserId(followers []model.Relation) []model.UserInformation {
	followerInfo := make([]model.UserInformation, len(followers))
	for i := 0; i < len(followers); i++ {
		db.Debug().Model(model.UserInformation{}).Where("id = ?", followers[i].UserId).Find(&followerInfo[i])
		fmt.Println(followerInfo)
	}
	return followerInfo
}

func GenerateRelationFollowerListResponse(statusCode int64, statusMsg string,
	followers []model.UserInformation, userId int64) model.RelationFollowerListResponseAll {
	followerList := make([]model.RelationFollowerListResponseUserList, len(followers))
	var exist int
	for i := 0; i < len(followers); i++ {
		followerList[i].ID = followers[i].Id
		followerList[i].Name = followers[i].Name
		followerList[i].FollowCount = followers[i].FollowCount
		followerList[i].FollowerCount = followers[i].FollowerCount
		_, exist = IfRelationActionYes(userId, followers[i].Id)
		if exist > 0 {
			followerList[i].IsFollow = true
		} else {
			followerList[i].IsFollow = false
		}
	}
	return model.RelationFollowerListResponseAll{
		StatusCode: int32(statusCode),
		StatusMsg:  statusMsg,
		UserList:   followerList,
	}
}
