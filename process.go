package snaker

import (
	"errors"
	"fmt"
	"tianwei.pro/snaker/entity"
	"tianwei.pro/snaker/parse"
)

// 流程定义业务类
type IProcessService interface {

	// 检查流程定义对象
	Check(process *entity.Process, idOrName string) error

	// 保存流程定义
	SaveProcess(process *entity.Process) error

	// 更新流程定义的类型
	// tt 类型
	UpdateType(id int64, tt int8) error

	// 根据主键ID获取流程定义对象
	GetProcessById(id int64) *entity.Process

	// 根据key获取流程定义对象
	GetProcessByKey(key string) *entity.Process

	// 根据key\version获取流程定义对象
	GetProcessByKeyAndVersion(key string) *entity.Process

	// 部署流程定义
	// return:  id -> 主键
	Deploy(input string) (int64, error)

	// 根据流程定义xml的输入流解析为字节数组，保存至数据库中，并且put到缓存中
	DeployByCreator(input string, creator string)

	// 重新部署流程
	ReDeploy(id int64, input string) error

	// 卸载流程，更新状态
	UnDeploy(id int64) error

	// 删除关联关系
	CascadeRemove(id int64) error
}

type ProcessService struct {
	
}

func (p *ProcessService) Check(process *entity.Process, idOrName string) error {
	if nil == process {
		return errors.New(fmt.Sprintf("指定的流程定义[id/name=%s]不存在", idOrName))
	} else if process.Status == entity.ProcessInit {
		return errors.New(fmt.Sprintf("指定的流程定义[id/name=%s,version=%d]为非活动状态", idOrName, process.Version))
	}
}

func (p *ProcessService) SaveProcess(process *entity.Process) error {
	panic("implement me")
}

func (p *ProcessService) UpdateType(id int64, tt int8) error {
	panic("implement me")
}

func (p *ProcessService) GetProcessById(id int64) *entity.Process {
	panic("implement me")
}

func (p *ProcessService) GetProcessByKey(key string) *entity.Process {
	panic("implement me")
}

func (p *ProcessService) GetProcessByKeyAndVersion(key string) *entity.Process {
	panic("implement me")
}

func (p *ProcessService) Deploy(input string) (int64, error) {
	return p.DeployByCreator(input, "")
}

func (p *ProcessService) DeployByCreator(input string, creator string) (int64, error) {
	if model, err := parse.Parse(input); err != nil {
		return 0, err
	} else {
		fmt.Println(model)
		return 0,nil
	}
}

func (p *ProcessService) ReDeploy(id int64, input string) error {
	panic("implement me")
}

func (p *ProcessService) UnDeploy(id int64) error {
	panic("implement me")
}

func (p *ProcessService) CascadeRemove(id int64) error {
	panic("implement me")
}
 