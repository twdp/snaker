package snaker

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

//任务实体类
type Task struct {
	BaseIdModel

	Version      int                                    //版本
	OrderId      int64  `orm:"index"`                   //流程实例ID
	TaskName     string `orm:"size(100) notnull index"` //任务名称
	DisplayName  string `orm:"size(200) notnull"`       //任务显示名称
	PerformType  PERFORM_ORDER                          //任务参与方式
	TaskType     TASK_ORDER                             //任务类型
	Operator     string `orm:"size(36)"`                //任务处理者ID
	FinishTime   time.Time        `orm:"null"`                      //任务完成时间
	ExpireTime   time.Time     `orm:"null"`                         //期望任务完成时间
	RemindTime   time.Time     `orm:"null"`                         //提醒时间
	Action       string     `orm:"size(200)"`           //任务关联的Action(WEB为表单URL)
	ParentTaskId int64      `orm:"index"`               //父任务ID
	Variable     string     `orm:"size(2000)"`          //任务附属变量(json存储)
	Model        *TaskModel `orm:"-"`                   //Model对象

	BaseTimeModel
}

func init() {
	orm.RegisterModelWithPrefix("snaker_", new(Task))
}

//根据ID得到任务
func (p *Task) GetTaskById(id int64) bool {
	p.Id = id

	err := orm.NewOrm().Read(p)
	if err != nil {
		logs.Error("fail to gettaskById. id: %d, err: %v", id, err)
		return false
	}
	return true
}

//得到活动任务
func (p *Task) GetActiveTasks() []*Task {
	tasks := make([]*Task, 0)
	//err := orm.Find(&tasks, p)
	//PanicIf(err, "fail to GetActiveTasks")
	return tasks
}

//根据OrderID得到活动任务
func (p *Task) GetActiveTasksByOrderId(orderId int64) []*Task {
	p.OrderId = orderId
	tasks := make([]*Task, 0)
	_, err := orm.NewOrm().QueryTable(p).Filter("OrderId", orderId).All(&tasks)
	if err != nil {
		logs.Error("fail to GetActiveTasksByOrderId. orderId: %d, cause: %v", orderId, err)
	}
	return tasks
}

//得到任务角色
func (p *Task) GetTaskActors() []*TaskActor {
	taskActors := make([]*TaskActor, 0)
	taskActor := &TaskActor{
		TaskId: p.Id,
	}

	_, err := orm.NewOrm().QueryTable(taskActor).Filter("TaskId", p.Id).All(&taskActors)
	if err != nil {
		logs.Error("fail to GetTaskActors. taskId: %d, cause: %v", p.Id, err)
	}


	return taskActors
}

//得到下一个ANY类型的任务
func GetNextAnyActiveTasks(parentTaskId int64) []*Task {
	task := &Task{
		ParentTaskId: parentTaskId,
	}
	tasks := make([]*Task, 0)

	_, err := orm.NewOrm().QueryTable(task).Filter("ParentTaskId", parentTaskId).All(&tasks)
	if err != nil {
		logs.Error("fail to GetNextAnyActiveTasks. parentTaskId: %d, cause: %v", parentTaskId, err)
	}

	return tasks
}

//得到下一个ALL类型的任务
func GetNextAllActiveTasks(orderId int64, taskName string, parentTaskId int64) []*Task {
	historyTask := &HistoryTask{
		OrderId:      orderId,
		TaskName:     taskName,
		ParentTaskId: parentTaskId,
	}
	tasks := make([]*Task, 0)

	historyTasks := make([]*HistoryTask, 0)

	_, err := orm.NewOrm().QueryTable(historyTask).Filter("OrderId", orderId, "TaskName", taskName, "ParentTaskId", parentTaskId).All(&historyTasks)
	if err != nil {
		logs.Error("fail to GetNextAllActiveTasks, orderId: %d, taskName: %s, parentTaskId: %d, cause: %v", orderId, taskName, parentTaskId, err)
		return tasks
	}

	ids := make([]int64, 0)
	for _, h := range historyTasks {
		ids = append(ids, h.Id)
	}
	_, err = orm.NewOrm().QueryTable(&Task{}).Filter("ParentTaskId__in", ids).All(&tasks)
	if err != nil {
		logs.Error("fail to GetNextAllActiveTasks two. ids: %v, err: %v", ids, err)
	}

	return tasks
}

//得到活动的任务（通过SQL）
func GetActiveTasksSQL(querystring string, args ...interface{}) ([]*Task, error) {
	tasks := make([]*Task, 0)
	_, err := orm.NewOrm().Raw(querystring, args).QueryRows(tasks)
	return tasks, err
}
