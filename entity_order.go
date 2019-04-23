package snaker

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

//流程工作单实体类（一般称为流程实例）
type Order struct {
	BaseIdModel

	Version        int                           //版本
	ProcessId      int64  `orm:"index"`          //流程定义ID
	Creator        string `orm:"size(36)"`       //流程实例创建者ID
	ParentId       int64  `orm:"index"`          //流程实例为子流程时，该字段标识父流程实例ID
	ParentNodeName string `orm:"size(100)"`      //流程实例为子流程时，该字段标识父流程哪个节点模型启动的子流程
	ExpireTime     time.Time   `orm:"null"`                  //流程实例期望完成时间
	LastUpdator    string `orm:"size(36)"`       //流程实例上一次更新人员ID
	Priority       int                           //流程实例优先级
	OrderNo        int64  `orm:"index"`          //流程实例编号
	Variable       string `orm:"size(3000)"` //流程实例附属变量

	BaseTimeModel
}

func init() {
	orm.RegisterModelWithPrefix("snaker_", new(Order))
}

//根据ID得到Order
func (p *Order) GetOrderById(id int64) bool {
	p.Id = id
	err := orm.NewOrm().Read(p)
	if err != nil {
		logs.Error("fail to GetOrderById, id: %d, err: %v", id, err)
		return false
	}
	return true
}

//得到活动的Order（通过SQL）
func GetActiveOrdersSQL(querystring string, args ...interface{}) []*Order {
	orders := make([]*Order, 0)
	_, err := orm.NewOrm().Raw(querystring, args).QueryRows(orders)
	if err != nil {
		logs.Error("fail to GetActiveOrdersSQL. querystring: %s, args: %v, err: %v", querystring, args, err)
	}
	return orders
}
