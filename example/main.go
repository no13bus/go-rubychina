package main

import (
	"fmt"
	"github.com/no13bus/go-rubychina"
)

func main() {
	// var nodes Nodes
	nodes, err := rubychina.GetNodes()
	if err != nil {
		fmt.Println(err)
		return
	}
	topics, err := rubychina.GetTopics()
	if err != nil {
		fmt.Println(err)
		return
	}
	users, err := rubychina.GetUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	user, err := rubychina.GetUserByName("no13bus")
	if err != nil {
		fmt.Println(err)
		return
	}
	topic, err := rubychina.GetTopicByID(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	sites, err := rubychina.GetSites()
	if err != nil {
		fmt.Println(err)
		return
	}
	fa_topics, err := rubychina.GetFavoriteTopic("no13bus")
	if err != nil {
		fmt.Println(err)
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
