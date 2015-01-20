package main

import (
	"fmt"
	"github.com/no13bus/go-rubychina"
)

func main() {
	// var nodes Nodes
	nodes, err := rubychina.GetNodes()
	if err != nil {
		fmt.Printf("node error:%s" % err)
		return
	}
	topics, err := rubychina.GetTopics()
	if err != nil {
		fmt.Printf("topics error:%s" % err)
		return
	}
	users, err := rubychina.GetUsers()
	if err != nil {
		fmt.Printf("users error:%s" % err)
		return
	}
	user, err := rubychina.GetUserByName("no13bus")
	if err != nil {
		fmt.Printf("userbyname error:%s" % err)
		return
	}
	topic, err := rubychina.GetTopicByID(1)
	if err != nil {
		fmt.Printf("GetTopicByID error:%s" % err)
		return
	}
	sites, err := rubychina.GetSites()
	if err != nil {
		fmt.Printf("GetSites error:%s" % err)
		return
	}
	fa_topics, err := rubychina.GetFavoriteTopic("no13bus")
	if err != nil {
		fmt.Printf("GetFavoriteTopic error:%s" % err)
		return
	}

	fmt.Println(nodes[0].Name)
	fmt.Println(topics[0].Title)
	fmt.Println(users[0].Name)
	fmt.Println(user.Topics[0].Title)
	fmt.Println(topic.Title)
	fmt.Println(topic.Replies[0].Body)
	fmt.Println(sites[0].Sites[0].Desc)
	if len(fa_topics) > 0 {
		fmt.Println(fa_topics[0].Title)
	}

}
