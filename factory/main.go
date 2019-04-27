package main

import "fmt"

const (
	Arm    = "arm"
	Legs   = "legs"
	Head   = "head"
	Finger = "finger"
)

type IRobot interface {
	GetBodyPart(bodyPart string) string
}

type RobotBodyParts struct{}

func NewRobot() IRobot {
	return &RobotBodyParts{}
}

func (r RobotBodyParts) GetBodyPart(bodyPart string) string {
	if bodyPart == Arm {
		return Arm
	}
	if bodyPart == Legs {
		return Legs
	}
	if bodyPart == Head {
		return Head
	}
	if bodyPart == Finger {
		return Finger
	}
	return ""
}

func main() {
	robotBodyPart := NewRobot()
	fmt.Printf("Robot body : %s \n", robotBodyPart.GetBodyPart("arm"))
	fmt.Printf("Robot body : %s \n", robotBodyPart.GetBodyPart("legs"))
}
