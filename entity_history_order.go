package snaker

import (
	"github.com/astaxie/beego/orm"
	"github.com/siddontang/go/log"
	"time"
)

//历史流程实例实体类
type HistoryOrder struct {
	BaseIdModel

	ProcessId  int64  `orm:"index"` //流程定义ID
	Creator    string `orm:"varchar(36)"`   //流程实例创建者ID
	ParentId   int64  `orm:"index"`         //流程实例为子流程时，该字段标识父流程实例ID
	ExpireTime time.Time     `orm:"null"`                //流程实例期望完成时间
	Priority   int                          //流程实例优先级
	OrderNo    int64  `orm:"index"`         //流程实例编号
	Variable   string `orm:"size(2000)"`    //流程实例附属变量
	OrderState FLOW_STATUS                  //流程实例状态
	FinishTime time.Time    `orm:"null"`                //完成时间

	BaseTimeModel
}

func init() {
	orm.RegisterModelWithPrefix("snaker_", new(HistoryOrder))
}

//从Order对象获取数据构件HistoryOrder
func (p *HistoryOrder) DataByOrder(order *Order) {
	p.Id = order.Id
	p.ProcessId = order.ProcessId
	p.ExpireTime = order.ExpireTime
	p.Creator = order.Creator
	p.ParentId = order.ParentId
	p.Priority = order.Priority
	p.OrderNo = order.OrderNo
	p.Variable = order.Variable
}

//根据ID得到HistoryOrder
func (p *HistoryOrder) GetHistoryOrderById(id int64) bool {
	p.Id = id
	err := orm.NewOrm().Read(p)
	if err != nil {
		log.Errorf("fail to get history by id. id: %v, err: %v", id, err)
		return false
	}
	return true
}

//通过HistoryOrder生成Order
func (p *HistoryOrder) Undo() *Order {
	order := &Order{
		BaseIdModel: BaseIdModel{
			Id: p.Id,
		},
		ProcessId:      p.ProcessId,
		ExpireTime:     p.ExpireTime,
		Creator:        p.Creator,
		LastUpdator:    p.Creator,
		//LastUpdateTime: p.FinishTime,
		ParentId:       p.ParentId,
		Priority:       p.Priority,
		OrderNo:        p.OrderNo,
		Variable:       p.Variable,
		Version:        0,
	}
	return order
}
