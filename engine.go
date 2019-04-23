package snaker

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"strconv"
	"tianwei.pro/snaker/entity"
)

const (
	ProcessKey  = "snaker.process"
	InstanceKey = "snaker.instance"
	TaskKey     = "snaker.task"
)

type Engine interface {
	// 获取process服务
	Process() IProcessService

	// 注册
	RegisterService(name string, instance interface{})

	// 根据name获取service
	GetByName(name string) interface{}

	// 根据name和type获取实例
	GetByNameAndType(name string, t interface{}) interface{}

	// 根据流程定义ID启动流程实例
	StartInstanceById(id int64) (*entity.Instance, error)

	// 根据流程定义id和操作人|flag启动流程实例
	StartInstanceByIdAndOperator(id int64, operator string) (*entity.Instance, error)

	// 根据流程定义id和操作人|flag和参数启动流程实例
	StartInstanceByIdAndOperatorAndArgs(id int64, operator string, args map[string]interface{}) (*entity.Instance, error)
}

// snakerEngine实现类
type SnakerEngineImpl struct {
	// 流程定义业务类
	processService IProcessService

	//// 流程实例业务类
	//instanceService *IInstanceService
	//
	//// 任务业务类
	//taskService *ITaskService
	//
	//// 查询业务类
	//queryService *IQueryService
	//
	//// 管理业务类
	//managerService *IManagerService

	di *Container
}

func NewEngine() Engine {
	return &SnakerEngineImpl{
		di:             New(),
		processService: NewProcessService(),
	}
}

// 注册
func (s *SnakerEngineImpl) RegisterService(name string, instance interface{}) {
	s.di.Provide(name, instance)
}

// 根据name获取service
func (s *SnakerEngineImpl) GetByName(name string) interface{} {
	return s.di.GetByName(name)
}

// 根据name和type获取实例
func (s *SnakerEngineImpl) GetByNameAndType(name string, t interface{}) interface{} {
	return s.di.GetByType(t)
}

//
func (s *SnakerEngineImpl) Process() IProcessService {
	if nil == s.processService {
		panic("IProcessService没有实例")
	}
	return s.processService
}

//
//func (s *SnakerEngineImpl) Query() *IQueryService {
//	return nil
//}
//
//func (s *SnakerEngineImpl) Manager() *IManagerService {
//	return nil
//}
//
//func (s *SnakerEngineImpl) Instance() *IInstanceService {
//	if nil == s.instanceService {
//		panic("IInstanceService没有实例")
//	}
//	return s.instanceService
//}

//func (s *SnakerEngineImpl) Task() *ITaskService {
//	if nil == s.taskService {
//		panic("ITaskService没有实例")
//	}
//	return s.taskService
//}

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
	if process, err := s.Process().GetProcessById(id); nil != err {
		return nil, err
	} else if err = s.Process().Check(process, strconv.FormatInt(id, 10)); err != nil {
		return nil, err
	} else if m, err := s.Process().ParseModel(process); err != nil {
		return nil, err
	} else {
		return s.startProcess(m, operator, args)
	}
}

func (s *SnakerEngineImpl) startProcess(m *ProcessModel, operator string, args map[string]interface{}) (*entity.Instance, error) {
	if execution, err := s.execute(m.Process, operator, args, 0, ""); nil != err {
		return nil, err
	} else {
		if start, err := m.GetStart(); err != nil {
			return nil, err
		} else {
			if err = start.Execute(execution); err != nil {
				return nil, err
			}
		}

		return execution.Instance, nil
	}
}

func (s *SnakerEngineImpl) execute(process *entity.Process, operator string, args map[string]interface{}, parentId int64, parentNodeName string) (*Execution, error) {
	//if instance, err := (*s.Instance()).CreateInstanceUseParentInfo(process, operator, args, parentId, parentNodeName); nil != nil {
	//	return nil, err
	//} else {
	current := &Execution{
		Engine:   s,
		Process:  process,
		Instance: &entity.Instance{},
		Args:     args,
		Operator: operator,
		Tasks:    arraylist.New(),
	}

	return current, nil
	//}

}
