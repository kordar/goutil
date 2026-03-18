package goutil

import (
	"errors"
	"log"
	"sync"
)

var instanceOfDbPool *DbConnPool
var onceOfDb sync.Once

type DBItem interface {
	GetName() string
	GetInstance() interface{}
	Close() error
}

type DbConnPool struct {
	handle map[string]DBItem
	mu     sync.RWMutex
}

func GetDbPool() *DbConnPool {
	onceOfDb.Do(func() {
		instanceOfDbPool = &DbConnPool{handle: make(map[string]DBItem)}
	})
	return instanceOfDbPool
}

// InitDataPool /*
// 初始化数据库连接(可在mail()适当位置调用)
func (m *DbConnPool) InitDataPool(items ...DBItem) (issucc bool) {
	if err := m.InitDataPoolE(items...); err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (m *DbConnPool) InitDataPoolE(items ...DBItem) error {
	for _, item := range items {
		if item == nil {
			continue
		}
		if existing, _ := m.Item(item.GetName()); existing != nil {
			log.Printf("实例[%s]已存在\n", item.GetName())
			continue
		}
		if err := m.Add(item); err != nil {
			return err
		}
	}

	// 关闭数据库，db会被多个goroutine共享，可以不调用
	// defer db.Close()
	return nil
}

// Add 添加数据库实例
func (m *DbConnPool) Add(db DBItem) error {
	if db == nil {
		return errors.New("db item is nil")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if m.handle[db.GetName()] != nil {
		return errors.New("db instance already exists")
	}
	m.handle[db.GetName()] = db
	return nil
}

// Remove 移除句柄
func (m *DbConnPool) Remove(name string) {
	m.mu.Lock()
	item := m.handle[name]
	delete(m.handle, name)
	m.mu.Unlock()

	if item == nil {
		return
	}
	if err := item.Close(); err != nil {
		log.Printf("移除句柄=%v\n", err)
	}
}

// Handle /*
// 对外获取数据库连接对象db
func (m *DbConnPool) Handle(name string) (conn interface{}) {
	m.mu.RLock()
	item := m.handle[name]
	m.mu.RUnlock()
	return item
}

func (m *DbConnPool) Item(name string) (DBItem, bool) {
	m.mu.RLock()
	item, ok := m.handle[name]
	m.mu.RUnlock()
	return item, ok
}

func (m *DbConnPool) Instance(name string) interface{} {
	item, ok := m.Item(name)
	if !ok || item == nil {
		return nil
	}
	return item.GetInstance()
}
