package snaker

import "tianwei.pro/snaker/entity"

type ProcessAccess interface {

	// 根据id查询process信息
	GetById(id int64) (*entity.Process, error)
}