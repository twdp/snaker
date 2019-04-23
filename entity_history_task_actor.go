package snaker

import "github.com/astaxie/beego/orm"

//历史任务参与者
type HistoryTaskActor struct {
	BaseIdModel

	TaskId  int64 `orm:"index"` //任务ID
	ActorId string `orm:"size(36)"`       //参与者ID

	BaseTimeModel
}

func init() {
	orm.RegisterModelWithPrefix("snaker_", new(HistoryTaskActor))
}