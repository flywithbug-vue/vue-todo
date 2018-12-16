package model

const (
	appCollection = "application"
)

type Application struct {
	Title        string //应用（组件）名称
	AppId        string
	Code         int
	Desc         string //项目描述
	CreateTime   int64  //创建时间
	ModifyTime   int64
	ModifyUserId string
}
