package main

import (
	"ThinkGOChat/thinkutils/logger"
	"github.com/emirpasic/gods/lists/arraylist"
	"reflect"
	"runtime"
)

var (
	log *logger.LocalLogger = logger.DefaultLogger()
)

type IAnimal interface {
	Eat() string
	OnInit() string
}

type BaseAnimal struct {
	IAnimal
}

func (this *BaseAnimal) Eat() string {
	return "BaseAnimal"
}

func (this *BaseAnimal) OnInit() string {
	logger.Info(reflect.TypeOf(this))
	return "Base OnInit"
}

func (this *BaseAnimal) OnInitReal(child interface{}) string {
	logger.Info(reflect.TypeOf(this))
	ref := reflect.ValueOf(child)
	method := ref.MethodByName("OnInit")
	if (method.IsValid()) {
		r := method.Call(make([]reflect.Value, 0))
		return r[0].String()
	} else {
		// 错误处理
	}

	return ""
}

type Cat struct {
	BaseAnimal
}

func (this *Cat) Eat() string {
	return "Cat"
}

func (this *Cat) OnInit() string {
	return this.Eat()
}

type Dog struct {
	BaseAnimal
}

func (this *Dog) OnInit() string {
	return this.Eat()
}

func (this *Dog) Eat() string {
	return "Dog"
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    log.Info("Hello World")

    //litAnimal := [Ba]{}

	//cat := &Cat{}
	dog := &Dog{}
	logger.Info(dog.OnInitReal(dog))

	lstAnimal := arraylist.New()
	//lstAnimal.Add(cat)
	lstAnimal.Add(dog)

	for i := 0; i<lstAnimal.Size(); i++ {
		item, _ := lstAnimal.Get(i)
		logger.Info(item.(IAnimal).OnInit())
	}
}
