package model

import "fmt"

type Motorcycle struct {
	Wheel         int
	Model         string
	Brand         string
	PublishYear   int
	MotorCapacity int
}

func NewMotorCycle() DriveAble  {
	return &Motorcycle{}
}

func (c *Motorcycle) Drive(velocity int) int {
	fmt.Println("motorcycle are going with speed of " , velocity)
	return  velocity
}

func (c *Motorcycle) Park() bool {
	fmt.Println("motor park shod")
	return true
}

func (c *Motorcycle) Break() bool {
	fmt.Println("cant break with this motor")
	return false
}
