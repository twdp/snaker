package cache

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/cache"
	"sync"
)

// snaker 缓存管理
type Manager interface {

	// 根据名称获取缓存
	GetByName(name string) (cache.Cache, error)

	// 销毁所有缓存
	Destroy()
}

// 创建cache的工厂
type Factory interface {

	NewCache() cache.Cache
}

type MemoryCacheManager struct {

	sync.Mutex

	Factory Factory

	container map[string]cache.Cache
}

func (m *MemoryCacheManager) GetByName(name string) (cache.Cache, error) {
	if "" == name {
		return nil, errors.New("名称不能为空")
	}
	if s, ok := m.container[name]; !ok {
		m.Lock()
		defer m.Unlock()
		if s, ok := m.container[name]; !ok {
			if m.Factory == nil {
				s = cache.NewMemoryCache()
				b, _ := json.Marshal(map[string]int{
					"interval": 60 * 12,
				})
				s.StartAndGC(string(b))

				m.container[name] = s
			} else {
				s = m.Factory.NewCache()
				m.container[name] = s
			}
			return s, nil
		}
		return s, nil
	} else {
		return s, nil
	}
}

func (m *MemoryCacheManager) Destroy() {
	m.container = nil
}


