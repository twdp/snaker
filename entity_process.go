package snaker

import (
	"github.com/astaxie/beego/orm"
	"github.com/siddontang/go/log"
)

//流程定义实体类
type Process struct {
	BaseIdModel

	Version        int                              //版本
	Name           string `orm:"size(100);index"`   //流程定义名称
	DisplayName    string `orm:"size(200)"`         //流程定义显示名称
	InstanceAction string `orm:"size(200)"`         //当前流程的实例Action,(Web为URL,一般为流程第一步的URL;APP需要自定义),该字段可以直接打开流程申请的表单
	State          FLOW_STATUS                      //状态
	Creator        string        `orm:"size(36)"`   //创建人
	Content        string        `orm:"type(text)"` //流程定义XML
	Model          *ProcessModel `orm:"-"`          //Model对象

	BaseTimeModel
}

func init() {
	orm.RegisterModelWithPrefix("snaker_", new(Process))
}

//根据ID得到Process
func (p *Process) GetProcessById(id int64) bool {
	p.Id = id
	return p.GetProcess()
}

//根据Process本身条件得到Process
func (p *Process) GetProcess() bool {
	err := orm.NewOrm().Read(p)
	if err != nil {
		log.Errorf("fail to get process. p: %v, err: %v", p, err)
		return false
	}
	return true
}

//根据Process本身条件得到Process
func (p *Process) GetProcessByNameAndVersion() bool {
	err := orm.NewOrm().Read(p, "Name", "Version")
	if err != nil {
		log.Errorf("fail to get process. p: %v, err: %v", p, err)
		return false
	}
	return true
}

//设定Model对象
func (p *Process) SetModel(model *ProcessModel) {
	p.Model = model
	p.Name = model.Name
	p.DisplayName = model.DisplayName
	p.InstanceAction = model.InstanceAction
}

//得到最新的Process
func GetLatestProcess(name string) *Process {
	process := &Process{
		Name: name,
	}
	processes := make([]*Process, 0)
	_, err := orm.NewOrm().QueryTable(process).Filter("Name", name).All(&processes)
	if err != nil {
		log.Errorf("get process lastest version failed. name: %s, err: %v", name, err)
	}
	if len(processes) > 0 {
		return processes[0]
	} else {
		return nil
	}
}
