package main

import "fmt"

type Notifier interface {
	Send(text string , receiver string)
	Delete(messageId int64)
}

type Sms struct {
	Address string
}

type Email struct {
	MailAddress string
}

func (e *Email) Send(text string, receiver string) {
	///send a request to mail server
}

func (e *Email) Delete(messageId int64) {
	return
}

type PushNotif struct {
	Address string
}

func (p *PushNotif) Send(text string, receiver string) {
	panic("implement me")
}

func (p *PushNotif) Delete(messageId int64) {
	panic("implement me")
}

func (s *Sms)Send(text string , receiver string)  {

}

func (s *Sms)Delete(id int64)  {

}

func NewEmail() Notifier  {
	return &Email{}
}

func NewSms() Notifier {
	return &Sms{}
}

func NewPush() Notifier  {
	return  &PushNotif{}
}

func main()  {
	var notifChannel string
	fmt.Scanln(&notifChannel)

	channel := getNotifChannel(notifChannel)

	channel.Send("salam" , "arash@gmail.com")
	fmt.Println(fmt.Sprintf("%T" , channel))
	fmt.Printf("%T" , channel)
	}

type Ak47 struct {

}

func getNotifChannel(channel string) Notifier {
	switch channel {
	case "sms":
		return NewSms()
	case "push":
		return NewPush()
	default:
		return NewEmail()
	}
	return nil
}

