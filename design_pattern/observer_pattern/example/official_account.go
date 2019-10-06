package main

import "fmt"

type OfficialAccount interface {
	AddSubscriber(s *Subscriber)
	RemoveSubscriber(s *Subscriber)
	NotifySubscribers(a *Article)
}

type WechatAccount struct {
	Subscribes []*Subscriber
	Article    []*Article
}

type Article struct {
	Title   string
	Content string
}

// 公众号添加订阅者
func (w *WechatAccount) AddSubscriber(s *Subscriber) {
	w.Subscribes = append(w.Subscribes, s)
	fmt.Println(s.Name, " subscribe success!")
}

// 公众号移除订阅者
func (w *WechatAccount) RemoveSubscriber(s *Subscriber) {
	var newUsers []*Subscriber

	for _, v := range w.Subscribes {
		if v.Name != s.Name {
			newUsers = append(newUsers, v)
		}
	}

	w.Subscribes = newUsers
	fmt.Println(s.Name, " cancel subscribe !")
}

// 写新文章并推送
func (w *WechatAccount) PublishArticle(title, content string) {
	fmt.Println("************************")
	fmt.Println("new article: ", title)
	article := NewArticle(title, content)
	w.Article = append(w.Article, article)
	w.NotifySubscribers(article)
}

// 通知所有订阅者有
func (w *WechatAccount) NotifySubscribers(a *Article) {
	fmt.Println("-------------------")
	fmt.Println("start notify subscribers")
	for _, v := range w.Subscribes {
		v.Notify(a)
	}
}

// 展示公众号的所有发布的文章
func (w *WechatAccount) ShowAllArticles() {
	fmt.Println("==============")
	fmt.Println("Article published by wechat official account")
	for k, v := range w.Article {
		fmt.Println(k, ": ", v.Title)
	}
}

func (w *WechatAccount) ShowAllSubscribers() {
	fmt.Println("==============")
	fmt.Println("Subscribers subscribes wechat official account")
	for k, v := range w.Subscribes {
		fmt.Println(k, ": ", v.Name)
	}
}

// 初始化公众号
func NewWechatAccount() *WechatAccount {
	return &WechatAccount{
		Subscribes: []*Subscriber{},
		Article:    []*Article{},
	}
}

// 初始化文章
func NewArticle(title, content string) *Article {
	return &Article{
		Title:   title,
		Content: content,
	}
}
