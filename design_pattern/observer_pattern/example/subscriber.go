package main

import "fmt"

type SubscriberInterface interface {
	Notify(a *Article)
}

type Subscriber struct {
	Name string
	Sex  int
}

func NewSubscriber(name string, sex int) *Subscriber {
	return &Subscriber{
		Name: name,
		Sex:  sex,
	}
}

// 订阅者收到公众号的推送
func (s *Subscriber) Notify(a *Article) {
	fmt.Println("++++++++++++++++++++++++++++++++++")
	if s.Sex == 1 {
		fmt.Println("Mr.", s.Name, "you have a new published article from wechat account")
	} else {
		fmt.Println("Miss.", s.Name, "you have a new published article from wechat account")
	}
	fmt.Println("title :", a.Title)
	fmt.Println("content :", a.Content)
}

// 订阅者新增公众号订阅
func (s *Subscriber) Subscribe(w OfficialAccount) {
	w.AddSubscriber(s)
}

// 订阅者取消公众号订阅
func (s *Subscriber) CancelSubscribe(w OfficialAccount) {
	w.RemoveSubscriber(s)
}
