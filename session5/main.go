package main

import (
	"fmt"
	"golang-class/session5/model"
)

func main()  {
	car1 := model.NewCar(10 , "Red" , "F10" , "Benz" , uint(2003),false)
	motor1 := model.NewMotorCycle()

	Drive(motor1)
	Drive(car1)
}

func Drive(vehicle model.DriveAble)  {
	fmt.Println(vehicle.Drive(25))
	JustPrint("test parameter", 20 , "string param" , 222.255 , 98 , -98 , model.Motorcycle{} , vehicle)
}

func JustPrint(test string, age int , value ...interface{})  {
	fmt.Println(value)
}
