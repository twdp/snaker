package snaker

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"sync"
	"time"
)

type BaseIdModel struct {
	Id int64
}

type BaseTimeModel struct {
	CreatedAt orm.DateTimeField `orm:"auto_now_add"`
	UpdatedAt orm.DateTimeField `orm:"auto_now"`
}

func UUID() int64 {
	date := time.Now().Format("20060102")
	id := fmt.Sprintf("%d", unix())
	i, err := strconv.ParseInt(fmt.Sprintf("%s%s", date, id[4: len(id) - 1]), 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

var lock sync.Locker = new(sync.Mutex)

func unix() int64 {
	lock.Lock()
	defer lock.Unlock()
	return time.Now().UnixNano() / 1000
}
