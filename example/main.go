package main

import (
	"fmt"
	"github.com/no13bus/go-rubychina"
)

func main() {
	// var nodes Nodes
	nodes, err := rubychina.GetNodes()
	if err != nil {
		fmt.Println("node error")
		return
	}
	topics, err := rubychina.GetTopics()
	if err != nil {
		fmt.Println("topics error")
		return
	}
	users, err := rubychina.GetUsers()
	if err != nil {
		fmt.Println("users error")
		return
	}
	user, err := rubychina.GetUserByName("no13bus")
	if err != nil {
		fmt.Println("userbyname error")
		return
	}
	topic, err := rubychina.GetTopicByID(1)
	sites, err := rubychina.GetSites()
	fa_topics, err := rubychina.GetFavoriteTopic("no13bus")

	fmt.Println(nodes[0].Name)
	fmt.Println(topics[0].Title)
	fmt.Println(users[0].Name)
	fmt.Println(user.Topics[0].Title)
	fmt.Println(topic.Title)
	fmt.Println(topic.Replies[0].Body)
	fmt.Println(sites[0].Sites[0].Desc)
	// fmt.Println(len(fa_topics))
	fmt.Println(fa_topics)

}
