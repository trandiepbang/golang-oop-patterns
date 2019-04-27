package main

import "fmt"

type Robot struct {
	name     string
	language string
}

func NewRobot() Robot {
	return Robot{}
}

func (r Robot) SetLang(lang string) Robot {
	r.language = lang
	return r
}

func (r Robot) SetName(name string) Robot {
	r.name = name
	return r
}

func (r Robot) Build() Robot {
	return r
}

func main() {
	robotBuilder := NewRobot().SetLang("english").SetName("x-men").Build()
	fmt.Printf("Robot name is %s - language he can speaks is %s ", robotBuilder.name, robotBuilder.language)
}
