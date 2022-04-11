package main

import (
	"ThinkGOChat/thinkutils/logger"
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"runtime"
)

var (
	log *logger.LocalLogger = logger.DefaultLogger()
)

type Animal interface {
	Eat()
}

type Cat struct {
}

func (this *Cat) Eat() {
	fmt.Println("FXXK")
}

type Dog struct {
}

func (this *Dog) Eat() {
	fmt.Println("FXXK123")
}


func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    log.Info("Hello World")

    cat := &Cat{}
	dog := &Dog{}

	lstAnimal := arraylist.New()
	lstAnimal.Add(cat)
	lstAnimal.Add(dog)

	for i := 0; i<lstAnimal.Size(); i++ {
		item, _ := lstAnimal.Get(i)
		item.(Animal).Eat()
	}
}
