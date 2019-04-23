package snaker

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"reflect"
)

// 本工程数据库使用beego orm

func Save(inf interface{}) {
	_, err := orm.NewOrm().Insert(inf)
	if err != nil {

		logs.Error("save obj failed. inf: %v", reflect.TypeOf(inf))
		panic(errors.New("保存失败"))
	}
}

func Update(inf interface{}) {
	_, e := orm.NewOrm().Update(inf)
	if e != nil {
		logs.Error("fail to update %v", reflect.TypeOf(inf))
		panic(errors.New("更新失败"))
	}
}

//删除实体对象
func Delete(inf interface{}) {
	orm.NewOrm().Delete(inf)
}

