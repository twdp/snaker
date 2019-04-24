package snaker

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

//委托代理
type Surrogate struct {
	BaseIdModel

	ProcessName string `orm:"size(36)"`       //流程名称
	Operator    string `orm:"size(36);index"` //授权人
	Surrogate   string `orm:"size(36)"`               //代理人
	OpTime      time.Time                             //操作时间
	StartTime   time.Time                             //开始时间
	EndTime     time.Time                             //结束时间
	State       SURROGATE_STATUS                      //状态

	BaseTimeModel
}

func init() {
	orm.RegisterModelWithPrefix("snaker_", new(Surrogate))
}

//得到代理人（通过SQL）
func GetSurrogateSQL(querystring string, args ...interface{}) []*Surrogate {
	surrogates := make([]*Surrogate, 0)
	_, err := orm.NewOrm().Raw(querystring, args).QueryRows(surrogates)
	if err != nil {
		logs.Error("fail to get surrogatesql. sql: {}, args: {}, cause: {}", querystring, args, err)
	}
	return surrogates
}
