package main

import (
	"fmt"

	"awesomeProject/internal/apple"
	"awesomeProject/internal/burn"
	"awesomeProject/internal/duck"
	"awesomeProject/internal/freeze"
	"awesomeProject/internal/mix"
	"awesomeProject/internal/potato"
)

func main() {
	fmt.Println()
	f := freeze.New(
		burn.New(
			mix.New(
				duck.New(),
				apple.New(),
				potato.New(),
				potato.New(),
				potato.New(),
				freeze.New(apple.New()),
			),
		),
	)
	fmt.Println(f.Name())
	warmed := f.Warm()
	fmt.Println(warmed.Name())
	Webinar(ASMyasnikov{})
}

type ASMyasnikov struct{}

func (ASMyasnikov) WriteCode() {

}

func (ASMyasnikov) Talk() {

}

func (ASMyasnikov) Cook() {

}

func (ASMyasnikov) ManageMinions() {

}

type speaker interface {
	WriteCode()
	Talk()
}

func Webinar(speaker speaker) {

}

type getter interface {
	Get() (string, error)
}

func HandleGet(s getter) error {
	return nil
}

func DoSome(ints []int) int {
	//return Do(ints)
	return Do(func() (args []interface{}) {
		args = make([]any, len(ints))
		for i, arg := range ints {
			args[i] = arg
		}
		return args
	}())
}

func Do(args []interface{}) int {
	return 0
}
