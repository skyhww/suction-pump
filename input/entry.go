package input

import (
	"errors"
	"fmt"
	"suction-pump/version"
	"sync"
)

//请求接口规范
type Input struct {
	//名称
	name string
	id   int
	path string
	//输入对象
	entry *[]Entry
	//负责创建以及维护Input
	factory Factory
	//Factory在创建input时负责维护此字段
	rw *sync.RWMutex
	//状态机
	status int
}

func (input *Input) getFactory() Factory {
	return input.factory
}

func (input *Input) Lock() {
	input.rw.Lock()
}
func (input *Input) unLock() {
	input.rw.Unlock()
}

func (input *Input) getStatus() int {
	input.rw.RLock()
	defer input.rw.Unlock()
	return input.status
}
func (input *Input) reload() error {
	input.rw.Lock()
	input.status = version.OFF
	defer input.rw.Unlock()
	err := input.factory.Reload(input)
	if err != nil {
		return err
	}
	input.status = version.ON
	return nil
}

func (input *Input) install() error {
	input.rw.Lock()
	defer input.rw.Unlock()
	if input.id == 0 {
		errors.New("模块id不能为空！")
	}
	if input.factory == nil {
		errors.New("模块工厂不能为空！")
	}
	if input.getStatus() == version.ON {
		errors.New(fmt.Sprintf("模块：%d,%s已处于运行状态！", input.id, input.name))
	}
	input.status = version.ON
	return nil
}
func (input *Input) unInstall() error {
	input.rw.Lock()
	defer input.rw.Unlock()
	if input.getStatus() == version.OFF {
		errors.New(fmt.Sprintf("模块：%d,%s已处于卸载状态！", input.id, input.name))
	}
	err := input.factory.unInstall(input)
	if err != nil {
		return err
	}
	input.status = version.OFF
	return nil
}

//只提供get方法供外部访问
func (input *Input) GetName() string {
	return input.name
}
func (input *Input) GetPath() string {
	return input.path
}
func (input *Input) GetId() int {
	return input.id
}
func (input *Input) GetEntry() []Entry {
	if input.entry == nil {
		return nil
	}
	return *input.entry
}

//输入字段
type Entry struct {
	name      string
	array     bool
	required  bool
	validator []FieldValidator
}

func (entry *Entry) GetName() string {
	return entry.name
}

func (entry *Entry) isArray() bool {
	return entry.array
}
func (entry *Entry) isRequired() bool {
	return entry.required
}
func (entry *Entry) Validate(value string) error {
	for _, validator := range entry.validator {
		err := validator.Validate(value)
		if err != nil {
			return err
		}
	}
	return nil
}

type PageInput struct {
	input    *Input
	pageSize int
	pageNo   int
}

func (pageInput *PageInput) GetName() string {
	return pageInput.input.GetName()
}
func (pageInput *PageInput) GetId() int {
	return pageInput.input.GetId()
}
func (pageInput *PageInput) GetEntry() []Entry {
	return pageInput.input.GetEntry()
}
func (pageInput *PageInput) GetPageSize() int {
	return pageInput.pageSize
}
func (pageInput *PageInput) GetPageNo() int {
	return pageInput.pageNo
}
