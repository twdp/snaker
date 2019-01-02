package snaker

import (
	"container/list"
	"errors"
	"fmt"
	"strconv"
	"tianwei.pro/snaker/core"
	"tianwei.pro/snaker/entity"
	"tianwei.pro/snaker/model"
	"time"
)

const (
	ProcessKey = "snaker.process"
	InstanceKey = "snaker.instance"
	TaskKey = "snaker.task"
)

type SnakerEngine interface {

	// 获取process服务
	Process() *IProcessService

	// 获取查询服务
	Query() *IQueryService

	// 获取实例服务
	Instance() *IInstanceService

	// 获取任务服务
	Task() *ITaskService

	// 获取管理服务
	Manager() *IManagerService

	// 根据流程定义ID启动流程实例
	StartInstanceById(id int64) (*entity.Instance, error)

	// 根据流程定义id和操作人|flag启动流程实例
	StartInstanceByIdAndOperator(id int64, operator string) (*entity.Instance, error)

	// 根据流程定义id和操作人|flag和参数启动流程实例
	StartInstanceByIdAndOperatorAndArgs(id int64, operator string, args map[string]interface{}) (*entity.Instance, error)

	// 根据唯一key启动实例
	StartInstanceByKey(key string) (*entity.Instance, error)

	// 根据唯一key和version 启动实例
	StartInstanceByKeyAndVersion(key string, version int) (*entity.Instance, error)

	// 根据唯一key、version、操作信息启动实例
	StartInstanceByKeyAndVersionAndOperator(key string, version int, operator string) (*entity.Instance, error)

	// 根据唯一key、version、操作信息、执行参数启动实例
	StartInstanceByKeyAndVersionAndOperatorAndArgs(key string, version int, operator string, args map[string]interface{}) (*entity.Instance, error)

	// 根据父执行对象启动子流程实例
	StartInstanceByExecution(execution *core.Execution) (*entity.Instance, error)

	// 根据任务主键ID执行任务
	ExecuteTaskById(taskId int64) (*list.List, error)

	// 根据任务主键ID，操作人执行任务
	ExecuteTaskByIdAndOperator(taskId int64, operator string) (*list.List, error)

	// 根据任务主键ID，操作人，参数列表执行任务
	ExecuteTaskByIdAndOperatorAndArgs(taskId int64, operator string, args map[string]interface{}) (*list.List, error)

	// 根据任务主键ID，操作人，参数列表执行任务，并且根据nodeName跳转到任意节点
	// 1、nodeName为nil时，则跳转至上一步处理
	// 2、nodeName不为nil时，则任意跳转，即动态创建转移
	ExecuteAndJumpTask(taskId int64, operator string, args map[string]interface{}, nodeName string) (*list.List, error)

	// 根据流程实例ID，操作人，参数列表按照节点模型model创建新的自由任务
	CreateFreeTask(instanceId int64, operator string, args map[string]interface{}, model *model.TaskModel) (*list.List, error)
}

// snakerEngine实现类
type SnakerEngineImpl struct {

	// 流程定义业务类
	processService *IProcessService

	// 流程实例业务类
	instanceService *IInstanceService

	// 任务业务类
	taskService *ITaskService

	// 查询业务类
	queryService *IQueryService

	// 管理业务类
	managerService *IManagerService
}

func NewEngine() SnakerEngine {
	return &SnakerEngineImpl{
	}
}

//
func (s *SnakerEngineImpl) Process() *IProcessService {
	if nil == s.processService {
		panic("IProcessService没有实例")
	}
	return s.processService
}

func (s *SnakerEngineImpl) Query() *IQueryService {
	return nil
}

func (s *SnakerEngineImpl) Manager() *IManagerService {
	return nil
}

func (s *SnakerEngineImpl) Instance() *IInstanceService {
	if nil == s.instanceService {
		panic("IInstanceService没有实例")
	}
	return s.instanceService
}

func (s *SnakerEngineImpl) Task() *ITaskService {
	if nil == s.taskService {
		panic("ITaskService没有实例")
	}
	return s.taskService
}

func (s *SnakerEngineImpl) StartInstanceById(id int64) (*entity.Instance, error) {
	return s.StartInstanceByIdAndOperatorAndArgs(id, "", nil)
}

func (s *SnakerEngineImpl) StartInstanceByIdAndOperator(id int64, operator string) (*entity.Instance, error) {
	return s.StartInstanceByIdAndOperatorAndArgs(id, operator, nil)
}

func (s *SnakerEngineImpl) StartInstanceByIdAndOperatorAndArgs(id int64, operator string, args map[string]interface{}) (*entity.Instance, error) {
	if nil == args {
		args = make(map[string]interface{})
	}
	process := (*s.Process()).GetProcessById(id)
	(*s.Process()).Check(process, strconv.FormatInt(id, 10))
	return s.startProcess(process, operator, args)
}

func (s *SnakerEngineImpl) StartInstanceByKey(key string) (*entity.Instance, error) {
	return s.StartInstanceByKeyAndVersionAndOperatorAndArgs(key, 0, "", nil)
}

func (s *SnakerEngineImpl) StartInstanceByKeyAndVersion(key string, version int) (*entity.Instance, error) {
	return s.StartInstanceByKeyAndVersionAndOperatorAndArgs(key, version, "", nil)
}

func (s *SnakerEngineImpl) StartInstanceByKeyAndVersionAndOperator(key string, version int, operator string) (*entity.Instance, error) {
	return s.StartInstanceByKeyAndVersionAndOperatorAndArgs(key, version, operator, nil)
}

func (s *SnakerEngineImpl) StartInstanceByKeyAndVersionAndOperatorAndArgs(key string, version int, operator string, args map[string]interface{}) (*entity.Instance, error) {
	if nil == args {
		args = make(map[string]interface{})
	}
	process := (*s.Process()).GetProcessByKeyAndVersion(key)
	(*s.Process()).Check(process, key)
	return s.startProcess(process, operator, args)

}

func (s *SnakerEngineImpl) StartInstanceByExecution(execution *core.Execution) (*entity.Instance, error) {
	process := execution.Process
	if start, err := process.Model.GetStart(); err != nil {
		return nil, err
	} else if current, err := s.execute(process, execution.Operator, execution.Args, execution.ParentInstance.Id, execution.ParentNodeName) ; err != nil {
		return nil, err
	} else if err = start.Execute(current); err != nil {
		return nil, err
	} else {
		return current.Instance, nil
	}
}

func (s *SnakerEngineImpl) ExecuteTaskById(taskId int64) (*list.List, error) {
	return s.ExecuteTaskByIdAndOperator(taskId, "")
}

func (s *SnakerEngineImpl) ExecuteTaskByIdAndOperator(taskId int64, operator string) (*list.List, error) {
	return s.ExecuteTaskByIdAndOperatorAndArgs(taskId, operator, nil)
}

func (s *SnakerEngineImpl) ExecuteTaskByIdAndOperatorAndArgs(taskId int64, operator string, args map[string]interface{}) (*list.List, error) {
	if execution, err := s.executeById(taskId, operator, args); err != nil {
		return nil, err
	} else if nil == execution {
		return &list.List{}, nil
	} else {
		model := execution.Process.Model
		if model != nil {
			if nodeModel, err := model.GetNode(execution.Task.TaskName); err != nil {
				return nil, err
			} else {
				if err = nodeModel.Execute(execution); err != nil {
					return nil, err
				}
			}
		}
		return execution.Tasks, nil
	}
}

func (s *SnakerEngineImpl) ExecuteAndJumpTask(taskId int64, operator string, args map[string]interface{}, nodeName string) (*list.List, error) {
	if execution, err := s.executeById(taskId, operator, args); err != nil {
		return nil, err
	} else if nil == execution {
		return &list.List{}, nil
	} else {
		m :=  execution.Process.Model
		if nil == m {
			return nil, errors.New("当前任务未找到流程定义模型")
		}
		if nodeName == "" {
			 if newTask, err := (*s.Task()).RejectTask(m, execution.Task); err != nil {
				return nil, err
			} else {
				execution.AddTask(newTask)
			}
		} else if nodeModel, err := m.GetNode(nodeName); err != nil {
			return nil, err
		} else if nodeModel == nil {
			return nil, errors.New(fmt.Sprintf("根据节点名称[%s]无法找到节点模型", nodeName))
		} else {
			tm := &model.TransitionModel{
				Target: nodeModel,
				Enable: true,
			}
			if err = tm.Execute(execution); err != nil {
				return nil, err
			}
		}
		return execution.Tasks, nil
	}
}

func (s *SnakerEngineImpl) CreateFreeTask(instanceId int64, operator string, args map[string]interface{}, model *model.TaskModel) (*list.List, error) {
	if instance, err := (*s.Query()).GetInstance(instanceId); nil != err {
		return nil, err
	} else if instance == nil {
		return nil, errors.New(fmt.Sprintf("指定的流程实例[id=%d]已完成或不存在", instanceId))
	} else {
		instance.CreatedBy = &Logger{
			LoggerS: operator,
		}
		instance.UpdatedAt = time.Now()

		p := (*s.Process()).GetProcessById(instance.ProcessId)
		if err = (*s.Process()).Check(p, strconv.FormatInt(instance.ProcessId, 10)); err != nil {
			return nil, err
		}

		execution := &core.Execution{
			Engine: s,
			Process: p,
			Instance: instance,
			Args: args,
		}
		execution.Operator = operator
		return (*s.Task()).CreateTask(model, execution)
	}
}

func (s *SnakerEngineImpl) startProcess(process *entity.Process, operator string, args map[string]interface{}) (*entity.Instance, error) {
	if execution, err := s.execute(process, operator, args, 0, ""); nil != err {
		return nil, err
	} else {
		if process.Model != nil  {
			if start, err := process.Model.GetStart(); err != nil {
				return nil, err
			} else {
				if err = start.Execute(execution); err != nil {
					return nil, err
				}
			}
		}
		return execution.Instance, nil
	}
}

func (s *SnakerEngineImpl) executeById(taskId int64, operator string, args map[string]interface{}) (*core.Execution, error) {
	if args == nil {
		args = make(map[string]interface{})
	}

	if task, err := (*s.Task()).CompleteByIdAndOperatorAndArgs(taskId, operator, args); err != nil  {
		return  nil, err
	} else if instance, err := (*s.Query()).GetInstance(task.InstanceId); err != nil {
		return  nil, err
	} else {
		instance.CreatedBy = &Logger{
			LoggerS: operator,
		}
		instance.UpdatedAt = time.Now()
		if err = (*s.Instance()).UpdateInstance(instance); err != nil {
			return nil, err
		}
		// 协办任务完成不产生执行对象
		if !task.IsMajor() {
			return nil, nil
		}
		if instance.Variable != nil {
			for k, v := range instance.Variable {
				if _, ok := args[k]; ok {

				} else {
					args[k] = v
				}
			}
		}
		process := (*s.Process()).GetProcessById(instance.ProcessId)
		if err := (*s.Process()).Check(process, strconv.FormatInt(instance.ProcessId, 10)); err != nil {
			return nil, err
		} else {
			return &core.Execution{
				Engine: s,
				Process: process,
				Instance: instance,
				Args: args,
				Operator: operator,
				Task: task,
			}, nil
		}
	}

}

func (s *SnakerEngineImpl) execute(process *entity.Process, operator string, args map[string]interface{}, parentId int64, parentNodeName string) (*core.Execution, error) {
	if instance, err := (*s.Instance()).CreateInstanceUseParentInfo(process, operator, args, parentId, parentNodeName); nil != nil {
		return nil, err
	} else {
		fmt.Println(instance)
		current := &core.Execution{
			Engine: s,
			Process: process,
			Instance: instance,
			Args: args,
			Operator: operator,
			Tasks: &list.List{},
		}

		return current, nil
	}

}