package input

import (
	"errors"
	"sync"
)

type Factory interface {
	Reload(input *Input) error
	getInput(id int) (*Input, error)
	unInstall(input *Input) error
}

type AbstractFactory struct {
	lock      *sync.RWMutex
	dbFactory Factory
}

func (factory *AbstractFactory) getDbFactory() Factory {
	if factory.dbFactory != nil {
		return factory.dbFactory
	}
	factory.lock.Lock()
	defer factory.lock.Unlock()
	if factory.dbFactory != nil {
		return factory.dbFactory
	}
	factory.dbFactory = &DbFactory{}
	return factory.dbFactory
}

var Single *AbstractFactory
var once sync.Once

func GetSingleton() *AbstractFactory {
	once.Do(func() {
		Single = &AbstractFactory{}
		Single.lock = &sync.RWMutex{}
	})
	return Single
}

type DbFactory struct {
	//此缓存数据量较小，不考虑缓存击穿跟缓存抖动等问题
	cache map[int]*Input
	sync  sync.RWMutex
}

func (factory *DbFactory) Reload(input *Input) error {
	if input == nil {
		return errors.New("试图重新加载一个空的input引用！")
	}
	if input.getFactory() != factory {
		return errors.New("非当前工厂生产的对象，无权限操作此对象！")
	}
	factory.sync.Lock()
	defer factory.sync.Unlock()
	in, err := factory.getInput(input.GetId())
	if err != nil {
		return err
	}
	//更换指针
	*input = *in
	factory.cache[in.GetId()] = in
	return input.install()
}

func (factory *DbFactory) getInput(id int) (*Input, error) {
	//不要求及时生效，可以暂时返回脏数据
	input := factory.cache[id]
	if input != nil {
		return input, nil
	}
	factory.sync.Lock()
	defer factory.sync.Unlock()
	input = factory.cache[id]
	if input != nil {
		return input, nil
	}
	//到数据库获取input
	return nil, nil
}
func (factory *DbFactory) unInstall(input *Input) error {
	if input == nil {
		return errors.New("试图重新加载一个空的input引用！")
	}
	if input.getFactory() != factory {
		return errors.New("非当前工厂生产的对象，无权限操作此对象！")
	}
	factory.sync.Lock()
	defer factory.sync.Unlock()
	in := factory.cache[input.id]
	if in == nil {
		return errors.New("脱离了factory管理的指针！")
	}
	if in != input {
		return errors.New("实际引用对象不一致！")
	}
	delete(factory.cache, in.id)
	return nil
}
