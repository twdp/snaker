package snaker

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

)

//字符串转整型
func StrToInt(value string) int {
	if value == "" {
		return 0
	}
	val, _ := strconv.Atoi(value)
	return val
}

//整型转字符串
func IntToStr(value int) string {
	return strconv.Itoa(value)
}

//装载XML文件
func LoadXML(xmlFile string) []byte {
	content, err := ioutil.ReadFile(xmlFile)
	if err != nil {
		logs.Error("parse xml failed. xmlFile: %v", xmlFile)
		panic(fmt.Errorf("error to read xml file!"))
	}
	return content
}

//map转json
func MapToJson(v map[string]interface{}) string {
	if v == nil {
		return ""
	}
	js := simplejson.New()
	js.Set("map", v)
	ret, _ := js.Get("map").Encode()
	return string(ret)
}

//json转map
func JsonToMap(v string) map[string]interface{} {
	js, _ := simplejson.NewJson([]byte(v))
	return js.MustMap()
}

//删除Slice中的元素
func StringsRemove(strings []string, start, end int) []string {
	return append(strings[:start], strings[end:]...)
}

//删除Slice中的元素
func StringsRemoveAtIndex(strings []string, index int) []string {
	return StringsRemove(strings, index, index+1)
}

//格式化时间字符串
func FormatTime(t time.Time, layout string) string {
	if t.IsZero() {
		return ""
	} else {
		return t.Format(layout)
	}
}

func ProcessTime(args map[string]interface{}, timeParam string) time.Time {
	if timeParam != "" {
		var timeStr string
		if timeInf, ok := args[timeParam]; ok {
			timeStr = timeInf.(string)
		} else {
			timeStr = timeParam
		}
		the_time, err := time.Parse(STD_TIME_LAYOUT, timeStr)
		if err == nil {
			return the_time
		}
	}
	return time.Time{}
}

//TaskType转换
func ProcessTaskType(taskType TASK_TYPE) TASK_ORDER {
	if strings.ToUpper(string(taskType)) == string(TT_ASSIST) {
		return TO_ASSIST
	} else {
		return TO_MAJOR
	}
}

//PerformType转换
func ProcessPerformType(performType PERFORM_TYPE) PERFORM_ORDER {
	if strings.ToUpper(string(performType)) == string(PT_ALL) {
		return PO_ALL
	} else {
		return PO_ANY
	}
}