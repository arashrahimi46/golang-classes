package model

type DriveAble interface {
	Drive(velocity int) (speed int)
	Park() bool
	Break() bool
}

type TestAble interface {

}
