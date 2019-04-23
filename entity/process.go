package entity

import "github.com/astaxie/beego/orm"

const (
	ProcessInit = iota

)

// 流程定义实体类
type Process struct {

	// 主键
	Id int64

	// 版本
	Version int8

	// 流程定义名称，根据此字段启动流程
	Name string

	// 页面上展示的名称
	DisplayName string

	// 流程定义类型(预留字段)
	Type string

	// 当前状态
	Status int8

	// 流程定义内容
	Content string

	// 创建人
	Creator string

	CreatedAt orm.TimeField `orm:"auto_now_add"`

	UpdatedAt orm.TimeField `orm:"auto_now"`
}

