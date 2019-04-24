package snaker

import (
	"github.com/astaxie/beego/orm"
	"github.com/siddontang/go/log"
	"time"
)

//任务实体类
type HistoryTask struct {
	BaseIdModel

	OrderId      int64  `orm:"index"`                   //流程实例ID
	TaskName     string `orm:"size(100);index"` //任务名称
	DisplayName  string `orm:"size(200)"`       //任务显示名称
	PerformType  PERFORM_ORDER                          //任务参与方式
	TaskType     TASK_ORDER                             //任务类型
	Operator     string `orm:"size(36)"`                //任务处理者ID
	FinishTime   time.Time    `orm:"null"`                          //任务完成时间
	ExpireTime   time.Time    `orm:"null"`                           //期望任务完成时间
	Action       string `orm:"size(200)"`               //任务关联的Action(WEB为表单URL)
	ParentTaskId int64  `orm:"index"`                   //父任务ID
	Variable     string `orm:"size(2000)"`              //任务附属变量(json存储)
	TaskState    FLOW_STATUS                            //任务状态

	BaseTimeModel
}

func init() {
	orm.RegisterModelWithPrefix("snaker_", new(HistoryTask))
}

//根据ID得到HistoryTask
func (p *HistoryTask) GetHistoryTaskById(id int64) bool {
	p.Id = id
	err := orm.NewOrm().Read(p)
	if err != nil {
		log.Errorf("fail to get history task by id. id: %d,err： %v", id, err)
		return false
	}
	return true
}

//通过HistoryTask生成Task
func (p *HistoryTask) Undo() *Task {
	task := &Task{
		BaseIdModel: BaseIdModel{
			Id:           p.Id,
		},
		TaskName:     p.TaskName,
		DisplayName:  p.DisplayName,
		TaskType:     p.TaskType,
		ExpireTime:   p.ExpireTime,
		Action:       p.Action,
		ParentTaskId: p.ParentTaskId,
		Variable:     p.Variable,
		PerformType:  p.PerformType,
		Operator:     p.Operator,
	}
	return task
}
