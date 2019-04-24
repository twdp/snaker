package snaker

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

//任务参与者
type TaskActor struct {
	BaseIdModel                             //主键ID
	TaskId  int64                           //任务ID
	ActorId string `orm:"size(36)"` //参与者ID

	BaseTimeModel
}

func init() {
	orm.RegisterModelWithPrefix("snaker_", new(TaskActor))
}

//通过任务ID，得到任务角色
func GetTaskActorsByTaskId(taskId int64) []*TaskActor {
	taskActor := &TaskActor{
		TaskId: taskId,
	}
	taskActors := make([]*TaskActor, 0)

	_, err := orm.NewOrm().QueryTable(taskActor).Filter("TaskId", taskId).All(&taskActors)
	if err != nil {
		logs.Error("fail to GetTaskActorsByTaskId. taskId: %d, cause: %v", taskId, err)
	}

	return taskActors
}
