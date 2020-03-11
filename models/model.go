
package models

import (
    "time"
)

//User struct dbo.UserInfo
type User struct{
	UserID int 
	UserNo string   //编号
	UserName string  //用户名 登录？
	NickName string
	password string
	plainCode string
	Email string
	Age int
	gender string
	imgID int		//头像ID
	levelID int		//等级
	roleID int		
	Introduction string  //简介
	isDeleted int
	dataOwner int
	createDateTime time.Time
	remark string
}

//Article struct
type Article struct{
	ArticleID int 
	Version int
	UserID int
	Title string   //标题
	Content string  //内容
	CoverID string //封面图
	ArticleType string
	IsPrivate string
	IsDeleted int
	DataOwner int
	CreateDateTime time.Time
	LastChagedBy int
	LastChangedDateTime time.Time
	Remark string
}
//Follower struct fans
type Follower struct{
	FollowerUserID int 
	UserID int
	BrowseCount int
	ReadCount int
	StarCount int
	LikeCount int
	IsDeleted int
	DataOwner int
	CreateDateTime time.Time
	LastChagedBy int
	LastChangedDateTime time.Time
	Remark string

	FollowerName string 
	FollowerIntroduction string

	UserName string
	UserIntroduction string
}

