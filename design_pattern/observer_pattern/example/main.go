package main

import "fmt"

const (
	MALE   = 1
	FEMALE = 2
)

func main() {

	wechatOfficialAccount := NewWechatAccount()

	jack := NewSubscriber("jack", MALE)
	lucy := NewSubscriber("lucy", FEMALE)
	tommy := NewSubscriber("tommy", MALE)

	fmt.Println("")

	jack.Subscribe(wechatOfficialAccount)
	lucy.Subscribe(wechatOfficialAccount)
	tommy.Subscribe(wechatOfficialAccount)

	fmt.Println("")

	wechatOfficialAccount.PublishArticle("Hell,Everyone", "thanks for your subscription")

	fmt.Println("")
	lucy.CancelSubscribe(wechatOfficialAccount)

	fmt.Println("")
	wechatOfficialAccount.PublishArticle("Study", "Good good study, Day day up")

	fmt.Println("")
	wechatOfficialAccount.ShowAllArticles()

	fmt.Println("")
	wechatOfficialAccount.ShowAllSubscribers()
}
