package snaker

import (
	"github.com/astaxie/beego/orm"
	"github.com/siddontang/go/log"
	"time"
)

//抄送实例表
type CCOrder struct {
	BaseIdModel

	OrderId    int64  `orm:"index"`    //流程实例ID
	ActorId    string `orm:"size(36)"` //操作者ID
	Creator    string `orm:"size(36)"` //流程实例创建者ID
	FinishTime time.Time               //流程实例完成时间
	State      FLOW_STATUS                    //流程实例状态

	BaseTimeModel
}

func init() {
	orm.RegisterModelWithPrefix("snaker_", new(CCOrder))
}

func GetCCOrder(orderId int64, actorIds ...string) []*CCOrder {
	ccorders := make([]*CCOrder, 0)
	_, err := orm.NewOrm().QueryTable(&CCOrder{}).Filter("OrderId", orderId).Filter("ActorId__in", actorIds).All(&ccorders)
	if err != nil {
		log.Errorf("get ccorder fail. orderId: %d, actorIds: %v", orderId, actorIds)
	}
	return ccorders
}
