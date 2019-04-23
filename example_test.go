package snaker_test

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"testing"
	"tianwei.pro/snaker"

	"github.com/Knetic/govaluate"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	orm.RegisterDataBase("default", "mysql", "root:anywhere@tcp(localhost:3306)/goflow_test", 30)

	// create table
	orm.RunSyncdb("default", false, true)

	orm.Debug = true

}

//测试表达式
func TestExpression(t *testing.T) {
	fmt.Printf("--- Start TestExpression ---\n")
	expression1, _ := govaluate.NewEvaluableExpression("content")
	parameters1 := make(map[string]interface{})
	parameters1["content"] = "toTask1"
	next1, _ := expression1.Evaluate(parameters1)
	t.Logf("next1 %v", next1)

	expression2, _ := govaluate.NewEvaluableExpression("content == 200")
	parameters2 := make(map[string]interface{})
	parameters2["content"] = 200.0
	next2, _ := expression2.Evaluate(parameters2)
	t.Logf("next2 %v", next2)

	expression3, _ := govaluate.NewEvaluableExpression("content > 200")
	parameters3 := make(map[string]interface{})
	parameters3["content"] = 200.0
	next3, _ := expression3.Evaluate(parameters3)
	t.Logf("next3 %v", next3)

	expression4, _ := govaluate.NewEvaluableExpression("content < 200")
	parameters4 := make(map[string]interface{})
	parameters4["content"] = 200.0
	next4, _ := expression4.Evaluate(parameters4)
	t.Logf("next4 %v", next4)

	fmt.Printf("--- End TestExpression ---\n")
}

//测试参与方式ALL
func TestActorAll(t *testing.T) {
	fmt.Printf("--- Start TestActorAll ---\n")
	bytes := snaker.LoadXML("res/actorall.xml")
	engine := snaker.NewEngineByConfig()
	processId := engine.Deploy(bytes, "")
	args := map[string]interface{}{
		"task1.operator": []string{"1", "2"},
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %d", order.Id)

	fmt.Printf("--- End TestActorAll ---\n")
}

//测试分叉和合并
func TestForkJoin(t *testing.T) {
	fmt.Printf("--- Start TestForkJoin ---\n")
	bytes := snaker.LoadXML("res/forkjoin.xml")
	engine := snaker.NewEngineByConfig()
	processId := engine.Deploy(bytes, "")
	args := map[string]interface{}{
		"task1.operator": []string{"1"},
		"task2.operator": []string{"1"},
		"task3.operator": []string{"1"},
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %d", order.Id)
	tasks := snaker.GetActiveTasksByOrderId(order.Id)
	for _, task := range tasks {
		engine.ExecuteTask(task.Id, "1", args)
	}
	fmt.Printf("--- End TestForkJoin ---\n")
}

//测试决策1
func TestDecision1(t *testing.T) {
	fmt.Printf("--- Start TestDecision1 ---\n")
	bytes := snaker.LoadXML("res/decision1.xml")
	engine := snaker.NewEngineByConfig()
	processId := engine.Deploy(bytes, "")

	args := map[string]interface{}{
		"task2.operator": []string{"1"},
		"content":        "toTask2",
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %d", order.Id)
	fmt.Printf("--- End TestDecision1 ---\n")
}

//测试决策2
func TestDecision2(t *testing.T) {
	fmt.Printf("--- Start TestDecision2 ---\n")
	bytes := snaker.LoadXML("res/decision2.xml")
	engine := snaker.NewEngineByConfig()
	processId := engine.Deploy(bytes, "")

	args := map[string]interface{}{
		"task1.operator": []string{"1"},
		"task2.operator": []string{"1"},
		"task3.operator": []string{"1"},
		"content":        250.0,
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %d", order.Id)
	fmt.Printf("--- End TestDecision2 ---\n")
}

//简单测试
func TestSimple(t *testing.T) {
	fmt.Printf("--- Start TestSimple ---\n")
	bytes := snaker.LoadXML("res/simple.xml")
	engine := snaker.NewEngineByConfig()
	processId := engine.Deploy(bytes, "")
	args := map[string]interface{}{
		"task1.operator": []string{"1"},
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %d", order.Id)
	tasks := snaker.GetActiveTasksByOrderId(order.Id)
	for _, task := range tasks {
		engine.ExecuteTask(task.Id, "1", args)
	}
	fmt.Printf("--- End TestSimple ---\n")
}

//测试协办流程
func TestAssist(t *testing.T) {
	fmt.Printf("--- Start TestAssist ---\n")
	bytes := snaker.LoadXML("res/assist.xml")
	engine := snaker.NewEngineByConfig()
	engine.Deploy(bytes, "")
	args := map[string]interface{}{}
	order := engine.StartInstanceByName("assist", -1, "2", args)
	tasks := snaker.GetActiveTasksByOrderId(order.Id)
	for _, task := range tasks {
		snaker.CreateNewTask(task.Id, snaker.TO_ASSIST, "test")
	}
	fmt.Printf("--- End TestAssist ---\n")
}

//测试子流程1
func TestSubProcess1(t *testing.T) {
	fmt.Printf("--- Start TestSubProcess1 ---\n")
	engine := snaker.NewEngineByConfig()
	bytes := snaker.LoadXML("res/subprocess.child.xml")
	processId := engine.Deploy(bytes, "")
	bytes = snaker.LoadXML("res/subprocess.sp1.xml")
	processId = engine.Deploy(bytes, "")

	args := map[string]interface{}{
		"task1.operator": []string{"1"},
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %d", order.Id)

	tasks := snaker.GetActiveTasksByOrderId(order.Id)
	for _, task := range tasks {
		engine.ExecuteTask(task.Id, "1", args)
	}
	fmt.Printf("--- End TestSubProcess1 ---\n")
}

//测试子流程2
func TestSubProcess2(t *testing.T) {
	fmt.Printf("--- Start TestSubProcess2 ---\n")
	engine := snaker.NewEngineByConfig()
	bytes := snaker.LoadXML("res/subprocess.child.xml")
	processId := engine.Deploy(bytes, "")
	bytes = snaker.LoadXML("res/subprocess.sp2.xml")
	processId = engine.Deploy(bytes, "")

	args := map[string]interface{}{
		"task1.operator": []string{"1"},
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %d", order.Id)

	tasks := snaker.GetActiveTasksByOrderId(order.Id)
	for _, task := range tasks {
		engine.ExecuteTask(task.Id, "1", args)
	}
	fmt.Printf("--- End TestSubProcess2 ---\n")
}

//测试小组
func TestGroup(t *testing.T) {
	fmt.Printf("--- Start TestGroup ---\n")
	bytes := snaker.LoadXML("res/group.xml")
	engine := snaker.NewEngineByConfig()
	engine.Deploy(bytes, "")
	args := map[string]interface{}{
		"task1.operator": []string{"role1"},
	}
	order := engine.StartInstanceByName("group", -1, "2", args)
	t.Logf("OrderId %d", order.Id)
	tasks := snaker.GetActiveTasksByOrderId(order.Id)
	for _, task := range tasks {
		//操作人改为test时，角色对应test，会无权处理
		engine.ExecuteTask(task.Id, "ADMIN", args)
	}
	fmt.Printf("--- End TestGroup ---\n")
}

//测试权限
func TestRight(t *testing.T) {
	fmt.Printf("--- Start TestRight ---\n")
	bytes := snaker.LoadXML("res/right.xml")
	engine := snaker.NewEngineByConfig()
	processId := engine.Deploy(bytes, "")
	args := map[string]interface{}{
		"task1.operator": []string{"2"},
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %d", order.Id)
	tasks := snaker.GetActiveTasksByOrderId(order.Id)
	for _, task := range tasks {
		engine.ExecuteTask(task.Id, string(snaker.ER_ADMIN), args)
	}
	fmt.Printf("--- End TestRight ---\n")
}

//测试任务提取
func TestTake(t *testing.T) {
	fmt.Printf("--- Start TestTake ---\n")
	bytes := snaker.LoadXML("res/take.xml")
	engine := snaker.NewEngineByConfig()
	processId := engine.Deploy(bytes, "")
	args := map[string]interface{}{
		"task1.operator": []string{"1"},
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %d", order.Id)
	tasks := snaker.GetActiveTasksByOrderId(order.Id)
	for _, task := range tasks {
		snaker.TakeTask(task.Id, "1")
	}
	fmt.Printf("--- End TestTake ---\n")
}

//测试时限控制
func TestExpire(t *testing.T) {
}

//转派任务测试
func TestTransfer(t *testing.T) {
}

//测试继续Order
func TestResume(t *testing.T) {
}

//测试委托代理
func TestSurrogate(t *testing.T) {
}

//测试驳回
func TestReject(t *testing.T) {
}

//测试抄送
func TestCC(t *testing.T) {
}

//测试局部拦截器
func TestInterceptor(t *testing.T) {
}
